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
				// send data to next step
				out <- data
				// Keep reading data from channel
				continue
			}

			// get this property
			fileInMemory := data.MemoryFile

			// classify the artifact
			if isGzip, _ := util.IsGzippedMemoryContent(fileInMemory); isGzip {
				data.ArtifactType = string(config.UrlTgz)
			} else if isExe, _ := util.IsMemoryContentAnExe(fileInMemory); isExe {
				data.ArtifactType = string(config.UrlExe)
			} else {
				data.ArtifactType = string(config.UrlXxx)
				logx.L.Infof("Unknown file type for '%s'; classified as '%s'", data.ArtifactName, data.ArtifactType)
			}

			// log information
			logx.L.Infof("Artifact File type is '%s'", data.ArtifactType)

			// send data to next step
			out <- data
		}
	}()
}

// func GuessFileType(in <-chan PipelineData, out chan<- PipelineData) {
// 	go func() {
// 		// close channel
// 		defer close(out)

// 		for data := range in {
// 			// propagate error if any
// 			if data.Err != nil {
// 				// send data to next step
// 				out <- data
// 				// Keep reading data from channel
// 				continue
// 			}

// 			// get this property
// 			fileInMemory := data.MemoryFile

// 			// do the job
// 			isGzip, err := util.IsGzippedMemoryContent(fileInMemory)
// 			if err != nil {
// 				data.Err = fmt.Errorf("error checking gzip content: %w", err)
// 				out <- data
// 				continue
// 			}

// 			isExe, err := util.IsMemoryContentAnExe(fileInMemory)
// 			if err != nil {
// 				data.Err = fmt.Errorf("error checking exe content: %w", err)
// 				out <- data
// 				continue
// 			}

// 			// define property
// 			switch {
// 			case isGzip:
// 				data.ArtifactType = string(config.UrlTgz)
// 			case isExe:
// 				data.ArtifactType = string(config.UrlExe)
// 			default:
// 				data.ArtifactType = string(config.UrlXxx)
// 			}

// 			// log information
// 			logx.L.Infof("Artifact File type is '%s'", data.ArtifactType)

// 			// send data to next step
// 			out <- data
// 		}
// 	}()
// }
