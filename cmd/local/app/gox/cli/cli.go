/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package cli

import (
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
}
