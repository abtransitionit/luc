/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package python

import (
	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/spf13/cobra"
)

// Description
var pythonSDesc = "Manage python applications."
var pythonLDesc = pythonSDesc + ` xxx.`

// root Command
var PythonCmd = &cobra.Command{
	Use:   "python",
	Short: pythonSDesc,
	Long:  pythonLDesc,
	Run: func(cmd *cobra.Command, args []string) {
		logx.L.Debugf(pythonSDesc)
	},
}
