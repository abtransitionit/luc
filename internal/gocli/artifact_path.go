/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package gocli

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/abtransitionit/luc/pkg/logx"
)

func ArtifactPath(in <-chan PipelineData, out chan<- PipelineData) {
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

			// define uniquePath
			uniquePath := filepath.Join("/tmp", fmt.Sprintf("%s_%d", data.ArtifactName, time.Now().UnixNano()))

			// define this property
			data.ArtifactPath = uniquePath

			logx.L.Infof("Artifact Path: '%s'", data.ArtifactPath)

			// send data to next step
			out <- data
		}
	}()
}

// logx.L.Infow("Specific URL generated", "cli", data.Config.Name, "specificUrl", data.SpecificUrl)
