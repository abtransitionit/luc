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
var getpropSDesc = "Get/Display an OS property."
var getpropLDesc = getpropSDesc + ` xxx.`

// provision Command
var getpropCmd = &cobra.Command{
	Use:   "getprop [PROPERTY NAME]",
	Short: getpropSDesc,
	Long:  getpropLDesc,
	Run: func(cmd *cobra.Command, args []string) {

		// Handle --show flag
		showFlag, _ := cmd.Flags().GetBool("show")
		if showFlag {
			util.ShowMapProperty()
			return
		}

		// Handle --remote flag
		remoteFlag, _ := cmd.Flags().GetString("remote")

		// manage arg
		if len(args) < 1 {
			cmd.Help()
			return
		}

		// get the property
		propertyName := args[0]

		// get the property parameters if any
		parameters := []string{}
		if len(args) > 1 {
			parameters = args[1:]
		}

		var propertyValue string
		var err error
		if remoteFlag != "" {
			// cli remote
			propertyValue, err = util.GetPropertyRemote(remoteFlag, propertyName, parameters...)
		} else {
			// cli local
			propertyValue, err = util.GetPropertyLocal(propertyName, parameters...)
		}

		// error
		if err != nil {
			logx.L.Debugf("%s", err)
			return
		}
		// success
		fmt.Println(propertyValue)

	},
}

func init() {
	// phase.CmdInit(getpropCmd)
	getpropCmd.Flags().BoolP("show", "s", false, "List available property name")
	getpropCmd.Flags().StringP("remote", "r", "", "Remote VM name")

}

// GetOsPropertyMap()
