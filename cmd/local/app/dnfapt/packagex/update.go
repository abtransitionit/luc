/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package packagex

import (
	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/spf13/cobra"
)

// Description
var updateSDesc = "update OS dnfapt package."
var updateLDesc = updateSDesc + ` xxx.`

// root Command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: updateSDesc,
	Long:  updateLDesc,
	Run: func(cmd *cobra.Command, args []string) {
		logx.L.Debugf(updateSDesc)
	},
}
