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

func GeRemotePropertyTest(vm string, property string) string {
	// get some OS property
	osfamily, err := util.GetPropertyRemote(vm, "osfamily")
	if err != nil {
		logx.L.Debugf("❌ output:  %v", osfamily)
		logx.L.Debugf("❌ error:   %v", err)
	}
	return osfamily

}

// locally create an empty file as current user
func TouchFileLocal(folderPath string, fileName string) error {
	filePath := filepath.Join(folderPath, fileName)

	// touch file
	cli := fmt.Sprintf(`touch %s`, filePath)
	_, err := util.RunCLILocal(cli)
	if err != nil {
		logx.L.Debugf("%s", err)
		return fmt.Errorf("failed to touch file locally: %w", err)
	}
	logx.L.Infof("✅ touched local file without garantee: %s", filePath)
	return nil
}

func DeleteFileLocal(folderPath string, fileName string) error {
	filePath := filepath.Join(folderPath, fileName)

	// delete file
	cli := fmt.Sprintf(`rm %s`, filePath)
	_, err := util.RunCLILocal(cli)
	if err != nil {
		logx.L.Debugf("%s", err)
		return fmt.Errorf("failed to delete file locally: %w", err)
	}
	logx.L.Infof("✅ deleted local file without garantee: %s", filePath)
	return nil
}

func TouchFileOnRemote(vm string, folderPath string, fileName string) error {
	filePath := filepath.Join(folderPath, fileName)

	// touch file
	cli := fmt.Sprintf(`touch %s`, filePath)
	_, err := util.RunCLIRemote(vm, cli)
	if err != nil {
		logx.L.Debugf("%s", err)
		return fmt.Errorf("failed to touch file remotely: %w", err)
	}
	logx.L.Infof("✅ touched remote file without garantee: %s", filePath)
	return nil
}

func DeleteFileOnRemote(vm string, folderPath string, fileName string) error {
	filePath := filepath.Join(folderPath, fileName)

	// delete
	cli := fmt.Sprintf(`rm %s`, filePath)
	_, err := util.RunCLIRemote(vm, cli)
	if err != nil {
		logx.L.Debugf("%s", err)
		return fmt.Errorf("failed to delete file remotely: %w", err)
	}
	logx.L.Infof("✅ deleted remote file without garantee: %s", filePath)
	return nil
}

func CheckFileLocalExits(fullPath string) bool {
	// Convert string to slice
	fnParameters := []string{fullPath}
	// check file exists
	result, err := util.PlayFnLocally("CheckFileExists", fnParameters)

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

func CheckFileRemoteExists(vm string, fullPath string) bool {
	// Convert string to slice
	fnParameters := []string{fullPath}
	// check file exists
	result, err := util.PlayFnOnRemote(vm, "CheckFileExists", fnParameters)

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

// func touchFileRemote(vm string) bool {
// 	// define var
// 	srcFileName := "test.txt"
// 	srcFolderName := "/tmp"
// 	srcFilePath := filepath.Join(srcFolderName, srcFileName)

// 	// create empty file
// 	cli := fmt.Sprintf(`touch %s`, srcFilePath)
// 	_, err := util.RunCLIRemote(vm, cli)
// 	if err != nil {
// 		logx.L.Debugf("%s", err)
// 		return false
// 	}

// 	// check file exists
// 	cli = fmt.Sprintf(`luc action FileExists %s`, srcFilePath)
// 	result, err := util.RunCLIRemote(vm, cli)
// 	if err != nil {
// 		logx.L.Debugf("%s", err)
// 		return false
// 	}

// 	// success
// 	logx.L.Info("✅ touched remote file")
// 	return strings.ToLower(strings.TrimSpace(result)) == "true"
// }
