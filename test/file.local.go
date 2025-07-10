/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package test

import (
	"strings"

	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/abtransitionit/luc/pkg/util"
)

func checkFileLocalExits(fullPath string) bool {
	// Convert string to slice
	fnParameters := []string{fullPath}
	// check file exists
	result, err := util.PlayFnLocally("IsFileExists", fnParameters)

	// error
	if err != nil {
		logx.L.Debugf("%s", err)
		return false
	}

	// normalize result
	res := strings.ToLower(strings.TrimSpace(result))

	switch res {
	case "false":
		logx.L.Debugf("❌ local file does not exist: %s", fullPath)
		return false
	case "true":
		logx.L.Debugf("✅ local file exists: %s", fullPath)
		return true
	default:
		logx.L.Infof("⚠️ ⚠️ Impossible to say, result unknown: %s", result)
		return false
	}
}
