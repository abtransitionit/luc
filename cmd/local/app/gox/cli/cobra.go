/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package cli

import (
	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/spf13/cobra"
)

// Description
var cobraSDesc = "Install the cli cobra-cli."
var cobraLDesc = cobraSDesc + ` xxx.`

// root Command
var cobraCmd = &cobra.Command{
	Use:   "cobra",
	Short: cobraSDesc,
	Long:  cobraLDesc,
	Run: func(cmd *cobra.Command, args []string) {
		logx.L.Debugf("%s", cobraSDesc)
	},
}
