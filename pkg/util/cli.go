/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/

package util

import (
	"errors"
	"os/exec"

	"github.com/abtransitionit/luc/pkg/errorx"
)

// checks if a CLI tool is available in the system's PATH.
// Returns:
//   - (true, nil) if the tool exists
//   - (false, nil) if the tool is not found (not treated as an error)
//   - (false, error) if PATH lookup fails (e.g., permission issues)
//
// Usage examples:
//
//		if exists, err := CliExist("go"); err != nil {
//			log.Fatalf("PATH lookup failed: %v", err)
//		} else if !exists {
//		  log.Println("Go toolchain not found - please install Go first")
//		}
//
//		// Check for Docker CLI
//		if exists, _ := CliExist("docker"); exists {
//		    fmt.Println("Docker is available")
//		} else {
//		    fmt.Println("Docker not found in PATH")
//		}
//	 Note: Function intentionally only returns and does not include logging. Caller should handle logging if needed based on their context`
func CliExists(name string) (bool, error) {
	_, err := exec.LookPath(name)
	if err == nil {
		return true, nil // Tool exists
	}
	if errors.Is(err, exec.ErrNotFound) {
		return false, nil // Tool missing (normal case)
	}
	// handle system FAILURE
	return errorx.BoolError("check cli exists", name, err)
}
