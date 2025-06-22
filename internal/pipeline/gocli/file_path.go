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

// generate the PATH of the CLI
func BuildPath(in <-chan PipelineData, out chan<- PipelineData) {
	defer close(out)

	for data := range in {
		if data.Err != nil {
			out <- data
			logx.L.Debugf("❌ Previous error detected")
			continue
		}

		// define the PATH to export
		switch data.Config.UrlType {
		case config.UrlTgz, config.UrlExe:
			rootFolder := "/usr/local/bin"
			path, err := util.BuildPath(rootFolder)
			if err != nil {
				data.Err = fmt.Errorf("[%s] unable to build PATH: %w", data.Config.Name, err)
				logx.L.Debugf("❌ Error detected 1")
				out <- data
				continue
			}
			logx.L.Debugf("[%s] export PATH=%s", data.Config.Name, path+":$PATH")

		default:
			logx.L.Debugf("[%s] UrlType '%s' is not impact by this stage", data.ArtifactName, data.Config.UrlType)
			out <- data
			continue
		}

		// send
		out <- data
	}
}
