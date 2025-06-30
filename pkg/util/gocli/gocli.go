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
// - remote install 1..1 go CLI
func RInstallC(vm string, osType string, cliConfig config.CLIConfig) (bool, error) {
	var cli = ""
	// check arg
	if vm == "" {
		return false, fmt.Errorf("❌ Error: vm is empty")
	}
	if osType == "" || osType != "linux" {
		return false, fmt.Errorf("❌ Error: osFamily is empty or is not linux")
	}
	// Play CLI
	_, err := util.RunCLIRemote2(cli, vm)
	if err != nil {
		return false, fmt.Errorf(" ❌ play cli > %s : %v", cli, err)
	}

	// on SUCCESS
	return true, nil
}
