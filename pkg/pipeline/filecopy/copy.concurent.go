/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package filecopy

import (
	"fmt"
	"sync"

	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/abtransitionit/luc/pkg/util"
)

// # Purpose
//
// - This step performs the actual copy operation
// - Multiple goroutines (workers) do the copy concurrently
func concurrentlyCopyFile(in <-chan PipelineData, out chan<- PipelineData, nbWorker int) {
	var wg sync.WaitGroup
	defer close(out)

	// Worker function
	worker := func() {
		defer wg.Done()
		for data := range in {
			if data.Err != nil {
				out <- data
				continue
			}
			cmd := fmt.Sprintf("scp %s %s:%s", data.SrcFile, data.Node, data.DstFile)
			_, err := util.RunCLILocal(cmd)
			if err != nil {
				data.Err = err
				logx.L.Debugf("❌ error detected")
			} else {
				logx.L.Infof("[%s] ✅ Copied successfully", data.Node)
			}

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
