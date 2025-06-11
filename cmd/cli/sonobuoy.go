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
var sonobuoySDesc = "download source code, then build and install the CLI."
var sonobuoyLDesc = sonobuoySDesc + ` xxx.`

// root Command
var sonobuoyCmd = &cobra.Command{
	Use:   "sonobuoy",
	Short: sonobuoySDesc,
	Long:  sonobuoyLDesc,
	Run:   deploy.SharedRun(gocli.Phases, sonobuoySDesc),
}

func init() {
	deploy.SharedInit(sonobuoyCmd)
}
