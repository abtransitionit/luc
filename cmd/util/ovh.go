/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package util

import (
	"github.com/abtransitionit/luc/cmd/util/ovh"
	"github.com/abtransitionit/luc/pkg/logx"
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
		cmd.Help()
	},
}

func init() {
	ovhCmd.AddCommand(ovh.ListvmCmd)
	ovhCmd.AddCommand(ovh.CplucCmd)
}

// list ovh vm
// cp luc to ovh vm provided as arg
