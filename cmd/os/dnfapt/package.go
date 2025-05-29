/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package dnfapt

import (
	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/spf13/cobra"
)

// Description
var packageSDesc = "Deploy a kubernetes cluster. It becomes a KBE (Kubernetes Easy) cluster."
var packageLDesc = packageSDesc + ` xxx.`

// root Command
var packageCmd = &cobra.Command{
	Use:   "package",
	Short: packageSDesc,
	Long:  packageLDesc,
	Run: func(cmd *cobra.Command, args []string) {
		logx.L.Infof("%s", packageSDesc)
	},
}
