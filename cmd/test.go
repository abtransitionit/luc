/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package cmd

import (
	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/abtransitionit/luc/pkg/netx"
	"github.com/spf13/cobra"
)

// Description
var testSDesc = "Test some code."
var testLDesc = testSDesc + ` xxx.`

// root Command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: testSDesc,
	Long:  testLDesc,
	Run: func(cmd *cobra.Command, args []string) {
		logx.L.Info(testSDesc)

		// get input
		param := args[0]

		// threat input
		// netx.IsVmSshConfigured(param)
		netx.IsSshConfiguredVmSshReachable(param)
		// if lucnet.IsSshReachable("o1u") {
		// 	logx.L.Info("SSH reachable")
		// }

	},
}
