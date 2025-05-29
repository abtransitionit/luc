/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package gox

import (
	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/spf13/cobra"
)

// Description
var runSDesc = "run a GO project (ie. without building it)."
var runLDesc = runSDesc + ` xxx.`

// root Command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: runSDesc,
	Long:  runLDesc,
	Run: func(cmd *cobra.Command, args []string) {
		logx.L.Infof("%s", runSDesc)
	},
}

func init() {
	runCmd.Flags().StringVarP(&pathFlag, "path", "p", "", "Path to the go project folder (containing the file go.mod)")
	runCmd.Flags().StringVar(&archFlag, "arch", "", "The cpu arch: must be arm or amd")
	runCmd.Flags().StringVar(&osFlag, "os", "", "The OS type: must be linux or darwin")
	runCmd.Flags().StringVar(&outFlag, "out", "", "Path of the output binary")
}
