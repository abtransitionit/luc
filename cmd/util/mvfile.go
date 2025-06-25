/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package util

import (
	"os"
	"strconv"

	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/abtransitionit/luc/pkg/util"
	"github.com/spf13/cobra"
)

// Description
var mvfileSDesc = "Move file from one folder to another."
var mvfileLDesc = mvfileSDesc + ` xxx.`

// provision Command
var mvfileCmd = &cobra.Command{
	Use:   "mvfile [Src File] [Dst Folder] [Permission] [IsRoot]",
	Short: mvfileSDesc,
	Long:  mvfileLDesc,
	Run: func(cmd *cobra.Command, args []string) {

		// handle arguments
		if len(args) == 4 {
			srcFilePath := args[0]
			dstFolderPath := args[1]
			n, _ := strconv.ParseUint(args[2], 8, 0)
			permission := os.FileMode(n)
			isRoot, _ := strconv.ParseBool(args[3])
			_, err := util.MvFile(srcFilePath, dstFolderPath, permission, isRoot)
			if err != nil {
				logx.L.Debugf("%s", err)
				return
			}
			return
		}

		// handle flag = --show
		if cmd.Flag("show").Value.String() == "true" {
			util.ShowPropertyMap()
			return
		} else {
			cmd.Help()
		}
	},
}

func init() {
	// phase.CmdInit(mvfileCmd)
	mvfileCmd.Flags().BoolP("show", "s", false, "List available property name")

}

// GetPropertyMap()
