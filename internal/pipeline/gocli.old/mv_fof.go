/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package gocliold

import (
	"path"

	"github.com/abtransitionit/luc/pkg/config"
	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/abtransitionit/luc/pkg/util"
)

// Move File or folder to final destination
func MvFof(in <-chan PipelineData, out chan<- PipelineData) {
	go func() {
		defer close(out) // close channel when done
		// loop over each item of type PipelineData in the channel
		for data := range in {
			// Step 1: propagate error if any
			if data.Err != nil {
				out <- data
				continue
			}

			// step 2: Move file or folder to final destination (based on UrlType)
			switch data.Config.UrlType {
			case config.UrlTgz:
				dstFolder := "/usr/local/bin"
				dstPath := path.Join(dstFolder, data.Config.Name)
				logx.L.Debugf("Moving '%s' to '%s'", data.FofTmpPath, dstPath)
				success, err := util.MvFolder(data.FofTmpPath, dstPath, 0755, true, true)
				if err != nil {
					logx.L.Debugf("❌ Failed to move folder: %s", err)
				}
				if success {
					logx.L.Infof("✅ Folder moved successfully to '%s'", dstPath)
				}

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

			/// step 3: send pipeline var to next pipeline step
			out <- data
		}
	}()
}

// logx.L.Infow("Specific URL generated", "cli", data.Config.Name, "specificUrl", data.SpecificUrl)
