/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/abtransitionit/luc/pkg/config"
	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/abtransitionit/luc/pkg/ui"
	"github.com/abtransitionit/luc/pkg/util"
	"github.com/abtransitionit/luc/pkg/util/dnfapt"
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
		value, err := dnfapt.UpdateOs()
		if err != nil {
			logx.L.Debugf("%s", err)
		}
		logx.L.Infof(" Status %v", value)
		return
		props := []string{
			"cpu", "ram", "osarch", "ostype", "osfamily", "osdistro", "osversion", "oskversion",
			"uuid", "cgroup", "osuser",
			"netgateway", "netip", "init",
			"selinfos", "osinfos"}
		var listPropertOut = []string{}
		for _, prop := range props {
			value, err := util.OsPropertyGet(prop)
			if err != nil {
				// logx.L.Debugf("%s", err)
				listPropertOut = append(listPropertOut, prop)
				continue
			}
			fmt.Printf("%-10s: %s\n", prop, value)
		}
		if len(listPropertOut) > 0 {
			logx.L.Debugf("❌ List property OUT: %s", listPropertOut)
		}
		return
		// config.DisplayCliCondfigInfo()
		cliUrl, err := config.GetCliProperty(logx.L, "nerdctl", "url")
		if err != nil {
			// logx.L.Infof("cliName: %s cliUrl: %s ", cliUrl)
			return
		}
		cliUrl, err = config.GetCliSpecificUrl(logx.L, "nerdctl", "linux", "amd64")
		if err == nil {
			logx.L.Infof("cliUrl: %s", cliUrl)
		}
		// download the file as memory data
		fileInMemory, err := util.GetPublicFile(cliUrl)
		if err != nil {
			return
		}
		// save memory data as file
		_, err = util.SaveToFile("/tmp/toto", fileInMemory)
		if err != nil {
			return
		}
		// logx.L.Infof("info: %s", info)

		// get the cli url
		// url, _ := cli.ResolveCliURL("nerdctl") // OS and Arch auto-detected
		// logx.L.Infof("url: %s", url)
		// // download the file
		// cli.DownloadCli(url, "nerdctl")

		// ui.AskForConfirmation(logx.L, testSDesc)
		// return
		// question := "yo vas ya ?"
		// input, err := ui.ReadUserInput(fmt.Sprintf("%s [y/N]: ", question))
		confirmed := ui.ConfirmAction(logx.L, "do action")
		if !confirmed {
			logx.L.Infof("action canceled or system failure")
			return
		}
		logx.L.Infof("action confirmed")

		// manage argument
		if len(args) == 0 {
			logx.L.Debugf("❌ No argument provided:exiting")
			os.Exit(1)
		} else if len(args) > 1 {
			logx.L.Debugf("❌ several argument provided, 1 required:exiting")
			os.Exit(2)
		}

		// get input
		param := args[0]
		logx.L.Debugf("provided argument : %s", param)

		// // _, err := netx.IsSshConfiguredVmSshReachable(param)
		// configured, err := netx.IsVmSshConfigured(param)
		// if err != nil {
		// 	logx.L.Errorf("❌ system failure : %w", err)
		// 	os.Exit(3)
		// } else if !configured {
		// 	logx.L.Debugf("❌ exiting: vm %s is not configured in ssh", param)
		// 	os.Exit(4)
		// } else {
		// 	logx.L.Infof("✅ vm %s is potentially configured", param)
		// }

		_, err = util.IsSshConfiguredVmSshReachable(param)
		if err != nil {
			os.Exit(5)
		}

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
