/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package kbe

import (
	"github.com/abtransitionit/luc/internal/phase/kbe"
	"github.com/abtransitionit/luc/pkg/phase"
	"github.com/spf13/cobra"
)

// Description
var provisionSDesc = "deploy a kubernetes cluster. It becomes a KBE (Kubernetes Easy) cluster."
var provisionLDesc = provisionSDesc + ` xxx.`

// init Command
var provisionCmd = &cobra.Command{
	Use:   "provision [phase name]",
	Short: provisionSDesc,
	Long:  provisionLDesc,
	Run:   phase.CmdRun(kbe.ProvisionPhases, provisionSDesc),
}

func init() {
	phase.CmdInit(provisionCmd)
}
