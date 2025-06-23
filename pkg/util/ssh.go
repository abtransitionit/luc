/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/

package util

import (
	"fmt"
	"strings"

	"github.com/abtransitionit/luc/pkg/errorx"
	"github.com/abtransitionit/luc/pkg/logx"
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
// _, err := util.IsSshConfiguredVmSshReachable(param)
// if err != nil {
// 	   os.Exit(2)
// }

func IsSshConfiguredVmSshReachable(vmName string) (bool, error) {

	// prerequisit: VM is configured in ~/.ssh/config.d/
	configured, err := IsVmSshConfigured(vmName)
	if err != nil {
		return false, err
	} else if !configured {
		return false, fmt.Errorf("VM %s is not configured in ~/.ssh/config.d/", vmName)
	}

	// Play CLI
	cli := fmt.Sprintf("ssh %s true", vmName)
	_, err = RunCLILocal(cli)
	if err != nil {
		logx.L.Debugf("❌ Error detected 1")
		return false, err
	}

	// handle system FAILURE - TODO: improve to know the real reaseon
	if err != nil {
		return errorx.BoolError("playing CLI", cli, err)
	}

	// hostname := strings.TrimSpace(out.String())
	// logx.L.Debugf("✅ vm %s is configured in ssh config and is ssh reachable : %v", vmName, hostname)
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

	// check ssh client is available
	cliExists, err := CliExists(cliSshName)
	if err != nil {
		return errorx.BoolError("check SSH CLI exists", cliSshName, err)
	}
	if !cliExists {
		return false, fmt.Errorf("not found CLI: ssh")
	}

	// Build CLI that helps to answer the function's question/query
	cli := fmt.Sprintf("ssh -G %s 2>/dev/null | grep ^hostname | tr -s ' ' | cut -d' ' -f2", vmName)
	out, err := RunCLILocal(cli)
	if err != nil {
		logx.L.Debugf("❌ Error detected 2")
		// return errorx.BoolError("check SSH config for vm", vmName, err)
		return false, err
	}

	// success: the hostname
	hostname := strings.TrimSpace(out)
	if hostname == vmName {
		return false, nil // VM not configured (normal case)
	}
	return true, nil // VM is properly configured
}
