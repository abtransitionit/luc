/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package cli

import (
	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/spf13/cobra"
)

// Description
var kindSDesc = "download binary and install it."
var kindLDesc = kindSDesc + ` xxx.`

// root Command
var kindCmd = &cobra.Command{
	Use:   "kind",
	Short: kindSDesc,
	Long:  kindLDesc,
	Run: func(cmd *cobra.Command, args []string) {
		logx.L.Debugf("%s", kindSDesc)
	},
}
