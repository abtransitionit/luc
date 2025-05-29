/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/

package util

import "os/exec"

// check if a CLI is available in the system's PATH.
func CliExist(name string) bool {
	_, err := exec.LookPath(name)
	return err == nil
}
