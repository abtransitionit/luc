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
var utgzSDesc = "unatr a targz file into a folder."
var utgzLDesc = utgzSDesc + ` xxx.`

// delete Command
var utgzCmd = &cobra.Command{
	Use:   "utgz [SRC FILE PATH] [DST FOLDER PATH]",
	Args:  cobra.ExactArgs(2), // Requires exactly 2 arguments
	Short: utgzSDesc,
	Long:  utgzLDesc,
	// define the set of phases for this cmd
	RunE: func(cmd *cobra.Command, args []string) error {
		logx.L.Infof("%s", utgzSDesc)

		// get the URL
		src := args[0]
		dst := args[1]

		if localFlag {
			if err := util.UnTgz(src, dst); err != nil {
				return err
			}
			// success
			logx.L.Infof("✅ untargzed file locally.")
			return nil
		}

		if remoteFlag != "" {
			// paly code
			cli := fmt.Sprintf("luc util utgz %s %s --local", src, dst)
			if _, err := util.RunCLIRemote(remoteFlag, cli); err != nil {
				return err
			}
			// success
			logx.L.Infof("✅ untargzed file remotely.")
			return nil
		}
		return nil
	},
}

func init() {
	utgzCmd.Flags().StringVar(&remoteFlag, "remote", "", "Download remotely from a target host (e.g., o1u)")
	utgzCmd.Flags().BoolVar(&localFlag, "local", false, "Download locally")
	utgzCmd.MarkFlagsMutuallyExclusive("local", "remote")
	utgzCmd.PreRunE = func(cmd *cobra.Command, args []string) error {
		if remoteFlag == "" && !localFlag {
			return fmt.Errorf("you must specify either --remote or --local")
		}
		return nil
	}
}
