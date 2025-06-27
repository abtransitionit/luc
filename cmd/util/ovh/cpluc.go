/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package ovh

import (
	"fmt"
	"strings"

	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/abtransitionit/luc/pkg/util"
	"github.com/spf13/cobra"
)

// Description
var cplucSDesc = "build and deploy LUC in local then copy it on remote:"
var cplucLDesc = cplucSDesc + `
    - build LUC go CLI locally for current OS/platform (where the command is run)
    - locally deploy LUC to /usr/local/bin on the current OS/platform
    - build LUC go CLI locally for other platforms (i.e. OVH VMs :linux/amd64)
    - copy this LUC go CLI extra platform to remote OVH VM(s).
	`

var CplucCmd = &cobra.Command{
	Use:   "cpluc OvhVm1 OvhVm2 ...",
	Short: cplucSDesc,
	Long:  cplucLDesc,
	Run: func(cmd *cobra.Command, args []string) {
		logx.L.Infof("%s", cplucSDesc)
		if !force {
			logx.L.Infof("use --force to run this command. also check --help for more details")
			return
		}

		// define CLI to build LUC locally and for platform linux/amd64 (ie. OVH VMs)
		outputCurrentPtfFilePath := "/tmp/luc"
		outputLinuxAmdFilePath := "/tmp/luc-linux"
		cli := fmt.Sprintf(`
		cd /var/tmp/luc 							&& 
		rm -rf /tmp/luc* &> /dev/null &&  
		go build -o %s 					&& 
		sudo mv %s /usr/local/bin/luc && 
		GOOS=linux GOARCH=amd64 go build -o %s && cd -
		`, outputCurrentPtfFilePath, outputCurrentPtfFilePath, outputLinuxAmdFilePath)
		logx.L.Debug("building LUC locally")

		// play CLI
		_, err := util.RunCLILocal(cli)
		if err != nil {
			logx.L.Debugf("%s", err)
		}
		logx.L.Debug("builded LUC locally")

		// args denote OVH VMs
		if len(args) == 0 {
			return
		}

		// build ListAsString from args (ie. OvhVm1 OvhVm2 ...)
		var listVm = ""
		listVm = strings.Join(args, " ")
		// copy LUC to OVH VMs
		logx.L.Debugf("copy LUC to OVH VMs: %s", listVm)

	},
}

var force bool

func init() {
	CplucCmd.Flags().BoolVar(&force, "force", false, "Force execution even if not needed")
}
