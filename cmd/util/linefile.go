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
var linefileSDesc = "Append line to end of file."
var linefileLDesc = linefileSDesc + ` xxx.`

// provision Command
var lineFileCmd = &cobra.Command{
	Use:   "linefile [STRING] [FILE PATH]",
	Short: linefileSDesc,
	Long:  linefileLDesc,
	RunE: func(cmd *cobra.Command, args []string) error {

		// Handle --remote flag
		remoteFlag, _ := cmd.Flags().GetString("remote")

		// manage arg
		if len(args) < 1 {
			cmd.Help()
			return nil
		}

		// manage flag - foce flag is mandatory to use this command
		if !forceFlag {
			return fmt.Errorf("use --force to run this command")
		}

		// get raw inputs
		fileContent := args[0]
		filePath := args[1]

		var err error

		// Play cli
		if remoteFlag != "" {
			// cli remote
			cli := fmt.Sprintf(`luc util linefile %q %s --force`, fileContent, filePath)
			_, err = util.RunCLIRemote(remoteFlag, cli)
		} else {
			// cli local
			_, err = util.AddLineToFile(filePath, fileContent)
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
	lineFileCmd.Flags().BoolVar(&forceFlag, "force", false, "Force execution is mandatory")
	lineFileCmd.Flags().StringP("remote", "r", "", "Remote VM name")

}

// // handle flag = --show
// if cmd.Flag("show").Value.String() == "true" {
// 	util.ShowMapProperty()
// 	return
// } else {
// 	cmd.Help()
// }
