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
		url := data.HostUrl

		// set instance property
		data.ArtName = path.Base(url)
		uniquePath := filepath.Join("/tmp", fmt.Sprintf("%s_%d", data.ArtName, time.Now().UnixNano()))
		data.ArtPath1 = uniquePath

		// log
		logx.L.Debugf("[%s] setted artifact property", data.Config.Name)
		// send
		out <- data
	}
}
