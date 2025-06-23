/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package cli

import (
	"github.com/abtransitionit/luc/pkg/config"
	"github.com/spf13/cobra"
)

// Description
var cliSDesc = "Manage GO CLI."
var cliLDesc = cliSDesc + ` xxx.`

// root Command
var CliCmd = &cobra.Command{
	Use:   "cli",
	Short: cliSDesc,
	Long:  cliLDesc,
	Run: func(cmd *cobra.Command, args []string) {
		// handle flag = --show
		if cmd.Flag("show").Value.String() == "true" {
			config.ShowCliConfigMap()
			return
		}
		cmd.Help()
	},
}

func init() {
	CliCmd.AddCommand(ovhCmd)
	//
	CliCmd.Flags().BoolP("show", "s", false, "show CLI config map")

}
