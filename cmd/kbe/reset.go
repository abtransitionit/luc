/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package kbe

import (
	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/spf13/cobra"
)

// Description
var resetSDesc = "uninstall a KBE cluster by removing all specific components on all KBE nodes."
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
