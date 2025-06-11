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

func GuessFileType(in <-chan PipelineData, out chan<- PipelineData) {
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

			// Guess file type from memory content
			fileInMemory := data.MemoryFile

			isGzip, err := util.IsGzippedMemoryContent(fileInMemory)
			if err != nil {
				data.Err = fmt.Errorf("error checking gzip content: %w", err)
				out <- data
				continue
			}

			isExe, err := util.IsMemoryContentAnExe(fileInMemory)
			if err != nil {
				data.Err = fmt.Errorf("error checking exe content: %w", err)
				out <- data
				continue
			}

			switch {
			case isGzip:
				data.ArtifactType = string(config.UrlTgz)
			case isExe:
				data.ArtifactType = string(config.UrlExe)
			default:
				data.ArtifactType = string(config.UrlXxx)
			}

			logx.L.Infof("Artifact File type is '%s'", data.ArtifactType)

			// send data to next step
			out <- data
		}
	}()
}

// logx.L.Infow("Downloaded file into memory", "cli", data.Config.Name, "size", len(fileBytes))
// logx.L.Infow("Downloaded file into memory", "cli", data.Config.Name, "size", len(fileBytes))
// logx.L.Infow("Specific URL generated: '%s'", data.SpecificUrl)
// logx.L.Errorw("Failed to download file", "cli", data.Config.Name, "url", data.CurUrl, "err", err)
// logx.L.Debugf("Failed to download file", "cli", data.Config.Name, "url", data.CurUrl, "err", err)
// logx.L.Infow("Specific URL generated: '%s'", data.SpecificUrl)
