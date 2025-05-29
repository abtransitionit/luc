/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package gox

import (
	"github.com/abtransitionit/luc/cmd/os/gox/cli"
	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/spf13/cobra"
)

// Description
var goSDesc = "Manage GO applications."
var goLDesc = goSDesc + ` xxx.`

// root Command
var GoCmd = &cobra.Command{
	Use:   "go",
	Short: goSDesc,
	Long:  goLDesc,
	Run: func(cmd *cobra.Command, args []string) {
		logx.L.Infof("%s", goSDesc)
	},
}

func init() {
	GoCmd.AddCommand(cli.CliCmd)
	GoCmd.AddCommand(buildCmd)
}
