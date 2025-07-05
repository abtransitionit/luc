/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package gox

import (
	"github.com/spf13/cobra"
)

// Description
var goSDesc = "manage GO CLI/App(s)."
var goLDesc = goSDesc + ` xxx.`

// root Command
var GoCmd = &cobra.Command{
	Use:   "go",
	Short: goSDesc,
	Long:  goLDesc,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var forceFlag bool
var archFlag string
var osFlag string

//	func init() {
//		goCmd.Flags().BoolP("show", "s", false, "show CLI config map")
//	}
func init() {
	GoCmd.AddCommand(buildCmd)
	GoCmd.AddCommand(runCmd)
	GoCmd.AddCommand(isdkCmd)
	GoCmd.AddCommand(bdLucCmd)
}
