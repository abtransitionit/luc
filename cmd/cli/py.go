/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package cli

import (
	"github.com/spf13/cobra"
)

// Description
var pyLDesc = pySDesc + ` xxx.`
var pySDesc = "manage python packages."

// root Command
var pyCmd = &cobra.Command{
	Use:   "py",
	Short: pySDesc,
	Long:  pyLDesc,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	pyCmd.Flags().BoolP("show", "s", false, "show CLI config map")
}
