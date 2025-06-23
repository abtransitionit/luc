/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package app

import (
	"github.com/abtransitionit/luc/cmd/local/app/dnfapt"
	"github.com/abtransitionit/luc/cmd/local/app/gox"
	"github.com/abtransitionit/luc/cmd/local/app/other"
	"github.com/abtransitionit/luc/cmd/local/app/python"
	"github.com/spf13/cobra"
)

// Description
var appSDesc = "manage linux open source apps."
var appLDesc = appSDesc + ` xxx.`

// root Command
var AppCmd = &cobra.Command{
	Use:   "app",
	Short: appSDesc,
	Long:  appLDesc,
}

func init() {
	AppCmd.AddCommand(dnfapt.DnfaptCmd)
	AppCmd.AddCommand(gox.GoCmd)
	AppCmd.AddCommand(python.PythonCmd)
	AppCmd.AddCommand(other.OtherCmd)
}
