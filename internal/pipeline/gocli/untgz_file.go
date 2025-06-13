/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package gocli

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/abtransitionit/luc/pkg/config"
	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/abtransitionit/luc/pkg/util"
)

func UnTgzFile(in <-chan PipelineData, out chan<- PipelineData) {
	go func() {
		// close channel
		defer close(out)

		// get config for this CLI - Did something gets wrong earlier
		for data := range in {
			if data.Err != nil {
				out <- data
				continue
			}
			// Do the work
			processedData := helperUnzip(data)

			// send data to next step
			out <- processedData
		}
	}()
}

func helperUnzip(data PipelineData) PipelineData {

	// manage decompression (based on UrlType)
	switch data.Config.UrlType {
	case config.UrlExe, config.UrlGo, config.UrlGit, config.UrlXxx:
		logx.L.Debugf("File type '%s' does not need decompression", data.Config.UrlType)

		// define property
		data.FofTmpPath = data.ArtifactPath

	case config.UrlTgz:
		logx.L.Infof("Decompressing '%s' (type: %s)", data.ArtifactPath, data.Config.UrlType)

		// define uniquePath
		ArtefactTmpPath := filepath.Join("/tmp", fmt.Sprintf("%s_%d", data.ArtifactName, time.Now().UnixNano()))

		// define property
		data.FofTmpPath = ArtefactTmpPath

		// decompress artifactt
		err := util.UntargzFile(data.ArtifactPath, ArtefactTmpPath)
		if err != nil {
			data.Err = fmt.Errorf("failed to decompress file: %w", err)
		} else {
			logx.L.Infof("✅ Successfully decompressed into '%s", ArtefactTmpPath)
		}

	default:
		logx.L.Infof("⚠️ UrlType '%s' — no decompression rule defined", data.Config.UrlType)
	}

	return data
}
