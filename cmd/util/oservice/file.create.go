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
	Use:   "cfile [FILE PATH]",
	Args:  cobra.ExactArgs(1), // Requires exactly 1 arguments
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
		// get input
		filePath := args[0]

		// use caases
		if localFlag {
			helperFlagLocal(filePath)
		}
		if remoteFlag != "" {
			helperFlagRemote(filePath)
		}

	},
}

var forceFlag bool
var remoteFlag string
var localFlag bool

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

func helperFlagLocal(filePath string) {
	logx.L.Info("saving a service unit file locally %s", filePath)
}

func helperFlagRemote(filePath string) {
	logx.L.Info("saving a service unit file remotely with RLUC %s", filePath)
	cli := fmt.Sprintf(`luc util oservice cfile %s --local --force`, remoteFlag)
	_, err := util.RunCLIRemote2(cli, remoteFlag)
	if err != nil {
		logx.L.Error(err)
	}
}
