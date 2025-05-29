/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package kbe

import (
	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/spf13/cobra"
)

// Description
var kbeSDesc = "Deploy and manage Kubernetes clusters (KBE = Kubernetes Easy)."
var kbeLDesc = kbeSDesc + ` xxx.`

// root Command
var KbeCmd = &cobra.Command{
	Use:   "kbe",
	Short: kbeSDesc,
	Long:  kbeLDesc,
	Run: func(cmd *cobra.Command, args []string) {
		logx.L.Infof("%s", kbeSDesc)
	},
}

func init() {
	KbeCmd.AddCommand(initCmd)
	KbeCmd.AddCommand(resetCmd)
}
