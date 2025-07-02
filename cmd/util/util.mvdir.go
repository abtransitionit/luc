/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package util

import (
	"fmt"
	"os"
	"strconv"

	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/abtransitionit/luc/pkg/util"
	"github.com/spf13/cobra"
)

// Description
var mvDirSDesc = "move a folder."
var mvDirLDesc = mvDirSDesc + ` xxx.`

// delete Command
// data.ArtPath1, dstPath, 0755, true
var mvdirCmd = &cobra.Command{
	Use:   "mvdir [SRC PATH] [DST PATH] [PERMISSION] [DO OVERWRITE] [IS ROOT]",
	Args:  cobra.ExactArgs(5), // Requires exactly one argument: the URL
	Short: mvDirSDesc,
	Long:  mvDirLDesc,
	// define the set of phases for this cmd
	RunE: func(cmd *cobra.Command, args []string) error {
		logx.L.Infof("%s", mvDirSDesc)

		// get arguments
		src := args[0]
		dst := args[1]
		permStr := args[2]
		doOverwriteStr := args[3]
		isRootStr := args[3]

		// convert arguments
		permUint64, err := strconv.ParseUint(permStr, 8, 32)
		if err != nil {
			return fmt.Errorf("❌ Error: invalid permission format: %w", err)
		}
		perm := os.FileMode(permUint64)
		doOverwrite, err := strconv.ParseBool(doOverwriteStr)
		if err != nil {
			return fmt.Errorf("❌ Error: invalid doOverwriteStr value: %w", err)
		}
		isRoot, err := strconv.ParseBool(isRootStr)
		if err != nil {
			return fmt.Errorf("❌ Error: invalid isroot value: %w", err)
		}

		// use arguments
		if localFlag {
			success, err := util.MvFolder(src, dst, perm, doOverwrite, isRoot)
			// error
			if err != nil {
				return err
			}
			// error
			if !success {
				return fmt.Errorf("❌ Error: File not moved to '%s'", dst)
			}
			// success
			logx.L.Infof("✅ moved file locally.")
			return nil
		}

		if remoteFlag != "" {
			cli := fmt.Sprintf("luc util mvdir %s %s %s %s %s --local", src, dst, permStr, doOverwriteStr, isRootStr)
			if _, err := util.RunCLIRemote2(cli, remoteFlag); err != nil {
				return err
			}
			// success
			logx.L.Infof("✅ moved file remotely.")
			return nil
		}
		return nil
	},
}

func init() {
	mvdirCmd.Flags().StringVar(&remoteFlag, "remote", "", "Download remotely from a target host (e.g., o1u)")
	mvdirCmd.Flags().BoolVar(&localFlag, "local", false, "Download locally")
	mvdirCmd.MarkFlagsMutuallyExclusive("local", "remote")
	mvdirCmd.PreRunE = func(cmd *cobra.Command, args []string) error {
		if remoteFlag == "" && !localFlag {
			return fmt.Errorf("you must specify either --remote or --local")
		}
		return nil
	}
}

// func HandleLocal(url string) error {

// 	// download
// 	_, err := util.GetPublicFile(url)
// 	if err != nil {
// 		return err
// 	}

// }

// func HandleRemote(url, host string) error {
// 	cli := fmt.Sprintf("luc util url get %s --local", url)

// 	// check vm is ssh reachable
// 	_, err := util.IsSshConfiguredVmSshReachable(host)
// 	if err != nil {
// 		return err
// 	}

// 	// remote download
// 	_, err = util.RunCLIRemote2(cli, host)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
