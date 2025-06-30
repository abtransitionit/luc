/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package cli

import (
	"github.com/spf13/cobra"
)

// Description
var daLDesc = daSDesc + ` xxx.`
var daSDesc = "manage dnfapt packages and repositories."

// root Command
var daCmd = &cobra.Command{
	Use:   "da",
	Short: daSDesc,
	Long:  daLDesc,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	daCmd.Flags().BoolP("show", "s", false, "show CLI config map")
}
