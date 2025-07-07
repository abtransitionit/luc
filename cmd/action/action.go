/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package action

import (
	"fmt"

	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/abtransitionit/luc/pkg/util"
	"github.com/spf13/cobra"
)

// Description
var actionSDesc = "play functions locally or remotely."
var ActionLDesc = actionSDesc + ` xxx.`

// root Command
var ActionCmd = &cobra.Command{
	Hidden: true, // available but not visible
	Use:    "action",
	Short:  actionSDesc,
	Long:   ActionLDesc,
	Run: func(cmd *cobra.Command, args []string) {
		logx.L.Infof("%s", actionSDesc)
		// Handle --show flag
		showFlag, _ := cmd.Flags().GetBool("show")
		if showFlag {
			util.ShowActionMap()
			return
		}

		// No action provided: show help
		if len(args) == 0 {
			cmd.Help()
			return
		}

		// get vm
		vmName, _ := cmd.Flags().GetString("remote")
		// get action
		action := args[0]
		logx.L.Debugf("action: %s", action)
		// get params
		param1 := ""
		param2 := ""
		param3 := ""
		// Assign only if present
		if len(args) > 1 {
			param1 = args[1]
			logx.L.Debugf("param1: %s", param1)
		}
		if len(args) > 2 {
			param2 = args[2]
			logx.L.Debugf("param2: %s", param2)
		}
		if len(args) > 3 {
			param3 = args[3]
			logx.L.Debugf("param3: %s", param3)
		}
		//
		var result string
		var err error

		if vmName != "" {
			// Remote execution
			result, err = util.PlayActionRemote(vmName, action, param1, param2, param3)
		} else {
			// Local execution
			result, err = util.PlayActionLocal(action, param1, param2, param3)
		}

		// error
		if err != nil {
			logx.L.Debugf("[%s] ❌ %v", action, err)
			return
		}

		// success
		fmt.Println(result)

	},
}

func init() {
	// phase.CmdInit(getpropCmd)
	ActionCmd.Flags().BoolP("show", "s", false, "List available property name")
	ActionCmd.Flags().StringP("remote", "r", "", "Remote VM name")

}
