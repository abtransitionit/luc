/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package cli

import (
	"github.com/abtransitionit/luc/internal/gocli"
	"github.com/abtransitionit/luc/pkg/deploy"
	"github.com/spf13/cobra"
)

// Description
var rootlesskitSDesc = "download tgz and install the CLI."
var rootlesskitLDesc = rootlesskitSDesc + ` xxx.`

// root Command
var rootlesskitCmd = &cobra.Command{
	Use:   "rootlesskit",
	Short: rootlesskitSDesc,
	Long:  rootlesskitLDesc,
	Run:   deploy.SharedRun(gocli.Phases, rootlesskitSDesc),
}

func init() {
	deploy.SharedInit(rootlesskitCmd)
}
