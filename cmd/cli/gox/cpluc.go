/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package gox

import (
	"fmt"

	"github.com/abtransitionit/luc/internal/config"
	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/abtransitionit/luc/pkg/util"
	"github.com/spf13/cobra"
)

var LucGitProjectFolder = config.LucGitProjectFolder
var LucBinaryPath = config.LucBinaryPath
var LucBinaryTmpPtfPath = config.LucBinaryTmpPtfPath

// Description
var cplucSDesc = "building and deploying luc"
var cplucLDesc = cplucSDesc + fmt.Sprintf(`:
- build luc locally from GIT folder %s for current platform
- deploy it locally to %s

Example usage:

luc cli go cpluc --force
`, LucGitProjectFolder, LucBinaryPath)

// root Command
var cplucCmd = &cobra.Command{
	Use:   "cpluc",
	Short: cplucSDesc,
	Long:  cplucLDesc,
	Run: func(cmd *cobra.Command, args []string) {
		logx.L.Debugf(cplucSDesc)

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

		// build
		if _, err := util.GoBuild(LucGitProjectFolder, LucBinaryTmpPtfPath); err != nil {
			logx.L.Debugf("%s", err)
			return
		}

		// deploy
		if _, err := util.MvFile(LucBinaryTmpPtfPath, LucBinaryPath, 0755, true); err != nil {
			logx.L.Debugf("%s", err)
			return
		}
		logx.L.Debugf("builded and deployed LUC to %s", LucBinaryPath)
	},
}

func init() {
	cplucCmd.Flags().BoolVarP(&forceFlag, "force", "f", false, "Bypass confirmation and force execution")
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
