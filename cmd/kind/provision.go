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
var provisionSDesc = "deploy a Kind cluster on a VM."
var provisionLDesc = provisionSDesc + ` xxx.`

// provision Command
var provisionCmd = &cobra.Command{
	Use:   "provision [phase name]",
	Short: provisionSDesc,
	Long:  provisionLDesc,
	// define the set of phases for this cmd
	Run: phase.CmdRun(kind.Phases, provisionSDesc),
}

func init() {
	phase.CmdInit(provisionCmd)
}
