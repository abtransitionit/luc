/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package gocliold

import (
	"github.com/abtransitionit/luc/pkg/config"
	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/abtransitionit/luc/pkg/util"
)

// Move File or folder to final destination
func UpdatePath(in <-chan PipelineData, out chan<- PipelineData) {
	go func() {
		defer close(out) // close channel when done
		// loop over each item of type PipelineData in the channel
		for data := range in {
			// Step 1: propagate error if any
			if data.Err != nil {
				out <- data
				continue
			}

			// step 2: Define the PATH to export
			switch data.Config.UrlType {
			case config.UrlTgz, config.UrlExe:
				rootFolder := "/usr/local/bin"
				path, err := util.BuildPath(rootFolder)
				if err != nil {
					logx.L.Debugf("❌ Failed to move file: %s", err)
				}
				logx.L.Debugf("export PATH=" + path + ":$PATH")

			default:
				// log information
				logx.L.Debugf("UrlType '%s' not yet managed", data.Config.UrlType)
				// send data to next step
				out <- data
				// Keep reading data from channel
				continue
			}

			// step 3: send pipeline var to next pipeline step
			out <- data
		}
	}()
}
