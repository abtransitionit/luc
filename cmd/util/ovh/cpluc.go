/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package ovh

import (
	"strings"

	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/abtransitionit/luc/pkg/pipeline/cpluc"
	"github.com/spf13/cobra"
)

// Description
var cplucSDesc = "build and deploy LUC (using a pipeline)."
var cplucLDesc = cplucSDesc + `
    - build LUC locally for the current OS/platform (where the command is run)
    - locally deploy LUC to /usr/local/bin on the current OS/platform
    - build LUC locally for another OS/platform (i.e. OVH VMs :linux/amd64)
    - copy this last build concurently to remote OVH VM(s).
	`

// provision Command
var CplucCmd = &cobra.Command{
	Use:   "cpluc OvhVm1 OvhVm2 ...",
	Short: cplucSDesc,
	Long:  cplucLDesc,
	// define the set of phases for this cmd
	Run: func(cmd *cobra.Command, args []string) {
		logx.L.Infof("%s", cplucSDesc)
		// foce is mandatory to use this command
		if !force {
			logx.L.Infof("use --force to run this command. also check --help for more details")
			return
		}

		// each args denote an OVH VMs on which LUC will be deployed
		if len(args) == 0 {
			return
		}

		// build ListAsString from args (ie. OvhVm1 OvhVm2 ...)
		var listVm = ""
		listVm = strings.Join(args, " ")

		// cpluc.RunPipeline(config.KbeListNode)
		cpluc.RunPipeline(listVm)

	},
}

var force bool

func init() {
	// phase.CmdInit(CplucCmd)
	CplucCmd.Flags().BoolVar(&force, "force", false, "Force execution even if not needed")
}
