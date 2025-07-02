/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package util

import (
	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/spf13/cobra"
)

// Description
var urlSDesc = "manage URLs."
var urlLDesc = urlSDesc + ` xxx.`

// delete Command
var urlCmd = &cobra.Command{
	Use:   "url",
	Short: urlSDesc,
	Long:  urlLDesc,
	// define the set of phases for this cmd
	Run: func(cmd *cobra.Command, args []string) {
		logx.L.Infof("%s", urlSDesc)
		cmd.Help()
	},
}

func init() {
	urlCmd.AddCommand(getCmd)
}

// list ovh vm
// cp luc to ovh vm provided as arg
