/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package cli

import (
	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/spf13/cobra"
)

// Description
var containerdSDesc = "download tgz and install the CLI."
var containerdLDesc = containerdSDesc + ` xxx.`

// root Command
var containerdCmd = &cobra.Command{
	Use:   "containerd",
	Short: containerdSDesc,
	Long:  containerdLDesc,
	Run: func(cmd *cobra.Command, args []string) {
		logx.L.Debugf("%s", containerdSDesc)
	},
}
