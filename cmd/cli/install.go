/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package cli

import (
	"github.com/abtransitionit/luc/pkg/config"
	"github.com/spf13/cobra"
)

// Description
var installSDesc = "install a GO CLI."
var installLDesc = installSDesc + ` xxx.`

// root Command
var installCmd = &cobra.Command{
	Use:   "install [cli name] [cli version] [Destination folder]",
	Short: installSDesc,
	Long:  installLDesc,
	Run: func(cmd *cobra.Command, args []string) {

		// handle flag = --show
		if cmd.Flag("show").Value.String() == "true" {
			config.ShowCliConfigMap()
			return
		}

		// // handle arguments
		// if len(args) == 3 {
		// 	// get CLI name
		// 	cliName := args[0]
		// 	cliVersion := args[1]
		// 	cliFolder := args[2]
		// 	// Create a proper map with the configuration
		// 	configMap := map[string]config.CustomCLIConfig{
		// 		cliName: {
		// 			Name:      cliName,
		// 			Version:   cliVersion,
		// 			DstFolder: cliFolder,
		// 		},
		// 	}

		// 	// Launch pipeline for that CLI
		// 	gocli.RunPipeline(configMap)
		// 	return
		// } else if len(args) > 1 {
		// 	logx.L.Error("❌ Needed 3 arguments")
		// 	return
		// }

		// fallback
		cmd.Help()
	},
}

func init() {
	installCmd.Flags().BoolP("show", "s", false, "show CLI config map")
}
