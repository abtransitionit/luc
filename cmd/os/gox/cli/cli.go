/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package cli

import (
	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/spf13/cobra"
)

// Description
var cliSDesc = "Manage GO applications."
var cliLDesc = cliSDesc + ` xxx.`

// root Command
var CliCmd = &cobra.Command{
	Use:   "cli",
	Short: cliSDesc,
	Long:  cliLDesc,
	Run: func(cmd *cobra.Command, args []string) {
		logx.L.Infof("%s", cliSDesc)
	},
}

func init() {
	CliCmd.AddCommand(cobraCmd)
	CliCmd.AddCommand(sdkCmd)
}
