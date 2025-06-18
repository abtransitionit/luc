/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/

package dnfapt

import (
	"fmt"
	"strings"

	"github.com/abtransitionit/luc/pkg/util"
)

// update a Linux OS dnfapt package and package repositories to version latest
func UpdateOs() (bool, error) {
	var cli = ""
	// get Li nux Os family
	osFamily, err := util.OsPropertyGet("osfamily")
	if err != nil {
		return false, err
	}
	switch strings.TrimSpace(osFamily) {
	case "debian":
		cli = "DEBIAN_FRONTEND=noninteractive sudo apt-get -o Dpkg::Options::='--force-confdef' -o Dpkg::Options::='--force-confold' update -qq -y > /dev/null && DEBIAN_FRONTEND=noninteractive sudo apt-get -o Dpkg::Options::='--force-confdef' -o Dpkg::Options::='--force-confold' upgrade -qq -y > /dev/null && DEBIAN_FRONTEND=noninteractive sudo apt-get -o Dpkg::Options::='--force-confdef' -o Dpkg::Options::='--force-confold' -qq clean > /dev/null"
	case "rhel", "fedora":
		cli = "sudo dnf update -q -y > /dev/null && sudo dnf upgrade -q -y  > /dev/null && sudo dnf clean all > /dev/null"
	default:
		return false, fmt.Errorf("❌ Error/Warning: unsupported Linux OS Family: %s", osFamily)
	}
	// Play CLI
	_, err = util.RunCLILocal(cli)
	if err != nil {
		return false, fmt.Errorf(" ❌ play cli > %s : %v", cli, err)
	}

	// on SUCCESS
	return true, nil
}
