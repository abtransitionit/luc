/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package gocli

import (
	"path"

	"github.com/abtransitionit/luc/pkg/config"
	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/abtransitionit/luc/pkg/util"
)

// Move File or folder to final destination
func FileMove(in <-chan PipelineData, out chan<- PipelineData) {
	defer close(out)

	for data := range in {
		if data.Err != nil {
			out <- data
			logx.L.Debugf("❌ Previous error detected")
			continue
		}
		var processedData = PipelineData{}
		// step 2: Move file or folder to final destination (based on UrlType)
		switch data.Config.UrlType {
		case config.UrlTgz:
			processedData = helperMvTgz(data)

		case config.UrlExe:
			processedData = helperMvExe(data)

		default:
			logx.L.Debugf("[%s] UrlType '%s' is not supported or not yet managed", data.ArtifactName, data.Config.UrlType) // logx.L.Debugf("[%s] UrlType '%s' is not impact by this stage", data.ArtifactName, data.Config.UrlType)
			out <- data
			continue
		}

		/// step 3: send pipeline var to next pipeline step
		out <- processedData
	}
}

func helperMvExe(data PipelineData) PipelineData {
	dstFolder := "/usr/local/bin"
	dstPath := path.Join(dstFolder, data.Config.Name)
	logx.L.Debugf("[%s] Moving '%s' to '%s'", data.Config.Name, data.FofTmpPath, dstPath)
	success, err := util.MvFile(data.FofTmpPath, dstPath, 0755, true)
	if err != nil {
		data.Err = err
		logx.L.Debugf("❌ Error detected 4")
		return data
	}
	if success {
		logx.L.Infof("✅ [%s] File moved to '%s'", data.Config.Name, dstPath)
	}
	return data
}

func helperMvTgz(data PipelineData) PipelineData {
	dstFolder := "/usr/local/bin"
	dstPath := path.Join(dstFolder, data.Config.Name)
	logx.L.Debugf("Moving '%s' to '%s'", data.FofTmpPath, dstPath)
	success, err := util.MvFolder(data.FofTmpPath, dstPath, 0755, true, true)
	if err != nil {
		data.Err = err
		logx.L.Debugf("❌ Error detected 4")
		return data
	}
	if success {
		logx.L.Infof("✅ Folder moved successfully to '%s'", dstPath)
	}

	return data
}
