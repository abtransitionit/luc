/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package test

import (
	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/abtransitionit/luc/test"
	"github.com/spf13/cobra"
)

// Description
var draftSDesc = "Test some code."
var draftLDesc = draftSDesc + ` xxx.`

// root Command
var draftCmd = &cobra.Command{
	Use:   "draft",
	Short: draftSDesc,
	Long:  draftLDesc,
	Run: func(cmd *cobra.Command, args []string) {
		logx.L.Info(draftSDesc)

		// test.CheckCliExits("gpg")
		// // out := test.GeRemoteProperty("o1u", "osfamily")
		// out, err := test.GetPackage1("o1u", "gpg")
		// if err != nil {
		// 	logx.L.Error("%v : %s", err, out)
		// 	return
		// }
		// logx.L.Infof("yo in test : %s", out)
		// local function tested
		// test.TouchFileLocal("/tmp", "titi")
		// test.CheckFileLocalExits("/tmp/test.txt")
		// remote function tested
		// test.TouchFileOnRemote("o1u", "/tmp", "toto")
		// test.CheckFileRemoteExists("o1u", "/tmp/toto")

		// test.TestCheckCliExistsOnremote(config.KbeListNode, "gpg")
		// test.TestCheckCliExistsOnremote(config.KbeListNode, "curl")

		url := "https://pkgs.k8s.io/core:/stable:/v1.32/rpm/repodata/repomd.xml.key"
		path := "/etc/apt/sources.list.d/kbe-k8s-apt-keyring.gpg"
		vm := "o1u"
		test.TestRemoteGetGpgFromUrl(vm, url, path, true)
		// test.TestGetGpgFromUrl(url, path, true)

		// test.TestVmAreSshReachable(config.KbeListNode)
		// test.DaAddRepoLocal("kbe-k8s")

		// createFileLocal()
		// touchFileRemote("o1u")
		// MoveFileLocal()
		// ListOvhVm()
		// ListMapKey()
		// installGoCli()
		// getPath()
		// fmt.Println(configi.KbeGoCliConfigMap)
		// addLineToFileRemote()

	},
}

func init() {
	draftCmd.Flags().BoolVarP(&forceFlag, "force", "f", false, "Bypass confirmation")
	draftCmd.Flags().BoolP("list", "l", false, "List all available phases")
	draftCmd.Flags().BoolP("runall", "r", false, "Run all phases in batch mode")
	// Make them mutually exclusive
	draftCmd.MarkFlagsMutuallyExclusive("list", "runall")
}
