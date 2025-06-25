/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package cli

import (
	"github.com/spf13/cobra"
)

// Description
var ovhSDesc = "install a GO CLI."
var ovhLDesc = ovhSDesc + ` xxx.`

// root Command
var ovhCmd = &cobra.Command{
	Use:   "ovh",
	Short: ovhSDesc,
	Long:  ovhLDesc,
	Run: func(cmd *cobra.Command, args []string) {

	},
}
