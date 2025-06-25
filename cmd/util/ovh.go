/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package util

import (
	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/abtransitionit/luc/pkg/util"
	"github.com/spf13/cobra"
)

// Description
var ovhSDesc = "manage OVH objects."
var ovhLDesc = ovhSDesc + ` xxx.`

// delete Command
var ovhCmd = &cobra.Command{
	Use:   "ovh",
	Short: ovhSDesc,
	Long:  ovhLDesc,
	// define the set of phases for this cmd
	Run: func(cmd *cobra.Command, args []string) {
		logx.L.Infof("%s", ovhSDesc)
		cli := "kind delete cluster"
		_, err := util.RunCLILocal(cli)
		if err != nil {
			logx.L.Debugf("❌ Error detected")
		}
		logx.L.Debugf("✅ deleted cluster")
	},
}
