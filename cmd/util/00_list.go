/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package util

import (
	"github.com/abtransitionit/luc/cmd/util/oservice"
	"github.com/abtransitionit/luc/cmd/util/ovh"
	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/spf13/cobra"
)

// Description
var utilSDesc = "manage Kind clusters."
var utilLDesc = utilSDesc + ` xxx.`

// root Command
var UtilCmd = &cobra.Command{
	Hidden: true, // available but not visible
	Use:    "util",
	Short:  utilSDesc,
	Long:   utilLDesc,
	Run: func(cmd *cobra.Command, args []string) {
		logx.L.Infof("%s", utilSDesc)
		// manage argument
		if len(args) == 0 {
			cmd.Help()
			return
		}

	},
}

var forceFlag bool
var remoteFlag string
var localFlag bool

func init() {
	UtilCmd.AddCommand(lineFileCmd)
	UtilCmd.AddCommand(strFileCmd)
	// UtilCmd.AddCommand(mvfileCmd)
	UtilCmd.AddCommand(mvdirCmd)
	UtilCmd.AddCommand(ovh.OvhCmd)
	UtilCmd.AddCommand(urlCmd)
	UtilCmd.AddCommand(utgzCmd)
	UtilCmd.AddCommand(oservice.OsServiceCmd)
}
