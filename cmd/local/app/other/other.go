/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package other

import (
	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/spf13/cobra"
)

// Description
var otherSDesc = "Manage other applications."
var otherLDesc = otherSDesc + ` xxx.`

// root Command
var OtherCmd = &cobra.Command{
	Use:   "other",
	Short: otherSDesc,
	Long:  otherLDesc,
	Run: func(cmd *cobra.Command, args []string) {
		logx.L.Debugf(otherSDesc)
	},
}
