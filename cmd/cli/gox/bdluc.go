/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package gox

import (
	"github.com/abtransitionit/luc/internal/config"
	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/abtransitionit/luc/pkg/util"
	"github.com/abtransitionit/luc/pkg/util/gocli"
	"github.com/spf13/cobra"
)

// Description
var bdLucSDesc = "building and deploying luc"
var bdLucLDesc = bdLucSDesc + `:
- build luc locally from GIT folder for current platform
- deploy it locally to final path

Example usage:

luc cli go bdLuc --force
`

// root Command
var bdLucCmd = &cobra.Command{
	Use:   "bdluc",
	Short: bdLucSDesc,
	Long:  bdLucLDesc,
	Run: func(cmd *cobra.Command, args []string) {
		logx.L.Debugf(bdLucSDesc)

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
		LucGitProjectFolder := config.LucGitProjectFolder
		LucBinaryPath := config.LucBinaryPath
		LucBinaryTmpPtfPath := config.LucBinaryTmpPtfPath

		// build
		if _, err := gocli.GoBuild(LucGitProjectFolder, LucBinaryTmpPtfPath); err != nil {
			logx.L.Debugf("%s", err)
			return
		}

		// deploy
		if _, err := util.MvFile2(LucBinaryTmpPtfPath, LucBinaryPath, 0755, true); err != nil {
			logx.L.Debugf("%s", err)
			return
		}
		logx.L.Debugf("builded and deployed LUC to %s", LucBinaryPath)
	},
}

func init() {
	bdLucCmd.Flags().BoolVarP(&forceFlag, "force", "f", false, "Mandatory to use this command")
	bdLucCmd.Flags().StringVar(&osTypeFlag, "os", "", "Target OS for cross-compilation (e.g., linux, windows, darwin)")
	bdLucCmd.Flags().StringVar(&osArchFlag, "arch", "", "Target architecture for cross-compilation (e.g., amd64, arm64)")

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
