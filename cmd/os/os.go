/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package os

import (
	"github.com/abtransitionit/luc/cmd/os/dnfapt"
	"github.com/abtransitionit/luc/cmd/os/gox"
	"github.com/abtransitionit/luc/cmd/os/kernelx"
	"github.com/abtransitionit/luc/cmd/os/service"
	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/spf13/cobra"
)

// Description
var osSDesc = "Manage linux OS objects."
var osLDesc = osSDesc + ` xxx.`

// root Command
var OsCmd = &cobra.Command{
	Use:   "os",
	Short: osSDesc,
	Long:  osLDesc,
	Run: func(cmd *cobra.Command, args []string) {
		logx.L.Infof("%s", osSDesc)
	},
}

func init() {
	OsCmd.AddCommand(dnfapt.DnfaptCmd)
	OsCmd.AddCommand(gox.GoCmd)
	OsCmd.AddCommand(kernelx.KernelCmd)
	OsCmd.AddCommand(service.ServiceCmd)
}
