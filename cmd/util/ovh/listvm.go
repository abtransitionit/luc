/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package ovh

import (
	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/abtransitionit/luc/pkg/phase"
	"github.com/spf13/cobra"
)

// Description
var listvmSDesc = "List configured and reachable OVH VMs."
var listvmLDesc = listvmSDesc + ` xxx.`

// init Command
var ListvmCmd = &cobra.Command{
	Use:   "listvm",
	Short: listvmSDesc,
	Long:  listvmLDesc,
	Run: func(cmd *cobra.Command, args []string) {
		logx.L.Infof("%s", listvmSDesc)
		cmd.Help()
	},
}

func init() {
	phase.CmdInit(ListvmCmd)
}
