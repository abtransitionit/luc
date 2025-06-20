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

func FileUntgz(in <-chan PipelineData, out chan<- PipelineData) {
	defer close(out)

	for data := range in {
		if data.Err != nil {
			out <- data
			logx.L.Debugf("❌ Previous error detected")
			continue
		}

		// check Artifact is impacted by this stage (based on UrlType)
		switch data.Config.UrlType {
		case config.UrlTgz:
			// proceed
		default:
			logx.L.Debugf("[%s] UrlType '%s' is not impact by this stage", data.Config.Name, data.Config.UrlType)
			out <- data
			continue
		}

		// Artifact is impacted by this stage
		processedData := helperUnzip(data)

		// send
		out <- processedData
	}
}

func helperUnzip(data PipelineData) PipelineData {

	// define uniquePath
	ArtifactTmpPath := filepath.Join("/tmp", fmt.Sprintf("%s_%d", data.ArtifactName, time.Now().UnixNano()))

	// define property
	data.FofTmpPath = ArtifactTmpPath

	// decompress artifactt
	err := util.UntargzFile(data.ArtifactPath, ArtifactTmpPath)
	if err != nil {
		data.Err = fmt.Errorf("[%s] failed to decompress file: %w", data.Config.Name, err)
		logx.L.Debugf("❌ Error detected 2")
	} else {
		logx.L.Debugf("[%s] Successfully decompressed into '%s", data.Config.Name, ArtifactTmpPath)
	}

	return data
}

// case config.UrlExe, config.UrlGo, config.UrlGit, config.UrlOth:
// 	logx.L.Debugf("File type '%s' does not need decompression", data.Config.UrlType)

// 	// define property
// 	data.FofTmpPath = data.ArtifactPath
