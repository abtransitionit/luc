/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package cli

import (
	"github.com/abtransitionit/luc/internal/pipeline/gocli"
	"github.com/abtransitionit/luc/pkg/config"
	"github.com/spf13/cobra"
)

// Description
var installSDesc = "install a GO CLI."
var installLDesc = installSDesc + ` xxx.`

// root Command
var installCmd = &cobra.Command{
	Use:   "install [cli name]",
	Short: installSDesc,
	Long:  installLDesc,
	Run: func(cmd *cobra.Command, args []string) {

		// handle flag = --show
		if cmd.Flag("show").Value.String() == "true" {
			config.ShowConfigMap()
			return
		}

		// handle arguments
		if len(args) == 1 {
			// get CLI name
			cliName := args[0]
			// Launch pipeline for that CLI
			gocli.Ep11(cliName)
			return
		}

		// fallback
		cmd.Help()
	},
}

func init() {
	installCmd.Flags().BoolP("show", "s", false, "show CLI config map")
}
