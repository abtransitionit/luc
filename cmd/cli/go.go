/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package cli

import (
	"github.com/spf13/cobra"
)

// Description
var goSDesc = "manage GO CLI/App(s)."
var goLDesc = goSDesc + ` xxx.`

// root Command
var goCmd = &cobra.Command{
	Use:   "go",
	Short: goSDesc,
	Long:  goLDesc,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	goCmd.Flags().BoolP("show", "s", false, "show CLI config map")
}
