/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package cli

import (
	"github.com/abtransitionit/luc/internal/pipeline/gocli"
	"github.com/abtransitionit/luc/pkg/deploy"
	"github.com/spf13/cobra"
)

// Description
var kubectlSDesc = "download binary and install it."
var kubectlLDesc = kubectlSDesc + ` xxx.`

// root Command
var kubectlCmd = &cobra.Command{
	Use:   "kubectl",
	Short: kubectlSDesc,
	Long:  kubectlLDesc,
	Run:   deploy.SharedRun(gocli.Phases, kubectlSDesc),
}

func init() {
	deploy.SharedInit(kubectlCmd)
}
