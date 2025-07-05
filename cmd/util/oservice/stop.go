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
var stopSDesc = "stop an OS service."
var stopLDesc = stopSDesc + `xxx`

// Definition
var stopCmd = &cobra.Command{
	Use:   "stop [SERVICE NAME]",
	Args:  cobra.ExactArgs(1), // Requires exactly 1 arguments
	Short: stopSDesc,
	Long:  stopLDesc,
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
			helperStopFlagLocal(args)
			return
		}
		if remoteFlag != "" {
			helperStopFlagRemote(args)
			return
		}

	},
}

func init() {
	stopCmd.Flags().BoolVar(&forceFlag, "force", false, "Force execution is mandatory")
	stopCmd.Flags().StringVar(&remoteFlag, "remote", "", "Stop remotelyonto a target host (e.g., o1u)")
	stopCmd.Flags().BoolVar(&localFlag, "local", false, "Stop locally")
	stopCmd.MarkFlagsMutuallyExclusive("local", "remote")
	stopCmd.PreRunE = func(cmd *cobra.Command, args []string) error {
		if remoteFlag == "" && !localFlag {
			return fmt.Errorf("you must specify either --remote or --local")
		}
		return nil
	}

}

func helperStopFlagLocal(args []string) {
	serviceName := args[0]

	logx.L.Infof("stop an os service locally")
	err := util.StopService(serviceName)
	if err != nil {
		logx.L.Debugf("%s", err)
		return
	}

}

func helperStopFlagRemote(args []string) {
	serviceName := args[0]
	logx.L.Infof("stop an os service rmotely")
	cli := fmt.Sprintf(`luc util oservice stop %s --local --force`, serviceName)
	_, err := util.RunCLIRemote(cli, remoteFlag)
	if err != nil {
		logx.L.Debugf("%s", err)
		return
	}
}
