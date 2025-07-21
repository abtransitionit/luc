/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package packagex

import (
	"fmt"
	"sync"

	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/abtransitionit/luc/pkg/util"
	"github.com/abtransitionit/luc/pkg/util/dnfapt"
)

func remoteInstall(in <-chan PipelineData, out chan<- PipelineData, nbVm int, vms []string, packages []string) {
	nbWorker := nbVm // as many workers as VM
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

			// ------------------------

			// remote install dnfapt packages
			logx.L.Debugf("[%s] remote installing packages", vm)
			for _, pkgName := range data.PackageList {
				// --- CHECK CONDITION ---
				if pkgName == "uidmap" && data.OsFamily != "debian" {
					logx.L.Debugf("[%s] [%s] Skipping package : OS family is %s, not 'debian'.", vm, pkgName, data.OsFamily)
					continue
				}
				// ------------------------
				logx.L.Debugf("[%s] [%s] remote installing package", vm, pkgName)
				output, err := dnfapt.RInstallP(vm, data.OsFamily, pkgName)
				if err != nil {
					data.Err = fmt.Errorf("%v, %v", err, output)
					logx.L.Debugf("[%s] ❌ error detected 1", vm)
					out <- data
					continue
				}
			}
			logx.L.Debugf("[%s] remote installed packages", vm)

			// set reboot status
			logx.L.Debugf("[%s] getting reboot status", vm)
			rebootStatus, err := util.GetPropertyRemote(vm, "rebootstatus")
			if err != nil {
				data.Err = fmt.Errorf("%v, %s", err, rebootStatus)
				logx.L.Debugf("[%s] ❌ Error detected 2", vm)
			}
			logx.L.Debugf("[%s] got reboot status : %s", vm, rebootStatus)

			// set instance property
			data.RebootStatus = rebootStatus

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
