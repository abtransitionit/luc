/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package test

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/abtransitionit/luc/pkg/util"
)

func checkFileRemoteExists(vm string, fullPath string) bool {
	// Convert string to slice
	fnParameters := []string{fullPath}
	// check file exists
	result, err := util.PlayFnOnRemote(vm, "IsFileExists", fnParameters)

	// error
	if err != nil {
		logx.L.Debugf("%s", err)
		return false
	}

	// normalize result
	res := strings.ToLower(strings.TrimSpace(result))

	switch res {
	case "false":
		logx.L.Debugf("❌ remote file does not exist: %s", fullPath)
		return false
	case "true":
		logx.L.Debugf("✅ remote file exists: %s", fullPath)
		return true
	default:
		logx.L.Infof("⚠️ ⚠️ Impossible to say, result unknown: -%s-", result)
		return false
	}
}

func touchFileRemote(vm string) bool {
	// define var
	srcFileName := "test.txt"
	srcFolderName := "/tmp"
	srcFilePath := filepath.Join(srcFolderName, srcFileName)

	// create empty file
	cli := fmt.Sprintf(`touch %s`, srcFilePath)
	_, err := util.RunCLIRemote(vm, cli)
	if err != nil {
		logx.L.Debugf("%s", err)
		return false
	}

	// check file exists
	cli = fmt.Sprintf(`luc action FileExists %s`, srcFilePath)
	result, err := util.RunCLIRemote(vm, cli)
	if err != nil {
		logx.L.Debugf("%s", err)
		return false
	}

	// success
	logx.L.Info("✅ touched remote file")
	return strings.ToLower(strings.TrimSpace(result)) == "true"
}
