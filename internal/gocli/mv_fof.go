/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package gocli

import (
	"path"

	"github.com/abtransitionit/luc/pkg/config"
	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/abtransitionit/luc/pkg/util"
)

// Move File or folder to final destination
func MvFof(in <-chan PipelineData, out chan<- PipelineData) {
	go func() {
		// close channel
		defer close(out)

		for data := range in {
			// Step 1: propagate error if any
			if data.Err != nil {
				out <- data
				continue
			}

			// Step 2: Mv file or folder to final destination (based on UrlType)
			switch data.Config.UrlType {
			case config.UrlExe:
				dstFolder := "/usr/local/bin"
				dstPath := path.Join(dstFolder, data.Config.Name)
				logx.L.Debugf("Moving '%s' to '%s'", data.FofTmpPath, dstPath)
				success, err := util.MvFile(data.FofTmpPath, dstPath, 0755, true)
				if err != nil {
					logx.L.Debugf("❌ Failed to move file: %s", err)
				}
				if success {
					logx.L.Infof("✅ File moved successfully to '%s'", dstPath)
				}

			default:
				// log information
				logx.L.Debugf("UrlType '%s' not yet managed", data.Config.UrlType)
				// send data to next step
				out <- data
				// Keep reading data from channel
				continue
			}

			// Step 3: send result to next pipeline step
			out <- data
		}
	}()
}

// logx.L.Infow("Specific URL generated", "cli", data.Config.Name, "specificUrl", data.SpecificUrl)
