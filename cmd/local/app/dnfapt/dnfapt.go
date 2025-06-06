/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package dnfapt

import (
	"github.com/abtransitionit/luc/cmd/local/app/dnfapt/packagex"
	"github.com/abtransitionit/luc/cmd/local/app/dnfapt/repository"
	"github.com/spf13/cobra"
)

// Description
var dnfaptSDesc = "manage linux OS packages and repositories (both Rhel and Debian based)."
var dnfaptLDesc = dnfaptSDesc + ` xxx.`

// root Command
var DnfaptCmd = &cobra.Command{
	Use:   "dnfapt",
	Short: dnfaptSDesc,
	Long:  dnfaptLDesc,
}

func init() {
	DnfaptCmd.AddCommand(packagex.PackageCmd)
	DnfaptCmd.AddCommand(repository.RepositoryCmd)
}
