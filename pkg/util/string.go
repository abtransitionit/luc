/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/

package util

import (
	"encoding/base64"
	"fmt"
	"os"
	"strings"
)

// # purpose
//
// - add a line, at the end of an existing file ONLY if it not ALREADY exists
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

	// Check if the line exists - 0: line found / 1: line not found / >1: other error
	cli := fmt.Sprintf(`grep -Fxq %q %s`, line, filePath)
	_, err = RunCLILocal(cli)

	// If status code is 0 : grep finds the line, do nothing
	if err == nil {
		return fmt.Sprintf("✅ done nothing, line already exists in file: %s", filePath), nil
	}

	if strings.Contains(err.Error(), "command failed") {
		return "", fmt.Errorf("error checking line with grep in %s: %w", filePath, err)
	}

	// now error code is 1 (grep not found the line): Append the line to the file
	cli = fmt.Sprintf(`echo %q >> %s`, line, filePath)
	outp, err := RunCLILocal(cli)
	if err != nil {
		return "", fmt.Errorf("%v, %s", err, outp)
	}

	// success
	return fmt.Sprintf("✅ added line to file : %s", filePath), nil
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
	if pathIsRoot {
		// Create an empty file as root
		_, err := helperCreateFileAsSudo(path)
		if err != nil {
			return "", fmt.Errorf("%s", err)
		}

		// Save file as root
		_, err = helperSaveStringToFileAsSudo(DeleteLeftSpace(data), path)
		if err != nil {
			return "", fmt.Errorf("%s", err)
		}

		return path, nil
	}

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

	return path, nil
}

// # Purpose
//
// gets the content of a Txt file
//
// # Parameters
//
//   - path: the path to the file
//
// # Returns
//
//   - string: a base64 encoded string
//   - error:  any error that occurred
//
// # Example
//
//	strin64, err := GetStringFromFile("/tmp/hello")
//	if err != nil {
//	    log.Fatal(err)
//	}
func GetStringFromFile(path string, pathIsRoot bool) (string, error) {
	var content []byte
	var err error

	// read file as root
	if pathIsRoot {
		content, err = helperReadFileAsSudo(path)
		if err != nil {
			return "", fmt.Errorf("failed to read file as root: %w", err)
		}
		return base64.StdEncoding.EncodeToString(content), nil
	}

	// read file as non-root
	content, err = os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("failed to read file: %w", err)
	}

	return base64.StdEncoding.EncodeToString(content), nil
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

func helperReadFileAsSudo(path string) ([]byte, error) {
	cmd := fmt.Sprintf("sudo cat %q", path)

	out, err := RunCLILocal(cmd)
	if err != nil {
		return nil, fmt.Errorf("failed to read file as sudo at %s: %w", path, err)
	}
	return []byte(out), nil
}
