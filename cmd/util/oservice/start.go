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
var startSDesc = "start an OS service."
var startLDesc = startSDesc + `xxx`

// Definition
var startCmd = &cobra.Command{
	Use:   "start [SERVICE NAME]",
	Args:  cobra.ExactArgs(1), // Requires exactly 1 arguments
	Short: startSDesc,
	Long:  startLDesc,
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
			helperStartFlagLocal(args)
			return
		}
		if remoteFlag != "" {
			helperStartFlagRemote(args)
			return
		}

	},
}

func init() {
	startCmd.Flags().BoolVar(&forceFlag, "force", false, "Force execution is mandatory")
	startCmd.Flags().StringVar(&remoteFlag, "remote", "", "Start remotelyonto a target host (e.g., o1u)")
	startCmd.Flags().BoolVar(&localFlag, "local", false, "start locally")
	startCmd.MarkFlagsMutuallyExclusive("local", "remote")
	startCmd.PreRunE = func(cmd *cobra.Command, args []string) error {
		if remoteFlag == "" && !localFlag {
			return fmt.Errorf("you must specify either --remote or --local")
		}
		return nil
	}

}

func helperStartFlagLocal(args []string) {
	serviceName := args[0]

	logx.L.Infof("start an os service locally")
	err := util.StartService(serviceName)
	if err != nil {
		logx.L.Debugf("%s", err)
		return
	}

}

func helperStartFlagRemote(args []string) {
	serviceName := args[0]
	logx.L.Infof("start an os service rmotely")
	cli := fmt.Sprintf(`luc util oservice start %s --local --force`, serviceName)
	_, err := util.RunCLIRemote(remoteFlag, cli)
	if err != nil {
		logx.L.Debugf("%s", err)
		return
	}
}
