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
		// display if flag = --list
		if cmd.Flag("list").Value.String() == "true" {
			config.DisplayCliCondfigInfo()
			return
		}
		cmd.Help()
	},
}

func init() {
	CliCmd.AddCommand(cobraCmd)
	CliCmd.AddCommand(containerdCmd)
	CliCmd.AddCommand(helmCmd)
	CliCmd.AddCommand(helm2Cmd)
	CliCmd.AddCommand(kindCmd)
	CliCmd.AddCommand(kubebuilderCmd)
	CliCmd.AddCommand(kubectlCmd)
	CliCmd.AddCommand(nerdctlCmd)
	CliCmd.AddCommand(rootlesskitCmd)
	CliCmd.AddCommand(runcCmd)
	CliCmd.AddCommand(sdkCmd)
	CliCmd.AddCommand(sonobuoyCmd)
	//
	CliCmd.Flags().BoolP("list", "l", false, "List CLI configurations")

}
