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
var mvFileSDesc = "move a file."
var mvFileLDesc = mvFileSDesc + ` xxx.`

// delete Command
// data.ArtPath1, dstPath, 0755, true
var mvfileCmd = &cobra.Command{
	Use:   "mvfile [SRC PATH] [DST PATH] [PERMISSION] [IS ROOT]",
	Args:  cobra.ExactArgs(4), // Requires exactly one argument: the URL
	Short: mvFileSDesc,
	Long:  mvFileLDesc,
	// define the set of phases for this cmd
	RunE: func(cmd *cobra.Command, args []string) error {
		logx.L.Infof("%s", mvFileSDesc)

		// get arguments
		src := args[0]
		dst := args[1]
		permStr := args[2]
		isRootStr := args[3]

		// convert arguments
		permUint64, err := strconv.ParseUint(permStr, 8, 32)
		if err != nil {
			return fmt.Errorf("❌ Error: invalid permission format: %w", err)
		}
		perm := os.FileMode(permUint64)
		isRoot, err := strconv.ParseBool(isRootStr)
		if err != nil {
			return fmt.Errorf("❌ Error: invalid isroot value: %w", err)
		}

		// use arguments
		if localFlag {
			success, err := util.MvFile2(src, dst, perm, isRoot)
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
			cli := fmt.Sprintf("luc util mvfile %s %s %s %s --local", src, dst, permStr, isRootStr)
			if _, err := util.RunCLIRemote(remoteFlag, cli); err != nil {
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
	mvfileCmd.Flags().StringVar(&remoteFlag, "remote", "", "Download remotely from a target host (e.g., o1u)")
	mvfileCmd.Flags().BoolVar(&localFlag, "local", false, "Download locally")
	mvfileCmd.MarkFlagsMutuallyExclusive("local", "remote")
	mvfileCmd.PreRunE = func(cmd *cobra.Command, args []string) error {
		if remoteFlag == "" && !localFlag {
			return fmt.Errorf("you must specify either --remote or --local")
		}
		return nil
	}
}
