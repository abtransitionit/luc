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
var cobraSDesc = "Install the cli cobra-cli."
var cobraLDesc = cobraSDesc + ` xxx.`

// root Command
var cobraCmd = &cobra.Command{
	Use:   "cobra [phase name]",
	Short: cobraSDesc,
	Long:  cobraLDesc,
	Run:   deploy.SharedRun(gocli.Phases, cobraSDesc),
}

func init() {
	deploy.SharedInit(cobraCmd)
}
