/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/

package util

import (
	"fmt"
	"os"
	"strings"
)

// # purpose
//
// - add a line to a file if it not exists
func AddLineToFile(filePath string, line string) (string, error) {

	// Check arg
	if filePath == "" {
		return "", fmt.Errorf("no file path provided")
	}
	if line == "" {
		return "", fmt.Errorf("no line provided")
	}

	// Check file exists
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return "", fmt.Errorf("file does not exist: %s", filePath)
	}

	// Check if the line exists
	cli := fmt.Sprintf(`grep -Fxq %q %s`, line, filePath)
	_, err = RunCLILocal(cli)

	// If grep finds the line, do nothing
	if err == nil {
		return "", nil
	}
	// If the file doesn't exist or line not found, continue (grep returns non-zero exit code)
	// If it's a different kind of error, return it
	if !strings.Contains(err.Error(), "command failed") {
		// return fmt.Errorf("error checking line in %s: %w", filePath, err)
		return "", err
	}

	// Append the line to the file
	cli = fmt.Sprintf(`echo %q >> %s`, line, filePath)

	// error
	if _, err = RunCLILocal(cli); err != nil {
		return "", err
	}

	// success
	return "", nil
}

// # Purpose
//
// deletes the left spaces on a multiline string and returns it
func DeleteLeftSpace(s string) string {
	lines := strings.Split(s, "\n")
	for i, line := range lines {
		lines[i] = strings.TrimLeft(line, " \t")
	}
	return strings.Join(lines, "\n")
}

// # Purpose
//
// saves a string to a file
//
// # Logic
//
// - Create an emprty file
// - save the content to the file
// - close the file
//
// # Parameters
//
//   - data: the string to save
//   - path: the path to the file
//   - pathIsRoot: whether the path is a root path
//
// # Returns
//
//   - string: the absolute path to the file
//   - error: any error that occurred
//
// # Example
//
//	absPath, err := SaveStringToFile("Hello, world!", "/tmp/hello.txt")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println("File saved at:", absPath)
func SaveStringToFile(data string, path string, pathIsRoot bool) (string, error) {
	var err error

	if pathIsRoot {
		// Create an empty file as root
		_, err = helperCreateFileAsSudo(path)
		if err != nil {
			return "", fmt.Errorf("%s", err)
		}

		// Save file as root
		_, err = helperSaveStringToFileAsSudo(DeleteLeftSpace(data), path)
		if err != nil {
			return "", fmt.Errorf("%s", err)
		}
	} else {
		// Create file as non root
		file, err := os.Create(path)
		if err != nil {
			return "", fmt.Errorf("failed to create file as non root: %w", err)
		}
		defer file.Close()

		// Save file as non root
		_, err = file.WriteString(DeleteLeftSpace(data))
		if err != nil {
			return "", fmt.Errorf("failed to write to file as non root: %w", err)
		}
	}
	return path, nil
}

func helperCreateFileAsSudo(path string) (string, error) {
	cmd := fmt.Sprintf("sudo touch %q", path)
	out, err := RunCLILocal(cmd)
	if err != nil {
		return "", fmt.Errorf("failed to create file as sudo at %s: %w", path, err)
	}
	return out, nil
}
func helperSaveStringToFileAsSudo(data string, path string) (string, error) {
	// escapedData := strings.ReplaceAll(data, `'`, `'\''`)
	// cmd := fmt.Sprintf("echo '%s' | sudo tee %q > /dev/null", escapedData, path)
	cmd := fmt.Sprintf("echo '%s' | sudo tee %q > /dev/null", data, path)

	out, err := RunCLILocal(cmd)
	if err != nil {
		return "", fmt.Errorf("failed to save string to file as sudo at %s: %w", path, err)
	}
	return out, nil
}
