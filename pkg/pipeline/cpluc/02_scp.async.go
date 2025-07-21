/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package cpluc

import (
	"fmt"
	"strings"
	"sync"

	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/abtransitionit/luc/pkg/util"
)

func scpLuc(in <-chan PipelineData, out chan<- PipelineData, vmList string) {
	vms := strings.Fields(vmList) // convert ListAsString to slice
	nbWorker := len(vms)
	var wg sync.WaitGroup
	defer close(out)

	// Worker function
	worker := func() {
		defer wg.Done()
		for data := range in {
			vm := data.VmName
			if data.Err != nil {
				out <- data
				logx.L.Debugf("❌ Previous error detected")
				continue
			}

			// scp luc from local to remote VM
			logx.L.Debugf("[%s] coping LUC from local to remote", vm)
			cli := fmt.Sprintf("scp %s %s:%s", data.localOutXptf, vm, data.remoteTmpPath)
			outp, err := util.RunCLILocal(cli)
			if err != nil {
				data.Err = fmt.Errorf("❌ Error: %v, %s", err, outp)
				logx.L.Debugf("[%s] ❌ error detected", vm)
				out <- data
				continue
			}
			logx.L.Debugf("[%s] copied LUC from local to remote tmp path", vm)

			// remote mv luc from temp to final folder
			logx.L.Debugf("[%s] remote moving LUC from tmp to final path", vm)
			cli = fmt.Sprintf("sudo mv  %s %s", data.remoteTmpPath, data.remoteExePath)
			outp, err = util.RunCLIRemote(vm, cli)
			if err != nil {
				data.Err = fmt.Errorf("❌ Error: %v, %s", err, outp)
				logx.L.Debugf("[%s] ❌ error detected", vm)
				out <- data
				continue
			}
			logx.L.Debugf("[%s] remote moved LUC from tmp to final path", vm)

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
