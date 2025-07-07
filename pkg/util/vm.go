/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/

package util

import (
	"fmt"
	"os/exec"
)

// # Purpose
//
// - Reboots the current OS
//
// Returns:
//
//   - []string: containing the Cartesian product
//   - error:    if either input slice is empty
func Reboot() (string, error) {
	// play CLI
	cli := "sudo reboot"
	output, err := RunCLILocal(cli)
	if err != nil {
		return "", fmt.Errorf("❌ Error rebooting : %v", err)
	}
	return output, nil
}

// # Purpose
//
// - Just lauch a reboot on a remote VM
//
// Returns:
//
//   - error:    if any error occurs
func RemoteReboot(vm string) error {
	// check arg
	if vm == "" {
		return fmt.Errorf("❌ Error: vm is empty")
	}

	// check VM is reachable
	_, err := IsSshConfiguredVmSshReachable(vm)
	if err != nil {
		return err
	}

	// remote reboot
	cmd := exec.Command("ssh", vm, "sudo reboot")
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("❌ Failed to send reboot command to '%s': %v", vm, err)
	}

	// success
	return nil
}
