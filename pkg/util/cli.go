/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/

package util

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
	"strings"

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

// # Purpose
//
// runs a local linux shell CLI and returns its stdout, stderr, and any error.
//
// Parameters:
//   - command: A shell command string (e.g., "ls -l /").
//
// Returns:
//   - stdout: The trimmed standard output of the command.
//   - err:    An error if the command failed to run or returned a non-zero exit code.
//
// Usage:
//
//	output, err := RunCLILocal("hostname")
//	if err != nil {
//	    fmt.Printf("Error: %v\n", err)
//	} else {
//	    fmt.Println("Hostname:", output)
//	}
func RunCLILocal(command string) (stdout string, err error) {
	cmd := exec.Command("bash", "-c", command)

	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out

	err = cmd.Run()
	stdout = strings.TrimSpace(out.String())

	if err != nil {
		return stdout, fmt.Errorf("command failed: %v\noutput:\n%s", err, stdout)
	}

	return stdout, nil
}

func RunCLILocalOld01(command string) (stdout string, err error) {
	cmd := exec.Command("bash", "-c", command)

	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out // still capture errors in output for debugging

	err = cmd.Run()
	stdout = strings.TrimSpace(out.String())

	return stdout, err
}

// RunCLIRemote runs a shell command on a remote machine via SSH.
func RunCLIRemote(vm string, command string) (stdout string, err error) {
	// Format SSH command: ssh user@host "command"
	fullCmd := fmt.Sprintf(`ssh %s "%s"`, vm, command)

	// cmd := exec.Command("bash", "-c", fullCmd)
	cmd := exec.Command("sh", "-c", fullCmd)

	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out

	err = cmd.Run()
	stdout = strings.TrimSpace(out.String())

	if err != nil {
		return stdout, fmt.Errorf("remote command failed: %v\noutput:\n%s", err, stdout)
	}

	return stdout, nil
}
