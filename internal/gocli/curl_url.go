/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package gocli

import (
	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/abtransitionit/luc/pkg/util"
)

func CurlUrl(in <-chan PipelineData, out chan<- PipelineData) {
	go func() {
		// close channel
		defer close(out)

		// get config for this CLI - Did something gets wrong earlier
		for data := range in {
			if data.Err != nil {
				// send data to next step
				out <- data
				// Keep reading data from channel
				continue
			}
			// curl the File into memory
			memoryFile, err := util.GetPublicFile(logx.L, data.SpecificUrl)
			if err != nil {
				data.Err = err
			} else {
				// define this property
				data.MemoryFile = memoryFile
				logx.L.Infof("File Downloaded into Memory")
			}

			// send data to next step
			out <- data
		}
	}()
}

// logx.L.Infow("Downloaded file into memory", "cli", data.Config.Name, "size", len(fileBytes))
// logx.L.Infow("Downloaded file into memory", "cli", data.Config.Name, "size", len(fileBytes))
// logx.L.Infow("Specific URL generated: '%s'", data.SpecificUrl)
// logx.L.Errorw("Failed to download file", "cli", data.Config.Name, "url", data.CurUrl, "err", err)
// logx.L.Debugf("Failed to download file", "cli", data.Config.Name, "url", data.CurUrl, "err", err)
// logx.L.Infow("Specific URL generated: '%s'", data.SpecificUrl)
