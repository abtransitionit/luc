/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package packagex

import (
	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/spf13/cobra"
)

// Description
var addSDesc = "add dnfapt package."
var addLDesc = addSDesc + ` xxx.`

// root Command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: addSDesc,
	Long:  addLDesc,
	Run: func(cmd *cobra.Command, args []string) {
		logx.L.Debugf(addSDesc)
	},
}
