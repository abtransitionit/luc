/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package dnfapt

import (
	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/spf13/cobra"
)

// Description
var repoSDesc = "Uninstall a KBE cluster by removing all specific components on all KBE nodes."
var repoLDesc = repoSDesc + ` xxx.`

// root Command
var repoCmd = &cobra.Command{
	Use:   "repo",
	Short: repoSDesc,
	Long:  repoLDesc,
	Run: func(cmd *cobra.Command, args []string) {
		logx.L.Infof("%s", repoSDesc)
	},
}
