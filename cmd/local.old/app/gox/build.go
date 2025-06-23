/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package gox

import (
	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/spf13/cobra"
)

var (
	pathFlag string
	archFlag string
	osFlag   string
	outFlag  string
)

// Description
var buildSDesc = "Build a GO project."
var buildLDesc = buildSDesc + ` xxx.`

// root Command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: buildSDesc,
	Long:  buildLDesc,
	Run: func(cmd *cobra.Command, args []string) {
		logx.L.Debugf(buildSDesc)
	},
}

func init() {
	buildCmd.Flags().StringVarP(&pathFlag, "path", "p", "", "Path to the go project folder (containing the file go.mod)")
	buildCmd.Flags().StringVar(&archFlag, "arch", "", "The cpu arch: must be arm or amd")
	buildCmd.Flags().StringVar(&osFlag, "os", "", "The OS type: must be linux or darwin")
	buildCmd.Flags().StringVar(&outFlag, "out", "", "Path of the output binary")
}
