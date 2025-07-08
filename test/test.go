/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package test

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	configi "github.com/abtransitionit/luc/internal/config"
	utili "github.com/abtransitionit/luc/internal/util"
	"github.com/abtransitionit/luc/pkg/config"
	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/abtransitionit/luc/pkg/util"
	"github.com/spf13/cobra"
)

// Description
var testSDesc = "Test some code."
var testLDesc = testSDesc + ` xxx.`

// root Command
var TestCmd = &cobra.Command{
	Use:   "test",
	Short: testSDesc,
	Long:  testLDesc,
	Run: func(cmd *cobra.Command, args []string) {
		logx.L.Info(testSDesc)

		// checkFileLocalExits("/tmp/test.txt")
		checkFileRemoteExists("o1u", "/tmp/toto")
		// checkFileLocalExists("/tmp/toto")
		// createFileLocal()
		// touchFileRemote("o1u")
		// MoveFileLocal()
		// ListOvhVm()
		// ListMapKey()
		// installGoCli()
		// getPath()
		// fmt.Println(configi.KbeGoCliConfigMap)
		// addLineToFileRemote()

	},
}

var forceFlag bool

// SetupCommonFlags configures flags that are shared across commands
func init() {
	TestCmd.Flags().BoolVarP(&forceFlag, "force", "f", false, "Bypass confirmation")
	TestCmd.Flags().BoolP("list", "l", false, "List all available phases")
	TestCmd.Flags().BoolP("runall", "r", false, "Run all phases in batch mode")
	// Make them mutually exclusive
	TestCmd.MarkFlagsMutuallyExclusive("list", "runall")
}

func checkFileLocalExits(fullPath string) bool {

	// check file exists
	result, err := util.PlayFnLocally("IsFileExists", fullPath)

	// error
	if err != nil {
		logx.L.Debugf("%s", err)
		return false
	}

	// app error
	if strings.ToLower(strings.TrimSpace(result)) == "false" {
		logx.L.Debugf("❌ local file does not exist: %s", fullPath)
		return false
	}
	if strings.ToLower(strings.TrimSpace(result)) == "true" {
		logx.L.Debugf("✅ local file exist: %s", fullPath)
		return true
	}

	//
	logx.L.Infof("⚠️ ⚠️ impossible to say result unknow : %s", result)
	return false
}

func checkFileRemoteExists(vm string, fullPath string) bool {

	// check file exists
	result, err := util.PlayFnOnRemote(vm, "IsFileExists", fullPath)

	// error
	if err != nil {
		logx.L.Debugf("Failed to check remote file exists: %s", err)
		return false
	}

	// app error
	if strings.ToLower(strings.TrimSpace(result)) == "false" {
		logx.L.Debugf("❌ local file does not exist: %s", fullPath)
		return false
	}
	if strings.ToLower(strings.TrimSpace(result)) == "true" {
		logx.L.Debugf("✅ local file exist: %s", fullPath)
		return true
	}

	//
	logx.L.Infof("⚠️ ⚠️ impossible to say result unknow : %s", result)
	return false
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

func createFileLocal() bool {
	// define var
	srcFileName := "test.txt"
	srcFolderName := "/tmp"
	srcFilePath := filepath.Join(srcFolderName, srcFileName)
	content := []byte("hello world")

	// create file
	if err := os.WriteFile(srcFilePath, content, 0644); err != nil {
		logx.L.Debugf("%s", err)
		return false
	}

	// check file exists
	cli := fmt.Sprintf(`luc action FileExists %s`, srcFilePath)
	result, err := util.RunCLILocal(cli)
	// error
	if err != nil {
		logx.L.Debugf("%s", err)
		return false
	}

	// app error
	if strings.ToLower(strings.TrimSpace(result)) == "false" {
		logx.L.Debugf("❌ local file does not exist: %s", srcFilePath)
		return false
	}

	// app success
	logx.L.Info("✅ created local file")
	return strings.ToLower(strings.TrimSpace(result)) == "true"
}

func MoveFileLocal() {
	logx.L.Info("move file")

}

func installGoCli() {

	// define vm
	vm := "o1u"

	// define 1 cli
	cliConfig := config.CustomCLIConfig{
		Name:      "runc",
		Version:   "1.3.0",
		DstFolder: "/usr/local/bin",
	}

	// get vm property
	osFamily, err := util.GetPropertyRemote(vm, "osfamily")
	if err != nil {
		logx.L.Debugf("%s", err)
	}

	// log
	logx.L.Infof("instal go cli on %s:%s", vm, osFamily)
	fmt.Println(cliConfig)

	// install cli(s) on VM
	// gocli.RInstallC(vm, &cliConfig)

}

func ListOvhVm() {
	logx.L.Info("List OVH Vm")
	logx.L.Info("🔹 List : %s", utili.ListOvhVm())
}

func ListMapKey() {
	logx.L.Info("List map:keys")
	// list := []string{"a", "b", "c"}
	listKey := util.GetMapKeys(configi.KindGoCliConfigMap)
	logx.L.Infof("🔹 as slice:      %s", listKey)
	logx.L.Infof("🔹 as StringList: %s", util.GetStringfromSliceWithSpace(listKey))
}

func getPath() {
	logx.L.Info("get path")
	path, err := util.GetSubdirRemote("/usr/local/bin", "o1u")
	if err != nil {
		logx.L.Debugf("%s", err)
	}
	logx.L.Infof("path: %s", path)
}

// test the method locally
func addLineToFile() {
	customRcFilePath := "$HOME/.profile.luc"
	RcFilePath := "$HOME/.bashrc"
	line := "source " + customRcFilePath

	if _, err := util.AddLineToFile(RcFilePath, line); err != nil {
		logx.L.Debugf("Error:", err)
	}

	logx.L.Infof("Line added or already present.")

}

// test local the call via luc util
func addLineToFileLocal() {
	customRcFilePath := "$HOME/.profile.luc"
	RcFilePath := "$HOME/.bashrc"
	line := "source " + customRcFilePath

	cli := fmt.Sprintf(`luc util linefile %q %s --force`, line, RcFilePath)

	if _, err := util.RunCLILocal(cli); err != nil {
		logx.L.Debugf("Error: %s", err)
	}

	logx.L.Infof("Line added or already present.")

}

func addLineToFileRemote() {
	customRcFilePath := "'$HOME/.profile.luc'"
	RcFilePath := "'$HOME/.bashrc'"
	line := fmt.Sprintf("source %s", customRcFilePath)

	cli := fmt.Sprintf(`luc util linefile %q %s --force --remote o1u`, line, RcFilePath)

	// logx.L.Debugf("Running CLI: %s", cli) // For debug

	if _, err := util.RunCLILocal(cli); err != nil {
		logx.L.Debugf("Error: %s", err)
	}

	logx.L.Infof("Line added or already present.")

}

// // check file exists
// cli = fmt.Sprintf(`test -f %s && echo true || echo false`, srcFilePath)
// result, err := util.RunCLIRemote(vm, cli)
// if err != nil {
// 	logx.L.Debugf("Failed to check remote file exists: %s", err)
// 	return false
// }
