/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package gox

import (
	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/abtransitionit/luc/pkg/util"
	"github.com/spf13/cobra"
)

var (
	pathFlag string
	outFlag  string
)

// Description
var buildSDesc = "Building a GO project locally."

// "Building Locally (for the current platform) a Go binary from a local GIT project folder."

var buildLDesc = buildSDesc + ` 
- for the current platform
- from a local GIT folder.

Example usage:

luc cli go build --path /var/tmp/luc --out /tmp/toto --force
`

// root Command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: buildSDesc,
	Long:  buildLDesc,
	Run: func(cmd *cobra.Command, args []string) {
		logx.L.Debugf(buildSDesc)

		// Validate flags
		if pathFlag == "" || outFlag == "" {
			logx.L.Debug("Both --path and --out flags are required")
			cmd.Help()
			return
		}

		// force flag is mandatory to use this command
		if !forceFlag {
			logx.L.Infof("use --force to run this command.")
			logx.L.Infof("also check --help for more details")
			return
		}

		// Build the CLI
		binaryPath, err := util.GoBuild(pathFlag, outFlag)

		// error
		if err != nil {
			logx.L.Debugf("%s", err)
			return
		}
		// success
		logx.L.Infof("✅ created binary from Go project '%s' at '%s'", pathFlag, binaryPath)
	},
}

func init() {
	buildCmd.Flags().StringVarP(&pathFlag, "path", "p", "", "Absolute path to the go project folder (containing the file go.mod)")
	buildCmd.Flags().StringVar(&outFlag, "out", "", "Absolute path of the output binary")
	buildCmd.Flags().BoolVarP(&forceFlag, "force", "f", false, "Bypass confirmation and force execution")
}

// buildCmd.Flags().StringVar(&archFlag, "arch", "", "The cpu arch: must be arm or amd")
// buildCmd.Flags().StringVar(&osFlag, "os", "", "The OS type must be linux or darwin")
