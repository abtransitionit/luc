/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package gox

import (
	"github.com/abtransitionit/luc/pkg/config"
	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/abtransitionit/luc/pkg/pipeline/gocli"
	"github.com/spf13/cobra"
)

// Description
var isdkSDesc = "install the GO chaintool/sdk (`go` binary and  libraries to start coding in go)."
var isdkLDesc = isdkSDesc + ` xxx.`

// root Command
var isdkCmd = &cobra.Command{
	Use:   "isdk",
	Short: isdkSDesc,
	Long:  isdkLDesc,
	Run: func(cmd *cobra.Command, args []string) {
		logx.L.Debugf(isdkSDesc)
		listVm := "o3r"
		GoCliCustomConfigMap := config.CustomCLIConfigMap{
			"go": {
				Name:      "go",
				Version:   "1.24.4",
				DstFolder: "/usr/local",
			},
		}

		// launch a pipeleine
		_, err := gocli.RunPipeline(listVm, GoCliCustomConfigMap)
		if err != nil {
			logx.L.Errorf("%s", err)
			return
		}
	},
}
