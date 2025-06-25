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
	Use:   "util",
	Short: utilSDesc,
	Long:  utilLDesc,
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
	UtilCmd.AddCommand(ovhCmd)
	UtilCmd.AddCommand(getpropCmd)
	UtilCmd.AddCommand(resetCmd)
}
