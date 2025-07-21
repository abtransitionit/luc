/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package upgrade

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/abtransitionit/luc/pkg/util"
)

func remoteReboot(in <-chan PipelineData, out chan<- PipelineData, nbWorker int) {
	var wg sync.WaitGroup
	defer close(out)

	// Worker function
	worker := func() {
		defer wg.Done()
		for data := range in {

			vm := data.HostName

			if data.Err != nil {
				out <- data
				logx.L.Debugf("❌ Previous error detected")
				continue
			}

			// remote reboot if needed
			if strings.ToLower(strings.TrimSpace(data.RebootStatus)) == "true" {

				// reboot
				logx.L.Debugf("[%s] remote rebooting", vm)
				err := util.RemoteReboot(vm)
				if err != nil {
					data.Err = err
					logx.L.Debugf("❌ [%s] error detected 1", vm)
					out <- data
					continue
				}
				logx.L.Debugf("[%s] remote rebooted", vm)

				// wait ssh to become reachable
				logx.L.Debugf("[%s] getting ssh reachability", vm)
				for {
					// get reachability
					isReachable, err := util.GetPropertyLocal("sshreachability", vm)

					// exit for loop if vm is reachable
					if err == nil && isReachable == "true" {
						break
					}

					// repeat for code if vm is not reachable
					logx.L.Debugf("[%s] ssh not reachable, waiting ..", vm)
					time.Sleep(2 * time.Second)
				} // end for loop

				// set instance property
				data.RebootStatus = "false"

			} // if

			// get property
			kernelVersion, err := util.GetPropertyRemote(vm, "oskversion")
			if err != nil {
				data.Err = fmt.Errorf("%v, %s", err, kernelVersion)
				logx.L.Debugf("[%s] ❌ Error detected 3", vm)
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

// for {
// 	isReachable, err := util.IsSshConfiguredVmSshReachable(vm)
// 	if err != nil {
// 		data.Err = fmt.Errorf("%v, %v", err, isReachable)
// 		logx.L.Debugf("❌ [%s] error detected 2", vm)
// 		out <- data
// 		continue
// 	}
// 	if err == nil && isReachable {
// 		break
// 	}
// 	time.Sleep(2 * time.Second)
// }
// log end wait
// logx.L.Debugf("[%s] got ssh reachability", vm)
