/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/

package dnfapt

import (
	"fmt"
	"strings"

	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/abtransitionit/luc/pkg/util"
)

// # Purpose
//
// remote upgrade a Linux OS dnfapt package and package repositories to version latest
func RUpgrade(vm string, osFamily string) (bool, error) {
	var cli = ""
	// check arg
	if vm == "" {
		return false, fmt.Errorf("vm is not provided")
	}
	if osFamily == "" {
		return false, fmt.Errorf("osFamily is not provided")
	}
	//
	switch strings.TrimSpace(osFamily) {
	case "debian":
		cli = "DEBIAN_FRONTEND=noninteractive sudo apt-get -o Dpkg::Options::='--force-confdef' -o Dpkg::Options::='--force-confold' update -qq -y > /dev/null && DEBIAN_FRONTEND=noninteractive sudo apt-get -o Dpkg::Options::='--force-confdef' -o Dpkg::Options::='--force-confold' upgrade -qq -y > /dev/null && DEBIAN_FRONTEND=noninteractive sudo apt-get -o Dpkg::Options::='--force-confdef' -o Dpkg::Options::='--force-confold' -qq clean > /dev/null"
	case "rhel", "fedora":
		cli = "sudo dnf update -q -y > /dev/null && sudo dnf upgrade -q -y  > /dev/null && sudo dnf clean all > /dev/null"
	default:
		return false, fmt.Errorf("unsupported Linux OS Family: %s", osFamily)
	}
	// Play CLI
	outp, err := util.RunCLIRemote(vm, cli)
	if err != nil {
		return false, fmt.Errorf("%v : %s", err, outp)
	}

	// on SUCCESS
	return true, nil
}

func Upgrade() (bool, error) {
	var cli = ""
	// get Li nux Os family
	osFamily, err := util.GetPropertyLocal("osfamily")
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

// # Purpose
//
// - remote install 1..1 dnfapt package
func RInstallP(vm string, osFamily string, packageName string) (bool, error) {
	var cli = ""
	// check arg
	if vm == "" {
		return false, fmt.Errorf("❌ Error: vm is empty")
	}
	if osFamily == "" {
		return false, fmt.Errorf("❌ Error: osFamily is empty")
	}
	//
	switch strings.TrimSpace(osFamily) {
	case "debian":
		cli = fmt.Sprintf(
			"DEBIAN_FRONTEND=noninteractive sudo apt-get -o Dpkg::Options::='--force-confdef' -o Dpkg::Options::='--force-confold' install -qq -y %s > /dev/null && DEBIAN_FRONTEND=noninteractive sudo apt-get -o Dpkg::Options::='--force-confdef' -o Dpkg::Options::='--force-confold' update -qq -y > /dev/null",
			packageName)
	case "rhel", "fedora":
		cli = fmt.Sprintf(
			"sudo dnf install -q -y %s > /dev/null && sudo dnf update -q -y > /dev/null",
			packageName)
	default:
		return false, fmt.Errorf("❌ Error/Warning: unsupported Linux OS Family: %s", osFamily)
	}
	// Play CLI
	out, err := util.RunCLIRemote(vm, cli)
	if err != nil {
		// return false, fmt.Errorf(" ❌ play cli > %s : %v", cli, err)
		return false, fmt.Errorf("%v > %s", err, out)
	}

	// on SUCCESS
	return true, nil
}

// # Purpose
//
// - remote install 1..1 dnfapt repository
func RInstallR(vm string, osFamily string, packageName string) (bool, error) {
	var cli = ""
	// check arg
	if vm == "" {
		return false, fmt.Errorf("❌ Error: vm is empty")
	}
	if osFamily == "" {
		return false, fmt.Errorf("❌ Error: osFamily is empty")
	}
	//
	switch strings.TrimSpace(osFamily) {
	case "debian":
		cli = fmt.Sprintf(
			"DEBIAN_FRONTEND=noninteractive sudo apt-get -o Dpkg::Options::='--force-confdef' -o Dpkg::Options::='--force-confold' install -qq -y %s > /dev/null && DEBIAN_FRONTEND=noninteractive sudo apt-get -o Dpkg::Options::='--force-confdef' -o Dpkg::Options::='--force-confold' update -qq -y > /dev/null",
			packageName)
	case "rhel", "fedora":
		cli = fmt.Sprintf(
			"sudo dnf install -q -y %s > /dev/null && sudo dnf update -q -y > /dev/null",
			packageName)
	default:
		return false, fmt.Errorf("❌ Error/Warning: unsupported Linux OS Family: %s", osFamily)
	}
	// Play CLI
	_, err := util.RunCLIRemote(vm, cli)
	if err != nil {
		return false, fmt.Errorf(" ❌ play cli > %s : %v", cli, err)
	}

	// on SUCCESS
	return true, nil
}

func AddRepo(RepoName string) (string, error) {

	// get Os:family
	osType, err := util.GetPropertyLocal("ostype")
	if err != nil {
		return "", fmt.Errorf("%v : %s", err, osType)
	}

	// get Os:family
	osFamily, err := util.GetPropertyLocal("osfamily")
	if err != nil {
		return "", fmt.Errorf("%v : %s", err, osFamily)
	}

	// install only if linux type
	if osType != "linux" {
		return "", fmt.Errorf("unsupported OS type: %s (only linux is supported)", osType)
	}

	switch strings.TrimSpace(strings.ToLower(osFamily)) {
	case "debian":
		// TODO
	case "rhel", "fedora":
		// TODO
	}

	logx.L.Infof("[%s] AddRepo > for osType/osFamily: %s/%s", RepoName, osType, osFamily)
	// on SUCCESS
	return RepoName, nil
}

func GetPackage(cli string) (string, error) {
	logx.L.Infof("GetPackage > %s", cli)
	logx.L.Infof("GetPackage > return the package that hold the CLI")
	return "", nil
}
