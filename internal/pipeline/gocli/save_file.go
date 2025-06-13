/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package gocli

import (
	"fmt"

	"github.com/abtransitionit/luc/pkg/config"
	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/abtransitionit/luc/pkg/util"
)

func SaveFile(in <-chan PipelineData, out chan<- PipelineData) {
	go func() {
		// close channel
		defer close(out)

		for data := range in {
			// propagate error if any
			if data.Err != nil {
				out <- data
				continue
			}

			// check Artifact need to be saved (based on UrlType)
			switch data.Config.UrlType {
			case config.UrlExe, config.UrlTgz:
				// proceed
			default:
				logx.L.Debugf("UrlType '%s' does not require saving", data.Config.UrlType)
				out <- data
				continue
			}

			// Check: ensure path is defined
			if data.ArtifactPath == "" {
				data.Err = fmt.Errorf("ArtifactPath is empty, cannot save")
				out <- data
				continue
			}

			// Save the file from memory into host FS
			_, err := util.SaveToFile(logx.L, data.ArtifactPath, data.MemoryFile)
			if err != nil {
				data.Err = fmt.Errorf("failed to save file: %w", err)
				out <- data
				continue
			}

			// log information
			logx.L.Debugf("File saved to '%s'", data.ArtifactPath)

			// step 3: send pipeline var to next pipeline step
			out <- data
		}
	}()
}
