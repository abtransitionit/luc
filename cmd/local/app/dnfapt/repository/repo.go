/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package repository

import (
	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/spf13/cobra"
)

// Description
var repositorySDesc = "Manage repositories."
var repositoryLDesc = repositorySDesc + ` xxx.`

// root Command
var RepositoryCmd = &cobra.Command{
	Use:   "repo",
	Short: repositorySDesc,
	Long:  repositoryLDesc,
	Run: func(cmd *cobra.Command, args []string) {
		logx.L.Debugf(repositorySDesc)
	},
}
