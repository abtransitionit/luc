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
			if data.Err != nil {
				out <- data
				logx.L.Debugf("❌ Previous error detected")
				continue
			}

			// remote upgrade
			logx.L.Debugf("[%s] remote upgrading", data.HostName)
			_, err := dnfapt.RUpgrade(data.HostName, data.OsFamily)
			if err != nil {
				data.Err = err
				logx.L.Debugf("[%s] ❌ error detected 1", data.HostName)
				out <- data
				continue
			}
			logx.L.Debugf("[%s] remote upgraded", data.HostName)

			// set reboot status
			logx.L.Debugf("[%s] getting reboot status", data.HostName)
			rebootStatus, err := util.GetPropertyRemote("rebootstatus", data.HostName)
			if err != nil {
				data.Err = err
				logx.L.Debugf("[%s] ❌ Error detected 2", data.HostName)
			}
			logx.L.Debugf("[%s] got reboot status : %s", data.HostName, rebootStatus)

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
