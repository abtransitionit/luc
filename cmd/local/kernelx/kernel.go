/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package kernelx

import (
	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/spf13/cobra"
)

// Description
var kernelSDesc = "manage linux OS kernel parameters and modules."
var kernelLDesc = kernelSDesc + ` xxx.`

// root Command
var KernelCmd = &cobra.Command{
	Use:   "kernel",
	Short: kernelSDesc,
	Long:  kernelLDesc,
	Run: func(cmd *cobra.Command, args []string) {
		logx.L.Debugf("%s", kernelSDesc)
	},
}
