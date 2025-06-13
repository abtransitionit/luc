/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package gocli

import (
	"path"

	"github.com/abtransitionit/luc/pkg/logx"
)

func ArtifactName(in <-chan PipelineData, out chan<- PipelineData) {
	go func() {
		// close channel
		defer close(out)

		for data := range in {
			// Step 1: propagate error if any
			if data.Err != nil {
				out <- data
				continue
			}

			// define property
			data.ArtifactName = path.Base(data.SpecificUrl)

			// log information
			logx.L.Infof("Artifact Name: '%s'", data.ArtifactName)

			// send data to next step
			out <- data
		}
	}()
}

// logx.L.Infow("Specific URL generated", "cli", data.Config.Name, "specificUrl", data.SpecificUrl)
// logx.L.Infow("Specific URL generated: '%s'", data.SpecificUrl)
