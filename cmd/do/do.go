/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package do

import (
	"fmt"

	"github.com/abtransitionit/luc/cmd/do/oservice"
	"github.com/abtransitionit/luc/pkg/action"
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
			action.ShowFnActionMap()
			return nil
		}

		// No action provided: show help
		if len(args) == 0 {
			cmd.Help()
			return nil
		}

		// define var from input
		vmName, _ := cmd.Flags().GetString("remote")

		// define var
		theAction := args[0]
		parameters := args[1:] // Pass all elements except the first one

		// does theAction exits
		if _, ok := action.FnActionMap[theAction]; !ok {
			return fmt.Errorf("❌ Error : theAction does not exist : %s", theAction)
		}

		// Play cli
		var result string
		var err error

		if vmName != "" {
			// Remote execution
			result, err = action.PlayFnOnRemote(vmName, theAction, parameters)
		} else {
			// Local execution
			result, err = action.PlayFnLocally(theAction, parameters)
		}

		// error
		if err != nil {
			// fmt.Println(result)
			return err
		}

		// success
		fmt.Println(result)
		return nil
	},
}

func init() {
	DoCmd.SilenceUsage = true // do not show usage on error
	// DoCmd.SilenceErrors = true
	DoCmd.AddCommand(getpropCmd)
	DoCmd.AddCommand(oservice.OsServiceCmd)
	DoCmd.Flags().BoolP("show", "s", false, "List available property name")
	DoCmd.Flags().StringP("remote", "r", "", "Remote VM name")
}
