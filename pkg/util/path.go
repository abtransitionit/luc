/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/

package util

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// # Purpose
//
// - build a tree PATH
//
// # Parameters
//
// - basePath: the path from wich to build the tree path
//
// # Returns
//
//   - string: a colon-separated list of all sub directories including the base dir
//   - error: any error encountered during directory traversal.
//
// # Example Usage
//
//	pathStr, err := BuildPathFromSubdirs("/usr/local/bin")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println("export PATH=" + pathStr + ":$PATH")
func GetListChuldDirLocal(basePath string) (string, error) {

	// check arg
	if basePath == "" {
		return "", fmt.Errorf("no base path provided")
	}

	// play code
	cli := fmt.Sprintf(`find %s -type d | sort | paste -sd\;`, basePath)
	path, err := RunCLILocal(cli)
	if err != nil {
		return "", err
	}
	return path, nil
}

func GetSubdirRemote(basePath string, vm string) (string, error) {

	// check arg
	if basePath == "" {
		return "", fmt.Errorf("no base path provided")
	}
	// check arg
	if vm == "" {
		return "", fmt.Errorf("no vm provided")
	}

	// play code
	cli := fmt.Sprintf(`find %s -type d | sort | paste -sd\;`, basePath)
	path, err := RunCLIRemote(vm, cli)
	if err != nil {
		return "", err
	}
	return path, nil
}

// # Purpose
//
// - adds a set of paths to the PATH environment variable avoiding duplication
//
// # Parameters
//
// - a semi-colon separated set of paths
//
// # Returns
//
// - the updated PATH string and an error if any.
//
// # Note
//
// - Does not add the path if it not exists

func UpdatePath(srcPath string) (string, error) {
	if srcPath == "" {
		return "", fmt.Errorf("no source path provided")
	}

	// Split source path into multiple entries
	newPaths := strings.Split(srcPath, string(os.PathListSeparator))

	// Get current PATH and convert to map for fast lookup
	currentPath := os.Getenv("PATH")
	existingPaths := strings.Split(currentPath, string(os.PathListSeparator))
	pathMap := make(map[string]bool)
	for _, p := range existingPaths {
		pathMap[p] = true
	}

	// Collect valid new paths
	var additions []string
	for _, np := range newPaths {
		absPath, err := CheckPath(np)
		if err != nil {
			continue // skip invalid paths silently
		}
		if !pathMap[absPath] {
			additions = append(additions, absPath)
			pathMap[absPath] = true
		}
	}

	// If no new paths to add, return current
	if len(additions) == 0 {
		return currentPath, nil
	}

	// Update PATH
	updatedPath := currentPath + string(os.PathListSeparator) + strings.Join(additions, string(os.PathListSeparator))
	if err := os.Setenv("PATH", updatedPath); err != nil {
		return "", fmt.Errorf("failed to update PATH: %v", err)
	}

	return updatedPath, nil
}

func CheckPath(path string) (string, error) {

	// check arg
	if path == "" {
		return "", fmt.Errorf("no path provided")
	}

	// check path is absolute
	absPath, err := filepath.Abs(path)
	if err != nil {
		return "", fmt.Errorf("invalid path: %v", err)
	}

	// check path exists
	if _, err := os.Stat(absPath); os.IsNotExist(err) {
		return "", fmt.Errorf("path does not exist: %s", absPath)
	} else if err != nil {
		return "", fmt.Errorf("error checking path: %v", err)
	}

	// success
	return absPath, nil
}
