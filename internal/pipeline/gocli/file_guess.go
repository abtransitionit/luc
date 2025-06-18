/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package gocli

import (
	"github.com/abtransitionit/luc/pkg/config"
	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/abtransitionit/luc/pkg/util"
)

func FileGuessType(in <-chan PipelineData, out chan<- PipelineData) {
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
			data.ArtifactType = "NA"
			// logx.L.Debugf("file type for '%s' is not guesssed", data.ArtifactName)
			logx.L.Debugf("[%s] UrlType '%s' is not impact by this stage", data.ArtifactName, data.Config.UrlType)
			out <- data
			continue
		}

		// get instance property
		fileInMemory := data.MemoryFile

		// step 2: classify the artifact
		if isGzip, _ := util.IsGzippedMemoryContent(fileInMemory); isGzip {
			data.ArtifactType = string(config.UrlTgz)
		} else if isExe, _ := util.IsMemoryContentAnExe(fileInMemory); isExe {
			data.ArtifactType = string(config.UrlExe)
		} else {
			// do nothing
		}

		// send
		out <- data
	}
}
