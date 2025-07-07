/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package cmd

import (
	"os"

	"github.com/abtransitionit/luc/cmd/cli"
	"github.com/abtransitionit/luc/cmd/kbe"
	"github.com/abtransitionit/luc/cmd/kind"
	"github.com/abtransitionit/luc/cmd/util"
	"github.com/spf13/cobra"
)

// Description
var rootSDesc = "LUC (aka. Linux Unified CLI) is a user-friendly, auto-documented command-line interface."
var rootLDesc = rootSDesc + ` It simplifies daily tasks for DevOps engineers and developers by providing a unified and consistent CLI experience. LUC can, for example:
	→ Manage containers and container images,
	→ Manage Linux OS packages and repositories using a unified interface — no need to worry about whether it's apt or dnf nor if it's debian, fedora or ubuntu
	→ Manage remote VM objects,
	→ Simplify the creation and management of Kubernetes clusters across virtual machines,
	→ ...and much more.

As a Linux cross-distribution CLI, LUC is also well-suited and ready for full automation and integration into any CI/CD pipelines.`

// root Command
var rootCmd = &cobra.Command{
	Use:   "luc",
	Short: rootSDesc,
	Long:  rootLDesc,
}

// called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.AddCommand(util.UtilCmd)
	rootCmd.AddCommand(cli.CliCmd)
	rootCmd.AddCommand(kbe.KbeCmd)
	rootCmd.AddCommand(kind.KindCmd)
	// rootCmd.AddCommand(local.LocalCmd)
	//
	rootCmd.AddCommand(testCmd)
}
