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
var nerdctlSDesc = "download tgz and install the CLI."
var nerdctlLDesc = nerdctlSDesc + ` xxx.`

// root Command
var nerdctlCmd = &cobra.Command{
	Use:   "nerdctl",
	Short: nerdctlSDesc,
	Long:  nerdctlLDesc,
	Run:   deploy.SharedRun(gocli.Phases, nerdctlSDesc),
}
