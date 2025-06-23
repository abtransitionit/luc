/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package cmd

import (
	"github.com/abtransitionit/luc/internal/config"
	"github.com/abtransitionit/luc/internal/util"
	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/abtransitionit/luc/pkg/pipeline/gocli"
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

		ListOvhVm()
		// provisionOneCli()

	},
}

var forceFlag bool

// SetupCommonFlags configures flags that are shared across commands
func init() {
	testCmd.Flags().BoolVarP(&forceFlag, "force", "f", false, "Bypass confirmation")
	testCmd.Flags().BoolP("list", "l", false, "List all available phases")
	testCmd.Flags().BoolP("runall", "r", false, "Run all phases in batch mode")
	// Make them mutually exclusive
	testCmd.MarkFlagsMutuallyExclusive("list", "runall")
}

func provisionOneCli() {
	// provision a cli
	configMap := map[string]config.CustomCLIConfig{
		"runc": {
			Name:      "runc",
			Version:   "1.3.0",
			DstFolder: "/usr/local/bin",
		},
	}
	gocli.RunPipeline(configMap)
}

func ListOvhVm() {
	logx.L.Infof("🔹 List OVH Vm: %s", util.ListOvhVm())
}
