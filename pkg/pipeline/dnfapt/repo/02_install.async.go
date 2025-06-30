/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package packagex

import (
	"sync"

	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/abtransitionit/luc/pkg/util"
)

func remoteInstall(in <-chan PipelineData, out chan<- PipelineData, nbVm int, vms []string, packages []string) {
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

			// remote install dnfapt packages
			logx.L.Debugf("[%s] remote installing packages", data.HostName)
			for _, pkgName := range data.PackageList {
				logx.L.Debugf("[%s] remote installing package %s", data.HostName, pkgName)
				// _, err := dnfapt.RemoteInstall(data.HostName, data.OsFamily, pkgName)
				// if err != nil {
				// 	data.Err = err
				// 	logx.L.Debugf("[%s] ❌ error detected 1", data.HostName)
				// 	out <- data
				// 	continue
				// }
			}
			logx.L.Debugf("[%s] remote installed packages", data.HostName)

			// set reboot status
			logx.L.Debugf("[%s] getting reboot status", data.HostName)
			rebootStatus, err := util.GetRemoteProperty("rebootstatus", data.HostName)
			if err != nil {
				data.Err = err
				logx.L.Debugf("[%s] ❌ Error detected 2", data.HostName)
			}
			logx.L.Debugf("[%s] got reboot status : ", data.HostName, rebootStatus)

			// set instance property
			data.RebootStatus = rebootStatus

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
