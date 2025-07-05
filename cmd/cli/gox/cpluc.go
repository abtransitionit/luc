/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package gox

import (
	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/abtransitionit/luc/pkg/pipeline/cpluc"
	"github.com/spf13/cobra"
)

// Description
var cplucSDesc = "install LUC binary on a VM."
var cplucLDesc = cplucSDesc + ` do:
- build luc locally for current platform
- deploy it locally
- copy the binary for platform linux/amd64 to the VM.
`

// root Command
var cplucCmd = &cobra.Command{
	Use:   "cpluc [VM1] [VM2] ...",
	Short: cplucSDesc,
	Long:  cplucLDesc,
	Run: func(cmd *cobra.Command, args []string) {
		logx.L.Debugf(cplucSDesc)

		// get the list of VMs from arg
		listVm := ""
		for _, vm := range args {
			listVm += " " + vm
		}

		// force flag is mandatory to use this command
		if !forceFlag {
			logx.L.Infof("use --force to run this command.")
			logx.L.Infof("also check --help for more details")
			return
		}

		// launch this pipeline
		_, err := cpluc.RunPipeline(listVm)
		if err != nil {
			logx.L.Debugf("%s", err)
			return
		}
	},
}

func init() {
	cplucCmd.Flags().BoolVarP(&forceFlag, "force", "f", false, "Bypass confirmation and force execution")
}
