/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package cli

import (
	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/spf13/cobra"
)

// Description
var helmSDesc = "download source code, then build and install the CLI."
var helmLDesc = helmSDesc + ` xxx.`

// root Command
var helmCmd = &cobra.Command{
	Use:   "helm",
	Short: helmSDesc,
	Long:  helmLDesc,
	Run: func(cmd *cobra.Command, args []string) {
		logx.L.Debugf("%s", helmSDesc)
	},
}
