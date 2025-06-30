/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package packagex

import (
	"strings"
	"sync"
	"time"

	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/abtransitionit/luc/pkg/util"
)

func remoteReboot(in <-chan PipelineData, out chan<- PipelineData, nbVm int) {
	nbWorker := nbVm // as many workers as VM
	var wg sync.WaitGroup
	defer close(out)

	// Worker function
	worker := func() {
		defer wg.Done()
		for data := range in {
			if data.Err != nil {
				out <- data
				logx.L.Debugf("❌ Previous error detected")
				continue
			}

			// remote reboot if needed
			if strings.ToLower(strings.TrimSpace(data.RebootStatus)) == "false" {
				logx.L.Debugf("[%s] Skipping reboot : reboot status is false", data.HostName)
				out <- data
				continue
			}
			// reboot
			logx.L.Debugf("[%s] remote rebooting", data.HostName)
			util.RemoteReboot(data.HostName)
			logx.L.Debugf("[%s] remote rebooted", data.HostName)

			// wait ssh reachable
			logx.L.Debugf("[%s] getting ssh reachability", data.HostName)
			for {
				isReachable, err := util.IsSshConfiguredVmSshReachable(data.HostName)
				if err != nil {
					data.Err = err
					logx.L.Debugf("[%s] ❌ error detected 2", data.HostName)
					out <- data
					continue
				} else if !isReachable {
					break
				}
				time.Sleep(2 * time.Second)
			}
			// log end wait
			logx.L.Debugf("[%s] got ssh reachability", data.HostName)

			// set instance property
			data.RebootStatus = "false"

			// get property
			kernelVersion, err := util.GetRemoteProperty("oskversion", data.HostName)
			if err != nil {
				data.Err = err
				logx.L.Debugf("[%s] ❌ Error detected 3", data.HostName)
				out <- data
				continue
			}

			// set instance property
			data.OskernelVersionAfter = kernelVersion

			// send
			out <- data
		} // for
	} // worker

	// Start N workers
	wg.Add(nbWorker)
	for i := 0; i < nbWorker; i++ {
		go worker()
	}

	wg.Wait()
}
