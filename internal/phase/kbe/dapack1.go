/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package kbe

import (
	"fmt"
	"strings"

	"github.com/abtransitionit/luc/internal/config"
	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/abtransitionit/luc/pkg/pipeline/dnfapt/packagex"
	"github.com/abtransitionit/luc/pkg/util"
	"github.com/abtransitionit/luc/test"
)

const DaPackStdDescription = "provision standard/required/missing OS CLI (via dnfapt  packages)."

func daPackStd(arg ...string) (string, error) {
	logx.L.Info(DaPackStdDescription)

	vms := util.GetSlicefromStringWithSpace(config.KbeListNode)

	// install mising standard dnfapt clis (per osDistro)
	for _, vm := range vms {

		// check vm is ssh reachable
		vmReachabiliy, err := util.GetPropertyLocal("sshreachability", vm)
		if err != nil {
			return "", fmt.Errorf("%v : %v", err, vmReachabiliy)
		}
		if vmReachabiliy != "true" {
			logx.L.Debugf("❌ [%s] : %s", vm, "skiping cause vm not reachable")
			continue
		}

		// get os:distro
		osDistro, err := util.GetPropertyRemote(vm, "osdistro")
		if err != nil {
			return "", err
		}

		// define dnfapt packages depending on os:distro - cliList is used to check package provisioning
		var packageList, cliList string
		switch strings.TrimSpace(osDistro) {
		case "debian":
			packageList = "gnupg"
			cliList = "gpg"
		}

		// add packages only if packageList is not empty for this VM
		if packageList != "" {
			// provision dnfapt packages
			_, err = packagex.RunPipeline(vm, strings.Fields(packageList))
			if err != nil {
				logx.L.Debugf("%s", err)
				return "", err
			}

			// check cli exists after install
			for _, cliName := range strings.Fields(cliList) {
				test.CheckCliExistsOnremote(vm, cliName)
			}
		} // if
	} // for

	// success
	return "", nil

}

// case "almalinux":
// 	packageList = "dnf-utils python3-dnf-plugin-versionlock"
// 	// purpose     = "provision CLI needs-restarting versionlock"
// case "rocky":
// 	packageList = "python3-dnf-plugin-versionlock"
// 	// purpose     = "provision CLI versionlock"  ;;
// case "fedora":
// 	packageList = "dnf-utils"
// 	// purpose     = "provision CLI needs-restarting"  ;;
