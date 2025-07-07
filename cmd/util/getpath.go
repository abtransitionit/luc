/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package util

import (
	"fmt"

	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/abtransitionit/luc/pkg/util"
	"github.com/spf13/cobra"
)

// Description
var pathSDesc = "create a PATH envar."
var pathLDesc = pathSDesc + `
Example usage:

luc util path /usr/local/bin;/usr/local/sbin/kind
`

// Definition
var pathCmd = &cobra.Command{
	Use:   "path [semi-colon separated list of paths]",
	Args:  cobra.ExactArgs(1), // Requires exactly one argument: the set of paths
	Short: pathSDesc,
	Long:  pathLDesc,
	// Code to play
	Run: func(cmd *cobra.Command, args []string) {

		// manage flag - foce flag is mandatory to use this command
		if !forceFlag {
			logx.L.Infof("use --force to run this command.")
			logx.L.Infof("also check --help for more details")
			return
		}

		// use cases
		var path string
		var err error
		if localFlag {
			if path, err = helperlingerFlagLocal(args); err != nil {
				logx.L.Debugf("%s", err)
				return
			}
		}
		if remoteFlag != "" {
			if path, err = helperlingerFlagRemote(args); err != nil {
				logx.L.Debugf("%s", err)
				return
			}
		}
		fmt.Println(path)
	},
}

func init() {
	pathCmd.Flags().BoolVar(&forceFlag, "force", false, "Force execution is mandatory")
	pathCmd.Flags().StringVar(&remoteFlag, "remote", "", "Download remotely from a target host (e.g., o1u)")
	pathCmd.Flags().BoolVar(&localFlag, "local", false, "Download locally")
	pathCmd.MarkFlagsMutuallyExclusive("local", "remote")
	pathCmd.PreRunE = func(cmd *cobra.Command, args []string) error {
		if remoteFlag == "" && !localFlag {
			return fmt.Errorf("you must specify either --remote or --local")
		}
		return nil
	}

}

func helperlingerFlagLocal(args []string) (string, error) {
	// get input
	treePath := args[0]

	// play cli
	path, err := util.UpdatePath(treePath)

	// error
	if err != nil {
		logx.L.Debugf("%s", err)
		return "", err
	}

	// success
	return path, nil
}

func helperlingerFlagRemote(args []string) (string, error) {

	// get input
	treePath := args[0]

	// play cli
	cli := fmt.Sprintf(`luc util path treePath --local --force`, treePath)
	path, err := util.RunCLIRemote(remoteFlag, cli)

	// error
	if err != nil {
		logx.L.Debugf("%s", err)
		return "", err
	}

	// success
	return path, nil
}
