/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/

package util

import (
	"fmt"
)

// computes the Cartesian product of two string slices.
//
// Returns:
//
//   - []string: containing the Cartesian product
//   - error:    if either input slice is empty
//
// Time complexity: O(n*m) where n = len(a), m = len(b)
// Space complexity: O(n*m)
func Reboot() (string, error) {
	// play CLI
	cli := "sudo reboot"
	output, err := RunCLILocal(cli)
	if err != nil {
		return "", fmt.Errorf("❌ Error rebooting : %v", err)
	}
	return output, nil
}
