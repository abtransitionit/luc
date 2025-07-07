/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package kind

import (
	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/abtransitionit/luc/pkg/util"
)

const RcDescription = "configure RC file."

func rc(arg ...string) (string, error) {
	logx.L.Info(RcDescription)
	// define var
	CustomRcFile := "$HOME/.profile.luc"
	RcFilePath := "$HOME/.bashrc"
	line := "source " + CustomRcFile
	// add line
	logx.L.Info("adding line to rc file")
	err := util.AddLineRcFile(RcFilePath, line)
	if err != nil {
		return "", err
	}
	// build a tree PATH
	basePath := "/usr/local/bin"
	logx.L.Debugf("building tree path from: '%s'", basePath)
	treePath, err := util.BuildPath(basePath)
	if err != nil {
		logx.L.Debugf("❌ Error detected 7")
		return "", err
	}

	// update PATH with this tree path
	updatedPath, err := util.UpdPath(treePath)
	if err != nil {
		logx.L.Debugf("❌ Error detected 8")
		return "", err
	}
	// add instruction
	logx.L.Info("adding instruction to rc file")
	line = "export PATH=" + updatedPath
	err = util.AddLineRcFile(CustomRcFile, line)
	if err != nil {
		return "", err
	}
	// add instruction
	logx.L.Info("adding instruction to rc file")
	line = "export CNI_PATH=/usr/local/bin/cni"
	err = util.AddLineRcFile(CustomRcFile, line)
	if err != nil {
		return "", err
	}

	return "", nil
}
