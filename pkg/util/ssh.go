/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/

package util

import (
	"fmt"
	"strings"

	"github.com/abtransitionit/luc/pkg/errorx"
	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/jedib0t/go-pretty/table"
	"github.com/jedib0t/go-pretty/text"
)

type SshStatus struct {
	Vm            string
	SshConfigured bool
	SshReachable  bool
}

type SshStatusMap map[string]SshStatus

// declare
var NodeSshStatusMap = SshStatusMap{}

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

	// check VM is configured in ~/.ssh/config.d/
	configured, err := IsVmSshConfigured(vmName)
	if err != nil {
		return false, fmt.Errorf("%v : %v", err, configured)
	} else if !configured {
		return false, fmt.Errorf("❌ Error: VM %s is not configured in ~/.ssh/config.d/", vmName)
	}

	// VM is configured : Play CLI
	cli := fmt.Sprintf("ssh %s true", vmName)
	outp, err := RunCLILocal(cli)
	if err != nil {
		return false, fmt.Errorf("%v : %v", err, outp)
	}

	// // handle system FAILURE - TODO: improve to know the real reaseon
	// if err != nil {
	// 	return errorx.BoolError("playing CLI", cli, err)
	// }

	// hostname := strings.TrimSpace(out.String())
	// logx.L.Debugf("✅ vm %s is configured in ssh config and is ssh reachable : %v", vmName, hostname)
	return true, nil
}
func IsVmSshReachable(vmName string) (bool, error) {

	// check VM is configured in ~/.ssh/config.d/
	configured, err := IsVmSshConfigured(vmName)
	if err != nil {
		return false, fmt.Errorf("%v : %v", err, configured)
	} else if !configured {
		return false, fmt.Errorf("❌ Error: VM %s is not configured in ~/.ssh/config.d/", vmName)
	}

	// VM is configured : Play CLI
	cli := fmt.Sprintf("ssh %s true", vmName)
	outp, err := RunCLILocal(cli)
	if err != nil {
		return false, fmt.Errorf("%v : %v", err, outp)
	}
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

const CheckSshDescription = "check VMs are SSH reachable."

// # Purpose
//
// Check if one are more VMs are SSH reachable.
// func CheckSshV2(listVm []string) (string, error) {

// 	// manage argument
// 	if len(listVm) == 0 {
// 		return "", fmt.Errorf("❌ Error : List of VMs is empty")
// 	}

// 	// do the check
// 	for _, vm := range listVm {
// 		test.CheckVmIsSshReachable(vm)
// 	}

// }
func CheckSshV1(arg ...string) (string, error) {
	logx.L.Info(CheckSshDescription)

	// initialize the map
	NodeSshStatusMap = SshStatusMap{}

	// define var
	var SliceNodes []string

	if len(arg) == 1 && strings.Contains(arg[0], " ") {
		SliceNodes = strings.Fields(arg[0]) // convert ListAsString to []string (ie. go slice)
	} else {
		// Already a slice of node names
		SliceNodes = arg
	}

	// check nodes are SSH configured
	logx.L.Info("check vms are SSH configured")
	for _, node := range SliceNodes {
		isSshConfigured, err := IsVmSshConfigured(node)
		if err != nil {
			return "", fmt.Errorf("node: %s: %v", node, err)
		}
		// Fill the map with a structure instance for each Node
		NodeSshStatusMap[node] = SshStatus{
			Vm:            node,
			SshConfigured: isSshConfigured,
			SshReachable:  false, // default value for now
		}

	} // for

	// check nodes are SSH reachable
	logx.L.Info("checking vms are SSH reachable")
	for _, node := range SliceNodes {
		isSssReachable, err := IsSshConfiguredVmSshReachable(node)
		if err != nil {
			logx.L.Debugf("%s", err)
			continue
		}
		// update the map for each node
		nodeStatus := NodeSshStatusMap[node]
		nodeStatus.SshReachable = isSssReachable
		NodeSshStatusMap[node] = nodeStatus

	} // for

	// success
	logx.L.Info("checked vms")
	fmt.Println((NodeSshStatusMap).String())
	return "", nil
}

// pretty display
func (m SshStatusMap) String() string {
	t := table.NewWriter()
	t.SetStyle(table.StyleLight)
	t.Style().Title.Align = text.AlignCenter

	t.SetTitle("SSH Status for Cluster nodes")
	t.AppendHeader(table.Row{"Node", "SSH Configured", "SSH Reachable"})

	for _, status := range m {
		t.AppendRow(table.Row{
			status.Vm,
			status.SshConfigured,
			status.SshReachable,
		})
	}

	return t.Render()
}
