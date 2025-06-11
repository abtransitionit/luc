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
var containerdSDesc = "download tgz and install the CLI."
var containerdLDesc = containerdSDesc + ` xxx.`

// root Command
var containerdCmd = &cobra.Command{
	Use:   "containerd",
	Short: containerdSDesc,
	Long:  containerdLDesc,
	Run:   deploy.SharedRun(gocli.Phases, containerdSDesc),
}

func init() {
	deploy.SharedInit(containerdCmd)
}
