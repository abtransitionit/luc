/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package rupgrade

import (
	"sync"

	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/abtransitionit/luc/pkg/util"
	"github.com/abtransitionit/luc/pkg/util/dnfapt"
)

func rUpgrade(in <-chan PipelineData, out chan<- PipelineData, nbWorker int) {
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

			// remote upgrade
			logx.L.Debugf("[%s] remote upgrading", vm)
			_, err := dnfapt.RUpgrade(vm, data.OsFamily)
			if err != nil {
				data.Err = err
				logx.L.Debugf("[%s] ❌ error detected 1", vm)
				out <- data
				continue
			}
			logx.L.Debugf("[%s] remote upgraded", vm)

			// set reboot status
			logx.L.Debugf("[%s] getting reboot status", vm)
			rebootStatus, err := util.GetPropertyRemote(vm, "rebootstatus")
			if err != nil {
				data.Err = err
				logx.L.Debugf("[%s] ❌ Error detected 2", vm)
			}
			logx.L.Debugf("[%s] got reboot status : %s", vm, rebootStatus)

			// set instance property
			data.RebootStatus = rebootStatus

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
