/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package kbe

import (
	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/spf13/cobra"
)

// Description
var initSDesc = "Deploy a kubernetes cluster. It becomes a KBE (Kubernetes Easy) cluster."
var initLDesc = initSDesc + ` xxx.`

// root Command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: initSDesc,
	Long:  initLDesc,
	Run: func(cmd *cobra.Command, args []string) {
		logx.L.Infof("%s", initSDesc)
	},
}
