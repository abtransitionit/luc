/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package util

import (
	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/spf13/cobra"
)

// Description
var utilSDesc = "manage Kind clusters."
var utilLDesc = utilSDesc + ` xxx.`

// root Command
var UtilCmd = &cobra.Command{
	Hidden: true, // available but not visible
	Use:    "util",
	Short:  utilSDesc,
	Long:   utilLDesc,
	Run: func(cmd *cobra.Command, args []string) {
		logx.L.Infof("%s", utilSDesc)
		// manage argument
		if len(args) == 0 {
			cmd.Help()
			return
		}

	},
}

func init() {
	UtilCmd.AddCommand(getpropCmd)
	UtilCmd.AddCommand(mvfileCmd)
	UtilCmd.AddCommand(ovhCmd)
	UtilCmd.AddCommand(resetCmd)
}
