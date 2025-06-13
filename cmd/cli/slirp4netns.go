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
var slirp4netnsSDesc = "download binary and install it."
var slirp4netnsLDesc = slirp4netnsSDesc + ` xxx.`

// root Command
var slirp4netnsCmd = &cobra.Command{
	Use:   "slirp4netns",
	Short: slirp4netnsSDesc,
	Long:  slirp4netnsLDesc,
	Run:   deploy.SharedRun(gocli.Phases, slirp4netnsSDesc),
}

func init() {
	deploy.SharedInit(slirp4netnsCmd)
}
