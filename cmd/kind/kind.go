/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package kind

import (
	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/spf13/cobra"
)

// Description
var kindSDesc = "Deploy and manage Kind clusters."
var kindLDesc = kindSDesc + ` xxx.`

// root Command
var KindCmd = &cobra.Command{
	Use:   "kind",
	Short: kindSDesc,
	Long:  kindLDesc,
	Run: func(cmd *cobra.Command, args []string) {
		logx.L.Infof("%s", kindSDesc)
	},
}
