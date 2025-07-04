/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package oservice

import (
	"fmt"

	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/abtransitionit/luc/pkg/util"
	"github.com/spf13/cobra"
)

// Description
var cFileSDesc = "create services unit files."
var cFileLDesc = cFileSDesc + `xxx`

// Definition
var cFileCmd = &cobra.Command{
	Use:   "cfile [STRING] [FILE PATH]",
	Args:  cobra.ExactArgs(2), // Requires exactly 1 arguments
	Short: cFileSDesc,
	Long:  cFileLDesc,
	// Code to play
	Run: func(cmd *cobra.Command, args []string) {

		// manage flag - foce flag is mandatory to use this command
		if !forceFlag {
			logx.L.Infof("use --force to run this command.")
			logx.L.Infof("also check --help for more details")
			return
		}

		// use caases
		if localFlag {
			helperCFileFlagLocal(args)
		}
		if remoteFlag != "" {
			helperCFileFlagRemote(args)
		}

	},
}

func init() {
	cFileCmd.Flags().BoolVar(&forceFlag, "force", false, "Force execution is mandatory")
	cFileCmd.Flags().StringVar(&remoteFlag, "remote", "", "Download remotely from a target host (e.g., o1u)")
	cFileCmd.Flags().BoolVar(&localFlag, "local", false, "Download locally")
	cFileCmd.MarkFlagsMutuallyExclusive("local", "remote")
	cFileCmd.PreRunE = func(cmd *cobra.Command, args []string) error {
		if remoteFlag == "" && !localFlag {
			return fmt.Errorf("you must specify either --remote or --local")
		}
		return nil
	}

}

func helperCFileFlagLocal(args []string) {
	content := args[0]
	filePath := args[1]
	logx.L.Debugf("saving a service unit file locally %s", filePath)
	err := util.CreateServiceFile(content, filePath)
	if err != nil {
		logx.L.Debugf("%s", err)
		return
	}

}

func helperCFileFlagRemote(args []string) {
	content := args[0]
	filePath := args[1]

	logx.L.Debugf("saving a service unit file remotely with RLUC %s", filePath)
	cli := fmt.Sprintf(`luc util oservice cfile %s %s --local --force`, content, filePath)
	_, err := util.RunCLIRemote(cli, remoteFlag)
	if err != nil {
		logx.L.Debugf("%s", err)
		return
	}
}
