/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package util

import (
	"fmt"
	"strconv"

	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/abtransitionit/luc/pkg/util"
	"github.com/spf13/cobra"
)

// Description
var strfileSDesc = "Create file from string."
var strfileLDesc = strfileSDesc + ` xxx.`

// provision Command
var strFileCmd = &cobra.Command{
	Use: "strfile [STRING] [FILE PATH] [IS ROOT]",
	// Args:  cobra.ExactArgs(3), // Requires exactly 2 argument
	Short: strfileSDesc,
	Long:  strfileLDesc,
	RunE: func(cmd *cobra.Command, args []string) error {

		// Handle --show flag
		showFlag, _ := cmd.Flags().GetBool("show")
		if showFlag {
			util.ShowMapProperty()
			return nil
		}

		// Handle --remote flag
		remoteFlag, _ := cmd.Flags().GetString("remote")

		// manage arg
		if len(args) < 3 {
			cmd.Help()
			return nil
		}

		// manage flag - foce flag is mandatory to use this command
		if !forceFlag {
			logx.L.Infof("use --force to run this command.")
			logx.L.Infof("also check --help for more details")
			return fmt.Errorf("use --force to run this command.")
		}

		// get raw inputs
		fileContent := args[0]
		filePtah := args[1]
		IsRootRaw := args[2]

		var err error
		// convert inputs
		IsRoot, err := strconv.ParseBool(IsRootRaw)
		if err != nil {
			logx.L.Debugf("%s", err)
			return err
		}

		// Play cli
		if remoteFlag != "" {
			// cli remote
			cli := fmt.Sprintf(`luc util strfile %s %s %s --force`, fileContent, filePtah, IsRootRaw)
			_, err = util.RunCLIRemote(remoteFlag, cli)
		} else {
			// cli local
			_, err = util.SaveStringToFile(fileContent, filePtah, IsRoot)
		}

		// error
		if err != nil {
			logx.L.Debugf("%s", err)
			return err
		}
		return nil
	},
}

func init() {
	strFileCmd.Flags().BoolVar(&forceFlag, "force", false, "Force execution is mandatory")
	strFileCmd.Flags().BoolP("show", "s", false, "List available property name")
	strFileCmd.Flags().StringP("remote", "r", "", "Remote VM name")

}

// // handle flag = --show
// if cmd.Flag("show").Value.String() == "true" {
// 	util.ShowMapProperty()
// 	return
// } else {
// 	cmd.Help()
// }
