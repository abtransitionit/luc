/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package service

import (
	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/spf13/cobra"
)

// Description
var addSDesc = "Add a custom unit file for a service."
var addLDesc = addSDesc + ` xxx.`

// root Command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: addSDesc,
	Long:  addLDesc,
	Run: func(cmd *cobra.Command, args []string) {
		logx.L.Debugf("%s", addSDesc)
	},
}
