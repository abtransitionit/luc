/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package service

import (
	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/spf13/cobra"
)

// Description
var serviceSDesc = "manage Linux OS services."
var serviceLDesc = serviceSDesc + ` xxx.`

// root Command
var ServiceCmd = &cobra.Command{
	Use:   "service",
	Short: serviceSDesc,
	Long:  serviceLDesc,
	Run: func(cmd *cobra.Command, args []string) {
		logx.L.Debugf("%s", serviceSDesc)
	},
}

func init() {
	ServiceCmd.AddCommand(addCmd)
	ServiceCmd.AddCommand(enableCmd)
}
