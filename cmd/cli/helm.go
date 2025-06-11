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
var helmSDesc = "download source code, then build and install the CLI."
var helmLDesc = helmSDesc + ` xxx.`

// root Command
var helmCmd = &cobra.Command{
	Use:   "helm",
	Short: helmSDesc,
	Long:  helmLDesc,
	Run:   deploy.SharedRun(gocli.Phases, helmSDesc),
}

func init() {
	deploy.SharedInit(helmCmd)
}
