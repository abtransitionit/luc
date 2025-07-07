/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package kind

import (
	"github.com/abtransitionit/luc/internal/config"
	"github.com/abtransitionit/luc/pkg/logx"
	pathp "github.com/abtransitionit/luc/pkg/pipeline/path"
)

const PathDescription = "configure OS PATH envvar."

func path(arg ...string) (string, error) {
	_, err := pathp.RunPipeline(config.KindVm, config.KindPathFile)
	if err != nil {
		logx.L.Debugf("%s", err)
		return "", err
	}
	return "", nil
}

// 	// build a tree PATH

// 	// update PATH with this tree path
// 	logx.L.Debugf("parevious $PATH is  : '%s'", envPath)
// 	logx.L.Debugf("current path tree is: '%s'", treePath)
// 	updatedPath, err := util.UpdPath(treePath)
// 	if err != nil {
// 		logx.L.Debugf("❌ Error detected 2")
// 		return "", err
// 	}
