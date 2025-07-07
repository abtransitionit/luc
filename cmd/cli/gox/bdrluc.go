/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package gox

import (
	"fmt"

	"github.com/abtransitionit/luc/internal/config"
	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/abtransitionit/luc/pkg/util/gocli"
	"github.com/spf13/cobra"
)

// Description
var bdrLucSDesc = "building Luc locally and deploying it remotley "
var bdrLucLDesc = bdrLucSDesc + `:
- build luc locally from local GIT folder for linux/amd64 platform
- deploy it to linux/amd64 VM(s)

Example usage:

luc cli go bdrLuc --force --remote o1u o2a
`

// root Command
var bdrLucCmd = &cobra.Command{
	Use:   "bdrluc",
	Short: bdrLucSDesc,
	Long:  bdrLucLDesc,
	Run: func(cmd *cobra.Command, args []string) {
		logx.L.Debugf(bdrLucSDesc)

		// If no flags and no args => show help and return
		if !cmd.Flags().Changed("force") && len(args) == 0 {
			cmd.Help()
			return
		}

		// force flag is mandatory to use this command
		if !forceFlag {
			logx.L.Infof("use --force to run this command.")
			logx.L.Infof("also check --help for more details")
			return
		}

		// define var
		osType := "linux"
		osArch := "amd64"
		LucGitProjectFolder := config.LucGitProjectFolder
		// LucBinaryPath := config.LucBinaryPath
		LucBinaryXptTmpfPath := fmt.Sprintf("/tmp/luc-%s-%s", osType, osArch)

		// build
		if _, err := gocli.GoBuildXPtf(LucGitProjectFolder, LucBinaryXptTmpfPath, osType, osArch); err != nil {
			logx.L.Debugf("%s", err)
			return
		}

		// // deploy
		// if _, err := util.MvFile(LucBinaryTmpPtfPath, LucBinaryPath, 0755, true); err != nil {
		// 	logx.L.Debugf("%s", err)
		// 	return
		// }
		logx.L.Debugf("builded and deployed LUC to %s", LucBinaryXptTmpfPath)
	},
}

func init() {
	bdrLucCmd.Flags().BoolVarP(&forceFlag, "force", "f", false, "Mandatory to use this command")
}

// // launch this pipeline
// _, err := cpluc.RunPipeline(listVm)
// if err != nil {
// 	logx.L.Debugf("%s", err)
// 	return
// }

// 		// get the list of VMs from arg
// listVm := ""
// for _, vm := range args {
// 	listVm += " " + vm
// }
