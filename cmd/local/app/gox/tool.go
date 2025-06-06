/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package gox

import (
	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/spf13/cobra"
)

// Description
var toolSDesc = "install the GO chaintool (`go` binary and  libraries to start coding in go)."
var toolLDesc = toolSDesc + ` xxx.`

// root Command
var toolCmd = &cobra.Command{
	Use:   "tool",
	Short: toolSDesc,
	Long:  toolLDesc,
	Run: func(cmd *cobra.Command, args []string) {
		logx.L.Debugf(toolSDesc)
	},
}
