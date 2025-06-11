/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package gocli

import (
	"fmt"

	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/abtransitionit/luc/pkg/util"
)

func SaveFile(in <-chan PipelineData, out chan<- PipelineData) {
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

			// Save the file from memory
			_, err := util.SaveToFile(logx.L, data.ArtifactPath, data.MemoryFile)
			if err != nil {
				data.Err = fmt.Errorf("failed to save file: %w", err)
				out <- data
				continue
			}

			logx.L.Infof("File saved to '%s'", data.ArtifactPath)

			// send data to next step
			out <- data
		}
	}()
}
