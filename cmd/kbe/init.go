/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package kbe

import (
	"github.com/abtransitionit/luc/internal/phase/kbe"
	"github.com/abtransitionit/luc/pkg/deploy"
	"github.com/spf13/cobra"
)

// Description
var initSDesc = "deploy a kubernetes cluster. It becomes a KBE (Kubernetes Easy) cluster."
var initLDesc = initSDesc + ` xxx.`

// init Command
var initCmd = &cobra.Command{
	Use:   "init [phase name]",
	Short: initSDesc,
	Long:  initLDesc,
	Run:   deploy.SharedRun(kbe.Phases, initSDesc),
}

func init() {
	deploy.SharedInit(initCmd)
}
