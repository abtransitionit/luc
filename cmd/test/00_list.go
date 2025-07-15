/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package test

import (
	"github.com/spf13/cobra"
)

// Description
var testSDesc = "Test some code."
var testLDesc = testSDesc + ` xxx.`

// root Command
var TestCmd = &cobra.Command{
	Use:   "test",
	Short: testSDesc,
	Long:  testLDesc,
}

func init() {
	TestCmd.AddCommand(unitCmd)
	TestCmd.AddCommand(draftCmd)
	// TestCmd.Flags().BoolVarP(&forceFlag, "force", "f", false, "Bypass confirmation")
	// TestCmd.Flags().BoolP("list", "l", false, "List all available phases")
	// TestCmd.Flags().BoolP("runall", "r", false, "Run all phases in batch mode")
	// // Make them mutually exclusive
	// TestCmd.MarkFlagsMutuallyExclusive("list", "runall")
}
