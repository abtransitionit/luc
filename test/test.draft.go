/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package test

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/abtransitionit/luc/pkg/action"
	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/abtransitionit/luc/pkg/util"
	"github.com/abtransitionit/luc/pkg/util/dnfapt"
)

func GetPackage1(vm string, cli string) (string, error) {
	logx.L.Debugf("GePackage for cli : %s on linux distro : default", cli)

	return "", nil
}
func CheckCliExits(cli string) (string, error) {
	logx.L.Debugf("Check cli : %s exists", cli)

	return "", nil
}
func GetPackage2(vm string, osDistro string, cli string) (string, error) {
	logx.L.Debugf("GePackage for cli : %s on linux distro : %s", cli, osDistro)
	return "", nil
}
func GeRemoteProperty(vm string, property string) string {
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
	result, err := action.PlayFnLocally("CheckFileExists", fnParameters)

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

	// do the check
	result, err := action.PlayFnOnRemote(vm, "CheckFileExists", fnParameters)

	// error
	if err != nil {
		logx.L.Debugf("%v : %s", err, result)
		return false
	}

	// normalize result
	res := strings.ToLower(strings.TrimSpace(result))

	switch res {
	case "false":
		logx.L.Debugf("❌ [%s] [%s] remote file does not exist", vm, fullPath)
		return false
	case "true":
		logx.L.Debugf("✅ [%s] [%s] remote cli exists", vm, fullPath)
		return true
	default:
		logx.L.Infof("⚠️ ⚠️ [%s] [%s] Impossible to say, result unknown: -%s-", vm, fullPath, result)
		return false
	}
}

func DaAddRepoLocal(repoName string) bool {
	result, err := dnfapt.AddRepo(repoName)
	if err != nil {
		logx.L.Debugf("%v : %s", err, result)
		return false
	}
	return true
}

func CheckCliExistsOnremote(vm string, cliName string) bool {

	// Convert string to slice
	fnParameters := []string{cliName}

	// check vm is ssh reachable
	result, err := util.GetPropertyLocal("sshreachability", vm)
	if err != nil {
		logx.L.Debugf("%v : %s", err, result)
		return false
	} else if result != "true" {
		logx.L.Debugf("❌ [%s] : %s", vm, "not reachable")
		return false
	}

	// do the check
	// logx.L.Infof("[%s] [%s] Checking cli exists", vm, cliName)
	result, err = action.PlayFnOnRemote(vm, "CheckCliExists", fnParameters)

	// error
	if err != nil {
		logx.L.Debugf("%v : %s", err, result)
		return false
	}

	// normalize result
	res := strings.ToLower(strings.TrimSpace(result))

	switch res {
	case "false":
		logx.L.Debugf("❌ [%s] [%s] : %s", vm, cliName, result)
		return false
	case "true":
		logx.L.Debugf("✅ [%s] [%s] : %s", vm, cliName, result)
		return true
	default:
		logx.L.Infof("⚠️ ⚠️ [%s] [%s] : %s", vm, cliName, result)
		return false
	}
}

func CheckVmIsSshReachable(vm string) bool {

	// do the check
	// logx.L.Infof("[%s] Checking vm is ssh reachable", vm)
	result, err := util.GetPropertyLocal("sshreachability", vm)

	// error
	if err != nil {
		logx.L.Debugf("%v : %s", err, result)
		return false
	}

	// normalize result
	res := strings.ToLower(strings.TrimSpace(result))

	switch res {
	case "false":
		logx.L.Debugf("❌ [%s] :%s", vm, result)
		return false
	case "true":
		logx.L.Debugf("✅ [%s] : %s", vm, result)
		return true
	default:
		logx.L.Infof("⚠️ ⚠️ [%s] [%s] Impossible to say", vm, result)
		return false
	}
}

func TestCheckCliExistsOnremote(listVmAsString string, cliName string) {
	for _, vm := range util.GetSlicefromStringWithSpace(listVmAsString) {
		CheckCliExistsOnremote(vm, cliName)
	}
}
func TestVmAreSshReachable(listVmAsString string) {
	for _, vm := range util.GetSlicefromStringWithSpace(listVmAsString) {
		result, err := util.GetPropertyLocal("sshreachability", vm)
		if err != nil {
			logx.L.Debugf("%v : %s", err, result)
			return
		}
		logx.L.Infof("[%s] is ssh reachable: %s", vm, result)
	}

}

func TestGetGpgFromUrl(url string, filePath string, isRootPath bool) {
	savedPath, err := util.GetGpgFromUrl(url, filePath, isRootPath)
	if err != nil {
		logx.L.Debugf("Failed to save GPG key to %s: %v", savedPath, err)
		return
	}
	logx.L.Debugf("GPG key successfully saved to %s", savedPath)
}

func TestRemoteGetGpgFromUrl(vm string, url string, filePath string, isRootPath bool) {
	cli := fmt.Sprintf(`luc do GetGpgFromUrl %s %s %v`, url, filePath, isRootPath)
	savedPath, err := util.RunCLIRemote(vm, cli)
	if err != nil {
		logx.L.Debugf("Failed to save GPG key to %s: %v", filePath, err)
		return
	}
	logx.L.Debugf("GPG key successfully saved to %s", savedPath)
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
