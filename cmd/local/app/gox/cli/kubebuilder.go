/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package cli

import (
	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/spf13/cobra"
)

// Description
var kubebuilderSDesc = "download source code, then build and install the CLI."
var kubebuilderLDesc = kubebuilderSDesc + ` xxx.`

// root Command
var kubebuilderCmd = &cobra.Command{
	Use:   "kubebuilder",
	Short: kubebuilderSDesc,
	Long:  kubebuilderLDesc,
	Run: func(cmd *cobra.Command, args []string) {
		logx.L.Debugf("%s", kubebuilderSDesc)
	},
}
