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
var sdkSDesc = "Install the standard GO libraries and the cli go to start coding."
var sdkLDesc = sdkSDesc + ` xxx.`

// root Command
var sdkCmd = &cobra.Command{
	Use:   "sdk",
	Short: sdkSDesc,
	Long:  sdkLDesc,
	Run:   deploy.SharedRun(gocli.Phases, sdkSDesc),
}

func init() {
	deploy.SharedInit(sdkCmd)
}
