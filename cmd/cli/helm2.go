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
var helm2SDesc = "download tgz and install the CLI."
var helm2LDesc = helm2SDesc + ` xxx.`

// root Command
var helm2Cmd = &cobra.Command{
	Use:   "helm2",
	Short: helm2SDesc,
	Long:  helm2LDesc,
	Run:   deploy.SharedRun(gocli.Phases, helm2SDesc),
}

func init() {
	deploy.SharedInit(helm2Cmd)
}
