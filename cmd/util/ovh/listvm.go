/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package ovh

import (
	"fmt"

	"github.com/abtransitionit/luc/internal/util"
	"github.com/abtransitionit/luc/pkg/logx"
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
		logx.L.Infof("wait a minute: time to test all possible OVH VMs, from your declared OVH configuration.")
		fmt.Println(util.ListOvhVm())
	},
}
