/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package cli

import (
	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/spf13/cobra"
)

// Description
var runcSDesc = "download tgz and install the CLI."
var runcLDesc = runcSDesc + ` xxx.`

// root Command
var runcCmd = &cobra.Command{
	Use:   "runc",
	Short: runcSDesc,
	Long:  runcLDesc,
	Run: func(cmd *cobra.Command, args []string) {
		logx.L.Debugf("%s", runcSDesc)
	},
}
