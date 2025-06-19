/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package gocli

import (
	"github.com/abtransitionit/luc/pkg/config"
	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/abtransitionit/luc/pkg/util"
)

// Move File or folder to final destination
func BuildPath(in <-chan PipelineData, out chan<- PipelineData) {
	defer close(out)
	for data := range in {
		if data.Err != nil {
			out <- data
			logx.L.Debugf("❌ Previous error detected")
			continue
		}

		// step 2: Define the PATH to export
		switch data.Config.UrlType {
		case config.UrlTgz, config.UrlExe:
			rootFolder := "/usr/local/bin"
			path, err := util.BuildPath(rootFolder)
			if err != nil {
				logx.L.Debugf("❌ Failed to move file: %s", err)
			}
			logx.L.Debugf("[%s] export PATH=%s", data.Config.Name, path+":$PATH")

		default:
			logx.L.Debugf("[%s] UrlType '%s' is not impact by this stage", data.ArtifactName, data.Config.UrlType)
			out <- data
			continue
		}

		// step 3: send pipeline var to next pipeline step
		out <- data
	}
}
