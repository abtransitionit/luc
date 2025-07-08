/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package do

import (
	"fmt"

	"github.com/abtransitionit/luc/pkg/util"
	"github.com/spf13/cobra"
)

// Description
var doSDesc = "play functions locally or remotely."
var doLDesc = doSDesc + ` xxx.`

// root Command
var DoCmd = &cobra.Command{
	Hidden: true, // available but not visible
	Use:    "do",
	Short:  doSDesc,
	Long:   doLDesc,
	RunE: func(cmd *cobra.Command, args []string) error {
		// logx.L.Infof("%s", doSDesc)

		// Handle --show flag
		showFlag, _ := cmd.Flags().GetBool("show")
		if showFlag {
			util.ShowFnActionMap()
			return nil
		}

		// No action provided: show help
		if len(args) == 0 {
			cmd.Help()
			return nil
		}

		// get vm
		vmName, _ := cmd.Flags().GetString("remote")
		// get action
		action := args[0]
		// logx.L.Debugf("action: %s", action)
		// get params
		param1 := ""
		param2 := ""
		param3 := ""
		// Assign only if present
		if len(args) > 1 {
			param1 = args[1]
			// logx.L.Debugf("param1: %s", param1)
		}
		if len(args) > 2 {
			param2 = args[2]
			// logx.L.Debugf("param2: %s", param2)
		}
		if len(args) > 3 {
			param3 = args[3]
			// logx.L.Debugf("param3: %s", param3)
		}
		//
		var result string
		var err error

		if vmName != "" {
			// Remote execution
			result, err = util.PlayFnOnRemote(vmName, action, param1, param2, param3)
		} else {
			// Local execution
			result, err = util.PlayFnLocally(action, param1, param2, param3)
		}

		// error
		if err != nil {
			// logx.L.Debugf("[%s] %v", action, err)
			// logx.L.Debugf("[%s] ❌ error detected", action)
			return err
		}

		// success
		fmt.Println(result)
		return nil
	},
}

func init() {
	DoCmd.AddCommand(getpropCmd)
	DoCmd.Flags().BoolP("show", "s", false, "List available property name")
	DoCmd.Flags().StringP("remote", "r", "", "Remote VM name")

}
