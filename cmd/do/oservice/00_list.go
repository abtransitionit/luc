/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package oservice

import (
	"github.com/spf13/cobra"
)

// Description
var osServiceSDesc = "manage OS services."
var osServiceLDesc = osServiceSDesc + ` xxx.`

// delete Command
var OsServiceCmd = &cobra.Command{
	Use:   "oservice [SRC FILE PATH] [DST FOLDER PATH]",
	Args:  cobra.ExactArgs(2), // Requires exactly one argument: the URL
	Short: osServiceSDesc,
	Long:  osServiceLDesc,
}

// var for child
var forceFlag bool
var remoteFlag string
var localFlag bool

func init() {
	// OsServiceCmd.AddCommand(lingerCmd)
	// OsServiceCmd.AddCommand(cFileCmd)
	OsServiceCmd.AddCommand(startCmd)
	OsServiceCmd.AddCommand(statusCmd)
	OsServiceCmd.AddCommand(stopCmd)
}
