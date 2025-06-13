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
var lucSDesc = "download tgz and install the CLI."
var lucLDesc = lucSDesc + ` xxx.`

// root Command
var lucCmd = &cobra.Command{
	Use:   "luc",
	Short: lucSDesc,
	Long:  lucLDesc,
	Run:   deploy.SharedRun(gocli.Phases, lucSDesc),
}

func init() {
	deploy.SharedInit(lucCmd)
}
