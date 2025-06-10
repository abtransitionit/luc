/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package cli

import (
	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/spf13/cobra"
)

// Description
var sonobuoySDesc = "download source code, then build and install the CLI."
var sonobuoyLDesc = sonobuoySDesc + ` xxx.`

// root Command
var sonobuoyCmd = &cobra.Command{
	Use:   "sonobuoy",
	Short: sonobuoySDesc,
	Long:  sonobuoyLDesc,
	Run: func(cmd *cobra.Command, args []string) {
		logx.L.Debugf("%s", sonobuoySDesc)
	},
}
