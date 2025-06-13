/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package kind

import (
	"github.com/abtransitionit/luc/internal/phase/kind"
	"github.com/abtransitionit/luc/pkg/phase"
	"github.com/spf13/cobra"
)

// Description
var initSDesc = "deploy a Kind cluster."
var initLDesc = initSDesc + ` xxx.`

// init Command
var initCmd = &cobra.Command{
	Use:   "init [phase name]",
	Short: initSDesc,
	Long:  initLDesc,
	Run:   phase.SharedRun(kind.Phases, initSDesc),
}

func init() {
	phase.SharedInit(initCmd)
}
