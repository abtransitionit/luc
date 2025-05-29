/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/

package netx

import (
	"bytes"
	"fmt"
	"net"
	"os/exec"
	"strings"
	"time"

	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/abtransitionit/luc/pkg/util"
	"go.uber.org/zap"
	// "github.com/kevinburke/ssh_config"
)

// IsSshReachable checks if a host is reachable on port 22 (SSH).
func IsSshReachable(host string) bool {
	timeout := 3 * time.Second
	address := net.JoinHostPort(host, "22")
	conn, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		return false
	}
	conn.Close()
	return true
}

// func ResolveHostname(alias string) string {
// 	host, _ := ssh_config.Get(alias, "Hostname")
// 	if host == "" {
// 		return alias
// 	}
// 	return host
// }

// func VmIsSshReachable(alias string) bool {
// 	address := net.JoinHostPort(ResolveHostname(alias), "22")
// 	conn, err := net.DialTimeout("tcp", address, 3*time.Second)
// 	if err != nil {
// 		return false
// 	}
// 	defer conn.Close()
// 	return true
// }

var funcPurpose = "checks whether a VM is configured in ~/.ssh/config.d/ before any remote action on it"

// checks whether a VM is configured in ~/.ssh/config.d/ before any remote action on it
func IsVmSshConfigured(vmName string) (bool, error) {
	logx.L.Debugf(funcPurpose)

	// decare and/or define var
	var out bytes.Buffer
	var cliSshName = "nerdctl"
	var shellCli = ""

	// prerequisite: CLI is available
	if !util.CliExist(cliSshName) {
		err := fmt.Errorf("❌ Error > required CLI not found: %s", cliSshName)
		logx.L.Warn("❌ Error > Prerequisite check failed ")
		logx.L.Warn(zap.String("requirement", "CLI availability"))
		logx.L.Warn(zap.String("cli", cliSshName))
		logx.L.Warn(zap.Error(err))
		// logx.L.Warn("Missing dependency", zap.String("cli", cliSshName), zap.Error(err))
		// logx.L.Debugf("CLI > %s > is NOT available", cliSshName)
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
		logx.L.Debugf("❌ Error > command failed > %v", err)
		return false, nil
	}

	// Here: No Failure: CLI provide a functional output
	hostname := strings.TrimSpace(out.String())
	logx.L.Debugf("VM name provided > %s - hostname found   > %s ", vmName, hostname)

	// if the hostname is the same as the VM name, the VM is not configured
	if hostname == vmName {
		logx.L.Debugf("❌ Error : VM > %s > is not configured in ssh config (%s)", vmName, hostname)
		return false, nil
	}
	// handle SUCCESS: Here: the VM is configured
	logx.L.Debugf("✅ Cool > VM > %s > is configured in ssh config (%s)", vmName, hostname)
	return true, nil
}

// func VmIsSshReachableOld(vmName string) (output bool, customErr string, err error) {

// 	// Define the CLI to get the hostname of the VM
// 	cli := fmt.Sprintf(`ssh -G %s 2>/dev/null | awk '/^hostname / {print $2}'`, vmName)

// 	// Play CLI and handle FAILURE
// 	logx.L.Debugf("👉 [%s] Get hostname", vmName)
// 	output, _, _, errSrc := RunCLILocal(cli)
// 	if errSrc != nil {
// 		customErr = fmt.Sprintf("❌ Error : Failed to get hostname for VM %s: %w", vmName, errSrc)
// 		logx.L.Errorf(customErr)
// 		return "", customErr, errSrc
// 	}

// 	// Clean up the hostname by removing extra spaces
// 	hostname := strings.Join(strings.Fields(output), "")
// 	logx.L.Debugf("👉 [%s] found hostname : %s", vmName, hostname)

// 	// Check if VM is defined in SSH config
// 	if hostname == vmName {
// 		// NO
// 		customErr = fmt.Sprintf("❌ Error : VM > %s > is not configure in ssh config (%s)", vmName, hostname)
// 		err = fmt.Errorf(customErr)
// 		logx.L.Errorf(customErr)
// 		return "", customErr, err
// 	}
// 	// Here the VM is defined in the SSH config : check if the VM is configured
// 	cli = fmt.Sprintf("ssh %s true", vmName)
// 	logx.L.Debugf("🔎 check SSH access")
// 	output, _, _, errSrc = RunCLILocal(cli)
// 	if errSrc != nil {
// 		customErr = fmt.Sprintf("❌ Error : VM > %s > IP (%s) exist but is not configured: %v", vmName, hostname, errSrc)
// 		logx.L.Errorf("❌ [%s] VM is not SSH reachable", vmName)
// 		return "", customErr, errSrc
// 	}
// 	customOutput := fmt.Sprintf("✅ VM is SSH reachable (%s)", hostname)
// 	logx.L.Debugf("👉 [%s] VM is SSH reachable", vmName)
// 	return customOutput, "", nil
// }
