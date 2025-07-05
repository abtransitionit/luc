/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package gox

import (
	"runtime"

	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/abtransitionit/luc/pkg/util/gocli"
	"github.com/spf13/cobra"
)

var (
	pathFlag string
	outFlag  string
)

// Description
var buildSDesc = "Building the binary from a GO project folder"

// "Building Locally (for the current platform) a Go binary from a local GIT project folder."

var buildLDesc = buildSDesc + `: 
- Build the Go binary locally for the current platform

Example usage:

luc cli go build --path /var/tmp/luc --out /tmp/luc --force

luc cli go build --path /var/tmp/luc --out /tmp/luc-linux-amd64 --ostype linux --osarch amd64 --force
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

		// define platform
		osType := osTypeFlag
		if osType == "" {
			osType = runtime.GOOS
		}
		osArch := osArchFlag
		if osArch == "" {
			osArch = runtime.GOARCH
		}

		// Build the CLI
		binaryPath, err := gocli.GoBuildXPtf(pathFlag, outFlag, osType, osArch)

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
	buildCmd.Flags().StringVarP(&pathFlag, "path", "p", "", "Absolute folder path to the go project folder (containing the file go.mod)")
	buildCmd.Flags().StringVar(&outFlag, "out", "", "Absolute file path of the output binary")
	buildCmd.Flags().BoolVarP(&forceFlag, "force", "f", false, "Mandatory to use this command")
	buildCmd.Flags().StringVar(&osTypeFlag, "ostype", "", "Target OS for cross-compilation (e.g., linux, windows, darwin)")
	buildCmd.Flags().StringVar(&osArchFlag, "osarch", "", "Target architecture for cross-compilation (e.g., amd64, arm64)")
}
