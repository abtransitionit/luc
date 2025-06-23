/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package local

import (
	"github.com/abtransitionit/luc/cmd/local/app"
	"github.com/abtransitionit/luc/cmd/local/kernelx"
	"github.com/abtransitionit/luc/cmd/local/service"
	"github.com/spf13/cobra"
)

// Description
var localSDesc = "Manage local OS objects."
var localLDesc = localSDesc + ` xxx.`

// root Command
var LocalCmd = &cobra.Command{
	Use:   "local",
	Short: localSDesc,
	Long:  localLDesc,
}

func init() {
	LocalCmd.AddCommand(app.AppCmd)
	LocalCmd.AddCommand(kernelx.KernelCmd)
	LocalCmd.AddCommand(service.ServiceCmd)
}
