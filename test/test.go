/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package test

import (
	"github.com/abtransitionit/luc/pkg/logx"
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

		// // local function currently in test

		// // remote function currently in test

		// local function tested
		touchFileLocal("/tmp", "titi")
		checkFileLocalExits("/tmp/test.txt")
		// remote function tested
		touchFileOnRemote("o1u", "/tmp", "toto")
		checkFileRemoteExists("o1u", "/tmp/toto")

		// TODO

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

func init() {
	TestCmd.Flags().BoolVarP(&forceFlag, "force", "f", false, "Bypass confirmation")
	TestCmd.Flags().BoolP("list", "l", false, "List all available phases")
	TestCmd.Flags().BoolP("runall", "r", false, "Run all phases in batch mode")
	// Make them mutually exclusive
	TestCmd.MarkFlagsMutuallyExclusive("list", "runall")
}
