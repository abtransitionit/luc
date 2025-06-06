/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package gox

import (
	"github.com/abtransitionit/luc/cmd/local/app/gox/cli"
	"github.com/spf13/cobra"
)

// Description
var goSDesc = "Manage GO applications and toolchain."
var goLDesc = goSDesc + ` xxx.`

// root Command
var GoCmd = &cobra.Command{
	Use:   "go",
	Short: goSDesc,
	Long:  goLDesc,
}

func init() {
	GoCmd.AddCommand(cli.CliCmd)
	GoCmd.AddCommand(buildCmd)
	GoCmd.AddCommand(runCmd)
	GoCmd.AddCommand(toolCmd)
}
