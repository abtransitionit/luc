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
var getpropSDesc = "Manage some OS property."
var getpropLDesc = getpropSDesc + ` xxx.`

// provision Command
var getpropCmd = &cobra.Command{
	Use:   "getprop [property name]",
	Short: getpropSDesc,
	Long:  getpropLDesc,
	Run: func(cmd *cobra.Command, args []string) {

		// handle arguments
		if len(args) == 1 {
			propertyName := args[0]
			propertyValue, err := util.GetLocalProperty(propertyName)
			if err != nil {
				logx.L.Debugf("%s", err)
				return
			}
			// logx.L.Infof("value for property '%s': %s", propertyName, propertyValue)
			fmt.Println(propertyValue)
			return
		}

		// handle flag = --show
		if cmd.Flag("show").Value.String() == "true" {
			util.ShowMapProperty()
			return
		} else {
			cmd.Help()
		}
	},
}

func init() {
	// phase.CmdInit(getpropCmd)
	getpropCmd.Flags().BoolP("show", "s", false, "List available property name")

}

// GetOsPropertyMap()
