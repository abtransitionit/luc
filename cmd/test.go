/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package cmd

import (
	configi "github.com/abtransitionit/luc/internal/config"
	utili "github.com/abtransitionit/luc/internal/util"
	configp "github.com/abtransitionit/luc/pkg/config"
	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/abtransitionit/luc/pkg/pipeline/gocli"
	"github.com/abtransitionit/luc/pkg/util"
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

		// ListOvhVm()
		ListMapKey()

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

func installGoCli() {
	// define test var
	vm := "o1u"
	// provision a cli
	configMap := configp.CustomCLIConfigMap{
		"runc": {
			Name:      "runc",
			Version:   "1.3.0",
			DstFolder: "/usr/local/bin",
		},
	}
	gocli.RunPipeline(vm, configMap)
}

func ListOvhVm() {
	logx.L.Info("List OVH Vm")
	logx.L.Info("🔹 List : %s", utili.ListOvhVm())
}

func ListMapKey() {
	logx.L.Info("List map:keys")
	// list := []string{"a", "b", "c"}
	listKey := util.GetMapKeys(configi.KindGoCliConfigMap)
	logx.L.Infof("🔹 as slice:      %s", listKey)
	logx.L.Infof("🔹 as StringList: %s", util.GetStringfromSliceWithSpace(listKey))
}
