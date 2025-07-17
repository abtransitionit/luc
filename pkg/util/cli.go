/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/

package util

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// # Purpose
//
// checks if a CLI is available in the system's PATH.
//
// # Returns
//   - (true, nil) if the tool exists
//   - (false, nil) if the tool is not found (not treated as an error)
//   - (false, error) if PATH lookup fails (e.g., permission issues)
//
// # Usage examples
//
// localExists, err := CliExists("luc")
//
//	if err != nil {
//		 log.Fatalf("error checking luc locally: %v", err)
//	}
//
// fmt.Println("luc available locally:", localExists)
func CliExists(name string) (bool, error) {
	_, err := exec.LookPath(name)
	if err != nil {
		if errors.Is(err, exec.ErrNotFound) {
			return false, nil // Tool is not installed (normal case)
		}
		return false, err // Unexpected error (e.g., permission issue)
	}
	return true, nil
}

// # Purpose
//
// checks if a CLI is available on a remote VM.
//
// # Example usage
//
//	exists, err := CliExistsRemote("luc", "o1u")
//	if err != nil {
//		log.Fatalf("error checking luc on remote 'o1u': %v", err)
//	}
//
// fmt.Println("luc available on o1u:", remoteExists)
func CliRemoteExists(name, vm string) (bool, error) {
	cmd := fmt.Sprintf("command -v %s", name)
	output, err := RunCLIRemote(vm, cmd)
	if err != nil {
		// Any SSH or command failure (not found) returns false
		return false, nil
	}
	return strings.TrimSpace(output) != "", nil
}

func RunCLILocal(command string, liveOutput ...bool) (string, error) {
	live := len(liveOutput) > 0 && liveOutput[0]
	cmd := exec.Command("bash", "-c", command)

	// Live output mode
	if live {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			// return "", fmt.Errorf("❌ Error: command failed: %v", err)
			return "", fmt.Errorf("❌ %v", err)
		}
		return "", nil
	}

	// Silent mode
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out

	err := cmd.Run()
	output := strings.TrimSpace(out.String())
	// error
	if err != nil {
		// return stdout, fmt.Errorf("❌ Error: command failed: %v\noutput:\n%s", err, stdout)
		return output, fmt.Errorf("%v", err)
	}
	// success
	return output, nil
}

// func RunCLILocalOLD(command string, liveOutput ...bool) (stdout string, err error) {
// 	// Set default value (false if not provided)
// 	live := false
// 	if len(liveOutput) > 0 {
// 		live = liveOutput[0]
// 	}

// 	cmd := exec.Command("bash", "-c", command)

// 	if live {
// 		// Live output mode - show in terminal
// 		cmd.Stdout = os.Stdout
// 		cmd.Stderr = os.Stderr

// 		err = cmd.Run()
// 		if err != nil {
// 			// return "", fmt.Errorf("❌ Error: command failed: %v", err)
// 			return stdout, fmt.Errorf("❌ %v", err)
// 		}
// 		return "", nil // No captured output in live mode
// 	} else {
// 		// Silent mode - capture output (original behavior)
// 		var out bytes.Buffer
// 		cmd.Stdout = &out
// 		cmd.Stderr = &out

// 		err = cmd.Run()
// 		stdout = strings.TrimSpace(out.String())

// 		if err != nil {
// 			// return stdout, fmt.Errorf("❌ Error: command failed: %v\noutput:\n%s", err, stdout)
// 			return stdout, fmt.Errorf("❌ %v", err)

// 		}
// 		return stdout, nil
// 	}
// }

func RunCLIRemote(vm, command string, liveOutput ...bool) (string, error) {
	live := len(liveOutput) > 0 && liveOutput[0]
	// fullCmd := fmt.Sprintf(`ssh %s '%s'`, vm, command)
	encodedCmd := base64.StdEncoding.EncodeToString([]byte(command))
	fullCmd := fmt.Sprintf(`ssh %s "echo '%s' | base64 --decode | sh"`, vm, encodedCmd)
	cmd := exec.Command("sh", "-c", fullCmd)
	// logx.L.Debugf("⚠️⚠️ Running on remote cli: %s", command)

	// Live output mode
	if live {
		// var stderr bytes.Buffer
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		// cmd.Stderr = io.MultiWriter(os.Stderr, &stderr) // capture + print

		if err := cmd.Run(); err != nil {
			// return "", fmt.Errorf("%v:%s", err, stderr.String())
			return "", fmt.Errorf("%v", err)
		}
		return "", nil
	}

	// Silent mode
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out

	err := cmd.Run()
	output := strings.TrimSpace(out.String())
	if err != nil {
		// return stdout, fmt.Errorf("❌ Error: remote command failed: %v\noutput:\n%s", err, stdout)
		return output, fmt.Errorf("%v", err)
	}
	return output, nil
}

// func RunCLIRemoteOLD(vm string, command string, liveOutput ...bool) (stdout string, err error) {
// 	// Set default value for liveOutput
// 	live := false
// 	if len(liveOutput) > 0 {
// 		live = liveOutput[0]
// 	}

// 	// Format SSH command: ssh user@host "command"
// 	// '%s' allow to not expand $XXX to local variable when exist
// 	fullCmd := fmt.Sprintf(`ssh %s '%s'`, vm, command)

// 	cmd := exec.Command("sh", "-c", fullCmd)

// 	if live {
// 		// Live output mode - show in terminal
// 		cmd.Stdout = os.Stdout
// 		cmd.Stderr = os.Stderr

// 		err = cmd.Run()
// 		if err != nil {
// 			// return "", fmt.Errorf("❌ Error: remote command failed: %v", err)
// 			// return stdout, err                   // // Return original output  + raw error
// 			return stdout, fmt.Errorf("❌ %v", err) // // Return original output  + raw error
// 		}
// 		return "", nil // No captured output in live mode
// 	} else {
// 		// Silent mode - capture output (original behavior)
// 		var out bytes.Buffer
// 		cmd.Stdout = &out
// 		cmd.Stderr = &out

// 		err = cmd.Run()
// 		stdout = strings.TrimSpace(out.String())

// 		if err != nil {
// 			// return stdout, fmt.Errorf("❌ Error: remote command failed: %v\noutput:\n%s", err, stdout)
// 			return stdout, fmt.Errorf("❌ %v", err)
// 		}
// 		return stdout, nil
// 	}
// }
