/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package dnfapt

import (
	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/spf13/cobra"
)

// Description
var dnfaptSDesc = "Manage packages and repositories (both Rhel and Debian based)."
var dnfaptLDesc = dnfaptSDesc + ` xxx.`

// root Command
var DnfaptCmd = &cobra.Command{
	Use:   "dnfapt",
	Short: dnfaptSDesc,
	Long:  dnfaptLDesc,
	Run: func(cmd *cobra.Command, args []string) {
		logx.L.Infof("%s", dnfaptSDesc)
	},
}

func init() {
	DnfaptCmd.AddCommand(repoCmd)
	DnfaptCmd.AddCommand(packageCmd)
}
