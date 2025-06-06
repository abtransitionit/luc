/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package service

import (
	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/spf13/cobra"
)

// Description
var enableSDesc = "Enable a service to start after a reboot."
var enableLDesc = enableSDesc + ` xxx.`

// root Command
var enableCmd = &cobra.Command{
	Use:   "enable",
	Short: enableSDesc,
	Long:  enableLDesc,
	Run: func(cmd *cobra.Command, args []string) {
		logx.L.Debugf("%s", enableSDesc)
	},
}
