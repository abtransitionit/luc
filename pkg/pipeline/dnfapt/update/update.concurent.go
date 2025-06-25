/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package update

import (
	"sync"

	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/abtransitionit/luc/pkg/util/dnfapt"
)

func updateConcurent(in <-chan PipelineData, out chan<- PipelineData, nbWorker int) {
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

			// Update the OS
			_, err := dnfapt.UpdateOs()
			if err != nil {
				data.Err = err
				logx.L.Debugf("Error detected")
				out <- data
				continue
			}

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
