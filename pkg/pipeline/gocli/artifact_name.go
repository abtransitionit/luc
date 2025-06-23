/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package gocli

import (
	"path"

	"github.com/abtransitionit/luc/pkg/logx"
)

func ArtifactName(in <-chan PipelineData, out chan<- PipelineData) {
	defer close(out)

	for data := range in {
		if data.Err != nil {
			out <- data
			logx.L.Debugf("❌ Previous error detected")
			continue
		}

		// set this instance property
		data.ArtifactName = path.Base(data.SpecificUrl)

		// send
		out <- data
	}
}
