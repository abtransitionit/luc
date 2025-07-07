/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package rfilecopy

import (
	"fmt"
	"path/filepath"
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

			// copy file to /tmp
			cli := fmt.Sprintf("scp %s %s:%s", data.SrcFile, data.Node, data.DstFile)
			_, err := util.RunCLILocal(cli)
			if err != nil {
				data.Err = err
				logx.L.Debugf("❌ error detected")
			}

			// sudo move file with luc EXCEPT if we must move LUC itself
			cliName := filepath.Base(data.DstFile)
			if cliName == "luc" {
				cli = fmt.Sprintf(`sudo mv %s "/usr/local/bin/" && chmod +x /usr/local/bin/luc`, data.DstFile)
			} else {
				cli = fmt.Sprintf(`luc util mvfile %s "/usr/local/bin/" 0755 true`, data.DstFile)
			}
			// play CLI
			_, err = util.RunCLIRemote(data.Node, cli)
			if err != nil {
				data.Err = err
				logx.L.Debugf("❌ error detected")
			}
			// set property
			data.DstFile = fmt.Sprintf("/usr/local/bin/%s", cliName)
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
