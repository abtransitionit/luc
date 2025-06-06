/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package cli

import (
	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/spf13/cobra"
)

// Description
var kubectlSDesc = "download binary and install it."
var kubectlLDesc = kubectlSDesc + ` xxx.`

// root Command
var kubectlCmd = &cobra.Command{
	Use:   "kubectl",
	Short: kubectlSDesc,
	Long:  kubectlLDesc,
	Run: func(cmd *cobra.Command, args []string) {
		logx.L.Debugf("%s", kubectlSDesc)
	},
}
