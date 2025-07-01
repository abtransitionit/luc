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
	defer close(out)

	for data := range in {
		if data.Err != nil {
			out <- data
			logx.L.Debugf("❌ Previous error detected")
			continue
		}

		// set this instance property - temporary folder whare to save the artifact
		uniquePath := filepath.Join("/tmp", fmt.Sprintf("%s_%d", data.ArtifactName, time.Now().UnixNano()))
		data.ArtifactPath = uniquePath

		// send
		out <- data
	}
}
