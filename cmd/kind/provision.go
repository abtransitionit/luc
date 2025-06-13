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
var provisionSDesc = "deploy a Kind cluster."
var provisionLDesc = provisionSDesc + ` xxx.`

// provision Command
var provisionCmd = &cobra.Command{
	Use:   "provision [phase name]",
	Short: provisionSDesc,
	Long:  provisionLDesc,
	Run:   phase.SharedRun(kind.Phases, provisionSDesc),
}

func init() {
	phase.SharedInit(provisionCmd)
}
