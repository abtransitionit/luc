/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package gocli

import (
	"fmt"

	"github.com/abtransitionit/luc/pkg/config"
	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/abtransitionit/luc/pkg/util"
)

func FileSave(in <-chan PipelineData, out chan<- PipelineData) {
	defer close(out)
	for data := range in {
		if data.Err != nil {
			out <- data
			logx.L.Debugf("❌ Previous error detected")
			continue
		}

		// check Artifact is impacted by this stage (based on UrlType)
		switch data.Config.UrlType {
		case config.UrlExe, config.UrlTgz:
			// proceed
		default:
			logx.L.Debugf("UrlType '%s' is not impact by this stage", data.Config.UrlType)
			out <- data
			continue
		}

		// Check path is defined
		if data.ArtifactPath == "" {
			data.Err = fmt.Errorf("ArtifactPath is empty for cli (%s), cannot save file", data.Config.Name)
			logx.L.Debugf("❌ Error detected 1")
			out <- data
			continue
		}

		// Save the file from memory into host FS
		_, err := util.SaveToFile(logx.L, data.ArtifactPath, data.MemoryFile)
		if err != nil {
			data.Err = fmt.Errorf("failed to save file for cli (%s): %w", data.Config.Name, err)
			logx.L.Debugf("❌ Error detected 2")
			out <- data
			continue
		}

		// send
		out <- data
	}
}
