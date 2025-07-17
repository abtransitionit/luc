/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package test

import (
	"fmt"
	"path/filepath"

	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/abtransitionit/luc/pkg/util"
	"github.com/spf13/cobra"
)

// Description
var unitSDesc = "Play go unit test:"
var unitLDesc = unitSDesc + `
- launch a bash script that manage the sequence of test to run

Example Usage:

 luc test unit rem				// play all test in folder remote
 luc test unit local			// play all test in folder local
 luc test unit rem Touch  // play [function]Test starting with "Touch" in remote test
`

// root Command
var unitCmd = &cobra.Command{
	Use:   "unit",
	Short: unitSDesc,
	Long:  unitLDesc,
	Run: func(cmd *cobra.Command, args []string) {

		// define var
		TestFolderName := "test"
		bashScriptName := "go.sh"

		// get gomod path
		cli := "go env GOMOD"
		gomodPath, err := util.RunCLILocal(cli)
		if err != nil {
			logx.L.Error(err)
			return
		}

		// get gomodFolder
		gomodFolder := filepath.Dir(gomodPath)

		// create folder path
		bashScriptPath := filepath.Join(gomodFolder, TestFolderName, bashScriptName)

		// manage argument if exists
		scriptParams := ""
		if len(args) > 0 {
			scriptParams = util.GetStringfromSliceWithSpace(args)
		}

		// play cli
		logx.L.Infof("play bash script run test: %s", bashScriptPath)
		cli = fmt.Sprintf(`%s %s`, bashScriptPath, scriptParams)
		out, err := util.RunCLILocal(cli, true)

		// error
		if err != nil {
			logx.L.Debugf("%s, %s", err, out)
			return
		}

		// success
		println(out)
	},
}

var forceFlag bool

func init() {
	unitCmd.Flags().BoolVarP(&forceFlag, "force", "f", false, "Bypass confirmation")
	unitCmd.Flags().BoolP("list", "l", false, "List all available phases")
	unitCmd.Flags().BoolP("runall", "r", false, "Run all phases in batch mode")
	// Make them mutually exclusive
	unitCmd.MarkFlagsMutuallyExclusive("list", "runall")
}
