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
var getSDesc = "donwload an object referenced by an URL into a folder."
var getLDesc = getSDesc + ` xxx.`

// delete Command
var getCmd = &cobra.Command{
	Use:   "get [URL] [PATH]",
	Args:  cobra.ExactArgs(2), // Requires exactly one argument: the URL
	Short: getSDesc,
	Long:  getLDesc,
	// define the set of phases for this cmd
	RunE: func(cmd *cobra.Command, args []string) error {
		logx.L.Infof("%s", getSDesc)

		// get the URL
		url := args[0]
		path := args[1]

		if localFlag {
			// play cli
			_, err := util.GetFile(url, path)

			// errror
			if err != nil {
				return err
			}

			// success
			logx.L.Infof("✅ downloaded object in locale memory.")
			return nil
		}

		if remoteFlag != "" {
			// play cli
			cli := fmt.Sprintf("luc util url get %s %s --local", url, path)
			_, err := util.RunCLIRemote(remoteFlag, cli)

			// error
			if err != nil {
				logx.L.Debugf("❌ Error detected 1")
				return err
			}

			// success
			logx.L.Infof("✅ [%s] downloaded object in remote memory.", remoteFlag)
			return nil
		}
		return nil
	},
}

func init() {
	getCmd.Flags().StringVar(&remoteFlag, "remote", "", "Download remotely from a target host (e.g., o1u)")
	getCmd.Flags().BoolVar(&localFlag, "local", false, "Download locally")
	getCmd.MarkFlagsMutuallyExclusive("local", "remote")
	getCmd.PreRunE = func(cmd *cobra.Command, args []string) error {
		if remoteFlag == "" && !localFlag {
			return fmt.Errorf("you must specify either --remote or --local")
		}
		return nil
	}
}
