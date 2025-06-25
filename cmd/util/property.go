/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package util

import (
	"github.com/abtransitionit/luc/internal/phase/kind"
	"github.com/abtransitionit/luc/pkg/phase"
	"github.com/spf13/cobra"
)

// Description
var propertySDesc = "Manage some OS property."
var propertyLDesc = propertySDesc + ` xxx.`

// provision Command
var propertyCmd = &cobra.Command{
	Use:   "property [phase name]",
	Short: propertySDesc,
	Long:  propertyLDesc,
	// define the set of phases for this cmd
	Run: phase.CmdRun(kind.Phases, propertySDesc),
}

func init() {
	// phase.CmdInit(propertyCmd)
	propertyCmd.Flags().BoolP("show", "s", false, "show CLI config map")

}

// GetPropertyMap()
