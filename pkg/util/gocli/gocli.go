/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/

package gocli

import (
	"fmt"

	"github.com/abtransitionit/luc/pkg/config"
	"github.com/abtransitionit/luc/pkg/util"
)

// # Purpose
//
// - remote install 1..1 Go CLI on 1..1 Linux VM
//
// # Parameters
//
// - vm: the VM name as defined in the ~/.ssh/config
// - cliConfig: a pointer to an instance
func RInstallC(vm string, cliConfig *config.CustomCLIConfig) (bool, error) {
	// var cli = ""

	// check arg
	if vm == "" {
		return false, fmt.Errorf("❌ Error: vm is empty")
	}
	if cliConfig == nil {
		return false, fmt.Errorf("❌ Error: cliConfig is empty")
	}

	// check vm
	_, err := util.IsSshConfiguredVmSshReachable(vm)
	if err != nil {
		return false, err
	}

	// get vm:property
	osType, err := util.GetRemoteProperty("ostype", vm)
	if err != nil {
		return false, err
	}

	// check vm:os
	if osType != "linux" {
		return false, fmt.Errorf("❌ Error: osType is '%s', not linux", osType)
	}

	// Before
	// - get other info for this Go CLI
	// 		- FileType
	// 		- UrlGen
	// - convert to Url specific
	// install CLI

	// fmt.Println(cliConfig)
	// logx.L.Info("Installing Go CLI in %s:", vm, cliConfig)
	// // Play CLI
	// _, err := util.RunCLIRemote(cli, vm)
	// if err != nil {
	// 	return false, fmt.Errorf(" ❌ play cli > %s : %v", cli, err)
	// }

	// on SUCCESS
	return true, nil
}
