/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package gocli

import (
	"fmt"
	"path"
	"path/filepath"
	"time"

	"github.com/abtransitionit/luc/pkg/logx"
)

func setArtifact(in <-chan PipelineData, out chan<- PipelineData) {
	defer close(out)

	for data := range in {
		if data.Err != nil {
			out <- data
			logx.L.Debugf("❌ Previous error detected")
			continue
		}

		// get instance property
		url := data.AppUrl
		// set instance property
		data.ArtifactName = path.Base(url)
		uniquePath := filepath.Join("/tmp", fmt.Sprintf("%s_%d", data.ArtifactName, time.Now().UnixNano()))
		data.ArtifactPath = uniquePath

		// log information
		logx.L.Debugf("[%s] setted artifact property", data.Config.Name)
		// send
		out <- data
	}
}
