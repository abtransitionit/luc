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
var lingerSDesc = "manage os non-root user service enabling after log out."
var lingerLDesc = lingerSDesc + `xxx`

// Definition
var lingerCmd = &cobra.Command{
	Use:   "linger",
	Short: lingerSDesc,
	Long:  lingerLDesc,
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
			helperlingerFlagLocal()
		}
		if remoteFlag != "" {
			helperlingerFlagRemote()
		}

	},
}

func init() {
	lingerCmd.Flags().BoolVar(&forceFlag, "force", false, "Force execution is mandatory")
	lingerCmd.Flags().StringVar(&remoteFlag, "remote", "", "Download remotely from a target host (e.g., o1u)")
	lingerCmd.Flags().BoolVar(&localFlag, "local", false, "Download locally")
	lingerCmd.MarkFlagsMutuallyExclusive("local", "remote")
	lingerCmd.PreRunE = func(cmd *cobra.Command, args []string) error {
		if remoteFlag == "" && !localFlag {
			return fmt.Errorf("you must specify either --remote or --local")
		}
		return nil
	}

}

func helperlingerFlagLocal() {
	logx.L.Debugf("Enabling linger for current user")
	if _, err := util.EnableLinger(); err != nil {
		logx.L.Debugf("%s", err)
		return
	}
	logx.L.Debugf("Enabled linger for current local user")
}

func helperlingerFlagRemote() {
	logx.L.Debugf("Enabling linger for current remote user")
	cli := `luc util oservice linger --local --force`
	if _, err := util.RunCLIRemote(remoteFlag, cli); err != nil {
		logx.L.Debugf("%s", err)
		return
	}
	logx.L.Debugf("Enabled linger for current remote user")
}
