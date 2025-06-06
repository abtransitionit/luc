/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/

package netx

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"

	"github.com/abtransitionit/luc/pkg/errorx"
	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/abtransitionit/luc/pkg/util"
)

// IsSshConfiguredVmSshReachable checks if a VM is both:
//  1. Configured in SSH config (~/.ssh/config.d/)
//  2. Currently reachable via SSH
//
// Prerequisites:
//   - VM must be configured in SSH config (checked via IsVmSshConfigured)
//   - SSH CLI must be available
//
// Returns:
//   - (true, nil)  if VM is properly configured and reachable
//   - (false, nil) if VM is configured but unreachable
//   - (false, error) for system failures or prerequisite violations
//
// Error Handling:
//   - Logs errors using logx (caller doesn't need to repeat logging)
//   - Returns wrapped errors with context
//
// Example:
//
//	reachable, err := IsSshConfiguredVmSshReachable("my-vm")
//	if err != nil {
//	    // handle system/configuration errors
//	}
//	if !reachable {
//	    // handle unreachable VM case
//	}
//
// Example:
//
// _, err := netx.IsSshConfiguredVmSshReachable(param)
// if err != nil {
// 	   os.Exit(2)
// }

func IsSshConfiguredVmSshReachable(vmName string) (bool, error) {
	// logx.L.Info("checks whether a VM configured in ~/.ssh/config.d/ is SSH reachable")

	// decare and/or define var
	var out bytes.Buffer

	// prerequisit: VM is configured in ~/.ssh/config.d/
	configured, err := IsVmSshConfigured(vmName)
	if err != nil {
		// handle system FAILURE
		msg := fmt.Sprintf("failed to check SSH config for VM: %s : %v", vmName, err)
		logx.L.Debugf("❌ %s", msg)
		return false, err
		// handle applogic FAILURE
	} else if !configured {
		msg := fmt.Sprintf("VM %s is not configured in ~/.ssh/config.d/", vmName)
		logx.L.Debugf("❌ %s", msg)
		return false, fmt.Errorf("VM %s is not configured in ~/.ssh/config.d/", vmName)
	}
	// handle applogic SUCCESS
	msg := fmt.Sprintf("VM %s is configured in ~/.ssh/config.d/", vmName)
	logx.L.Debugf("✅ %s", msg)

	// prepare CLI to play : define CLI that helps to answer the function question
	shellCli := fmt.Sprintf("ssh %s true", vmName)
	shellCmd := exec.Command("sh", "-c", shellCli)
	// intermediate variable
	shellCmd.Stdout = &out
	// play CLI
	err = shellCmd.Run()

	// handle system FAILURE - TODO: improve to know the real reaseon
	if err != nil {
		msg := fmt.Sprintf("failed to play CLI %s to check vm %s is SSH reachable (or VM is not SSH reachable)", shellCli, vmName)
		logx.L.Debugf("❌ %s", msg)
		return errorx.BoolError("playing CLI", shellCli, err)
	}

	// handle applogic SUCCESS
	msg = fmt.Sprintf("VM %s is SSH reachable", vmName)
	logx.L.Debugf("✅ %s", msg)

	hostname := strings.TrimSpace(out.String())
	logx.L.Debugf("✅ vm %s is potentially configured in ssh consig and is ssh reachable : %v", vmName, hostname)
	return true, nil
}

// IsVmSshConfigured checks if a VM is properly configured in the SSH config.
//
// It checks:
//  1. The SSH CLI is available
//  2. The VM has a dedicated hostname configuration in SSH config
//
// Returns:
//   - (true, nil)  if VM is properly configured in SSH config
//   - (false, nil) if VM is not configured (normal case)
//   - (false, error) if system error occurs (e.g., permission issues, SSH CLI check failed)
//
// Example:
//
//		configured, err := IsVmSshConfigured("my-vm")
//		if err != nil {
//		    // handle system error
//		}
//		if !configured {
//		    // handle unconfigured VM case
//		}
//	 ...  // handle configured VM case
//
// Note: Function intentionally only returns and does not include logging. Caller should handle logging if needed based on their context
// Note: Function only return.
func IsVmSshConfigured(vmName string) (bool, error) {
	var cliSshName = "ssh"
	var out bytes.Buffer

	// prerequisite: ssh client is available : critical error
	cliExists, err := util.CliExists(cliSshName)
	// handle FAILURE
	if err != nil {
		return errorx.BoolError("check SSH CLI exists", cliSshName, err)
	}
	// handle SUCCESS : a boolean
	if !cliExists {
		return false, nil // SSH not available → treat as "not configured"
	}

	// Here: ssh cli exists
	// Build CLI that helps to answer the function's question
	cmd := fmt.Sprintf("ssh -G %s 2>/dev/null | grep ^hostname | tr -s ' ' | cut -d' ' -f2", vmName)
	// Prepare the CLI
	shellCmd := exec.Command("sh", "-c", cmd)
	// Play the CLI
	shellCmd.Stdout = &out

	// handle system FAILURE
	if err := shellCmd.Run(); err != nil {
		return errorx.BoolError("check SSH config for vm", vmName, err)
	}

	// handle SUCCESS: the hostname
	hostname := strings.TrimSpace(out.String())
	if hostname == vmName {
		return false, nil // VM not configured (normal case)
	}
	return true, nil // VM is properly configured
}
