/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package gocliold

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
		defer close(out) // close channel when done
		// loop over each item of type PipelineData in the channel
		for data := range in {
			// Step2: is there a previous error ? if yes propagate it
			if data.Err != nil {
				out <- data
				continue
			}
			// step 3: do
			processedData := helperUnzip(data)

			// step 4: send pipeline var to next pipeline step
			out <- processedData
		}
	}()
}

func helperUnzip(data PipelineData) PipelineData {

	// manage decompression (based on UrlType)
	switch data.Config.UrlType {
	case config.UrlExe, config.UrlGo, config.UrlGit, config.UrlOth:
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
