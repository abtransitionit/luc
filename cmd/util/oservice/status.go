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
var statusSDesc = "status an OS service."
var statusLDesc = statusSDesc + `xxx`

// Definition
var statusCmd = &cobra.Command{
	Use:   "status [SERVICE NAME]",
	Args:  cobra.ExactArgs(1), // Requires exactly 1 arguments
	Short: statusSDesc,
	Long:  statusLDesc,
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
			helperStatusFlagLocal(args)
			return
		}
		if remoteFlag != "" {
			helperStatusFlagRemote(args)
			return
		}

	},
}

func init() {
	statusCmd.Flags().BoolVar(&forceFlag, "force", false, "Force execution is mandatory")
	statusCmd.Flags().StringVar(&remoteFlag, "remote", "", "Status remotelyonto a target host (e.g., o1u)")
	statusCmd.Flags().BoolVar(&localFlag, "local", false, "Status locally")
	statusCmd.MarkFlagsMutuallyExclusive("local", "remote")
	statusCmd.PreRunE = func(cmd *cobra.Command, args []string) error {
		if remoteFlag == "" && !localFlag {
			return fmt.Errorf("you must specify either --remote or --local")
		}
		return nil
	}

}

func helperStatusFlagLocal(args []string) {
	serviceName := args[0]

	logx.L.Infof("start an os service locally")
	result, err := util.StatusService(serviceName)
	if err != nil {
		logx.L.Debugf("%s", err)
		return
	}
	fmt.Println("⚠️ ⚠️")
	fmt.Println(result)
}

func helperStatusFlagRemote(args []string) {
	serviceName := args[0]
	logx.L.Infof("start an os service rmotely")
	cli := fmt.Sprintf(`luc util oservice start %s --local --force`, serviceName)
	_, err := util.RunCLIRemote(cli, remoteFlag)
	if err != nil {
		logx.L.Debugf("%s", err)
		return
	}
}
