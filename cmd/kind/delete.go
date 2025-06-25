/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package kind

import (
	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/abtransitionit/luc/pkg/util"
	"github.com/spf13/cobra"
)

// Description
var deleteSDesc = "delete default Kind cluster on a VM."
var deleteLDesc = deleteSDesc + ` xxx.`

// delete Command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: deleteSDesc,
	Long:  deleteLDesc,
	// define the set of phases for this cmd
	Run: func(cmd *cobra.Command, args []string) {
		logx.L.Infof("%s", deleteSDesc)

		// are we on a linux VM
		cli := "luc util getprop ostype"
		ostype, err := util.RunCLILocal(cli)
		if err != nil {
			logx.L.Debugf("%s", err)
		}
		if ostype != "linux" {
			logx.L.Debugf("❌ not on a linux VM")
			return
		}
		logx.L.Debugf("✅ on a linux VM. will delete cluster")

		// cli := "kind delete cluster"
		// _, err := util.RunCLILocal(cli)
		// if err != nil {
		// 	logx.L.Debugf("❌ Error detected")
		// }
		// logx.L.Debugf("✅ deleted cluster")
	},
}
