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
var runcSDesc = "download tgz and install the CLI."
var runcLDesc = runcSDesc + ` xxx.`

// root Command
var runcCmd = &cobra.Command{
	Use:   "runc",
	Short: runcSDesc,
	Long:  runcLDesc,
	Run:   deploy.SharedRun(gocli.Phases, runcSDesc),
}
