/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package gox

import (
	"log"

	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/abtransitionit/luc/pkg/util"
	"github.com/spf13/cobra"
)

// Description
var runSDesc = "run a GO project for test (ie. without building it)."
var runLDesc = runSDesc + ` xxx.`

// root Command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: runSDesc,
	Long:  runLDesc,
	Run: func(cmd *cobra.Command, args []string) {
		logx.L.Debugf(runSDesc)

		// prerequisit: check go CLI exists
		cliName := "go"
		if cliExist, err := util.CliExists(cliName); err != nil {
			log.Fatalf("CLI not exist: %s : %v", cliName, err)
			return
		} else if !cliExist {
			log.Println("Go toolchain not found - please install Go first")
		}
		logx.L.Debugf("CLI: %s exists", cliName)
		// get project Path
		return
		// check flags
		folerPath := "/var/tmp"
		bool, err := util.FolderExists(folerPath)
		if err != nil {
			logx.L.Errorf("❌ Error checking folder: %s > %v", folerPath, err)
			return
		} else if !bool {
			logx.L.Errorf("❌ Error: folder %s does not exist", folerPath)
			return
		} else {
			logx.L.Infof("folder : %s exists", folerPath)
		}
	},
}

func init() {
	runCmd.Flags().StringVarP(&pathFlag, "path", "p", "", "Path to the go project folder (containing the file go.mod/main.go)")
	runCmd.Flags().StringVar(&archFlag, "arch", "", "The cpu arch: must be arm or amd")
	runCmd.Flags().StringVar(&osFlag, "os", "", "The OS type: must be linux or darwin")
	runCmd.Flags().StringVar(&outFlag, "out", "", "Path of the output binary")
	runCmd.MarkFlagRequired("path") // This makes the flag mandatory
}
