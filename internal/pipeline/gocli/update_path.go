/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package gocli

import (
	"github.com/abtransitionit/luc/pkg/config"
	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/abtransitionit/luc/pkg/util"
)

// Move File or folder to final destination
func UpdatePath(in <-chan PipelineData, out chan<- PipelineData) {
	go func() {
		// close channel
		defer close(out)

		for data := range in {
			// Step 1: propagate error if any
			if data.Err != nil {
				out <- data
				continue
			}

			// Step 2: Define the PATH to export
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

			// Step 3: send result to next pipeline step
			out <- data
		}
	}()
}
