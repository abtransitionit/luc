/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/

package netx

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"

	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/abtransitionit/luc/pkg/util"
)

func IsSshConfiguredVmSshReachable(vmName string) (bool, error) {
	logx.L.Debug("checks whether a VM configured in ~/.ssh/config.d/ is SSH reachable")

	// decare and/or define var
	var out bytes.Buffer
	var shellCli = ""

	// prerequisit: VM is configured in ~/.ssh/config.d/
	_, err := IsVmSshConfigured(vmName)
	if err != nil {
		return false, err
	}
	// Here: the VM is properly configured in ~/.ssh/config.d/
	// define CLI that helps to answer the function question
	shellCli = fmt.Sprintf("ssh %s true", vmName)
	// logx.L.Debugf("CLI to play > %s > ", shellCli)

	// prepare CLI to play
	shellCmd := exec.Command("sh", "-c", shellCli)

	// intermediate variable
	shellCmd.Stdout = &out

	// play CLI
	err = shellCmd.Run()

	// handle FAILURE
	if err != nil {
		err := fmt.Errorf("❌ Error > occured: %v with command: %s", err, shellCmd)
		logx.L.Debugf("❌ Error > command failed > %s", shellCmd)
		return false, err
	}

	// Here: No Failure: CLI provide a functional result (no output or error)
	// result := strings.TrimSpace(out.String())
	// logx.L.Debugf("ssh acces to VM  > %s > %s ", vmName, result)

	// Here: the VM is configured and SSH reachable
	logx.L.Debugf("✅ Cool > VM > %s > is configured in  ~/.ssh/config.d/ and ssh reachable", vmName)
	return true, nil
}

// checks whether a VM is configured in ~/.ssh/config.d/
func IsVmSshConfigured(vmName string) (bool, error) {
	logx.L.Debug("checks whether a VM is configured in ~/.ssh/config.d/")

	// decare and/or define var
	var out bytes.Buffer
	var cliSshName = "ssh"
	var shellCli = ""

	// prerequisite: CLI is available
	if !util.CliExist(cliSshName) {
		err := fmt.Errorf("❌ Error > required CLI not found: %s", cliSshName)
		logx.L.Warnf("❌ Error > Prerequisite check failed : CLI availability of %s", cliSshName)
		return false, err
	}

	// define CLI that helps to answer the function question
	shellCli = fmt.Sprintf("ssh -G %s 2>/dev/null |  grep ^hostname | tr -s ' ' | cut -d' ' -f2", vmName)
	// logx.L.Debugf("CLI to play > %s > ", shellCli)

	// prepare CLI to play
	shellCmd := exec.Command("sh", "-c", shellCli)

	// intermediate variable
	shellCmd.Stdout = &out

	// play CLI
	err := shellCmd.Run()

	// handle FAILURE
	if err != nil {
		err := fmt.Errorf("❌ Error > occured: %v with command: %s", err, shellCmd)
		logx.L.Debugf("❌ Error > command failed > %s", shellCmd)
		return false, err
	}

	// Here: No Failure: CLI provide a functional result: get it
	hostname := strings.TrimSpace(out.String())

	// if the hostname is the same as the VM name, the VM is not configured
	if hostname == vmName {
		logx.L.Debugf("❌ Error : VM name provided > %s > is not configured in ssh config (hostname found > %s)", vmName, hostname)
		return false, nil
	}

	// Here: the VM is configured
	logx.L.Debugf("✅ Cool > VM > %s > is potentially configured in ssh config (%s)", vmName, hostname)
	return true, nil
}
