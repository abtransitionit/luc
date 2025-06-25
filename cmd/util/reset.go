/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package util

import (
	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/spf13/cobra"
)

// Description
var resetSDesc = "uninstall a Kind cluster by removing all specific components on the Kind node."
var resetLDesc = resetSDesc + ` xxx.`

// root Command
var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: resetSDesc,
	Long:  resetLDesc,
	Run: func(cmd *cobra.Command, args []string) {
		logx.L.Infof("%s", resetSDesc)
	},
}
