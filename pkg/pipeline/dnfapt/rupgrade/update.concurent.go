/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package rupgrade

import (
	"fmt"
	"strings"
	"sync"

	"github.com/abtransitionit/luc/pkg/logx"
)

func upgradeConcurent(in <-chan PipelineData, out chan<- PipelineData, vmList string) {
	vms := strings.Fields(vmList) // convert ListAsString to slice
	nbWorker := len(vms)
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

			// Update the VM OS based on its osFamily
			cmd := fmt.Sprintf("luc util osupgrade")
			logx.L.Debugf("⚠️ will play loccally on the remote VM: %s", cmd)
			// cmd := fmt.Sprintf("pwd; ls -ial")
			// result, err := util.RunCLIRemote(cmd, data.HostName)
			// if err != nil {
			// 	data.Err = err
			// 	logx.L.Debugf("❌ error detected")
			// }
			// logx.L.Debugf("⚠️ result: %s", result)

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
