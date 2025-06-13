/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package gocli

import (
	"github.com/abtransitionit/luc/pkg/config"
	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/abtransitionit/luc/pkg/util"
)

func GuessFileType(in <-chan PipelineData, out chan<- PipelineData) {
	go func() {
		// close channel
		defer close(out)

		for data := range in {
			// propagate error if any
			if data.Err != nil {
				out <- data
				continue
			}

			// get property
			fileInMemory := data.MemoryFile

			// classify the artifact
			if isGzip, _ := util.IsGzippedMemoryContent(fileInMemory); isGzip {
				data.ArtifactType = string(config.UrlTgz)
			} else if isExe, _ := util.IsMemoryContentAnExe(fileInMemory); isExe {
				data.ArtifactType = string(config.UrlExe)
			} else {
				logx.L.Debugf("file type for '%s' is not guesssed", data.ArtifactName)
			}

			// log information
			logx.L.Debugf("Artifact gueessed File type is '%s'", data.ArtifactType)

			// send data to next step
			out <- data
		}
	}()
}
