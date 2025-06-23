/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/

package util

import (
	"strings"
)

// # purpose
//
// add a line to an rc file
func AddLineRcFile(filePath, line string) error {
	// Check if the line exists
	checkCmd := `grep -Fxq '` + line + `' ` + filePath
	_, err := RunCLILocal(checkCmd)
	if err == nil {
		// fmt.Printf("INFO: Line already exists in %s\n", filePath)
		return nil
	}

	// If grep fails with something other than "not found", return the error
	if !strings.Contains(err.Error(), "command failed") {
		// return fmt.Errorf("error checking line in %s: %w", filePath, err)
		return err
	}

	// Add the line
	appendCmd := `echo '` + line + `' >> ` + filePath
	_, err = RunCLILocal(appendCmd)
	if err != nil {
		// return fmt.Errorf("failed to append line to %s: %w", filePath, err)
		return err
	}

	// fmt.Printf("INFO: Line added to %s\n", filePath)
	return nil
}
