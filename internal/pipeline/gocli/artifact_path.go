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

		for data := range in {
			// Step 1: propagate error if any
			if data.Err != nil {
				out <- data
				continue
			}

			// step 2: define property
			uniquePath := filepath.Join("/tmp", fmt.Sprintf("%s_%d", data.ArtifactName, time.Now().UnixNano()))
			data.ArtifactPath = uniquePath

			// log information
			logx.L.Infof("Artifact Path: '%s'", data.ArtifactPath)

			// step 3: send pipeline var to next pipeline step
			out <- data
		}
	}()
}

// logx.L.Infow("Specific URL generated", "cli", data.Config.Name, "specificUrl", data.SpecificUrl)
