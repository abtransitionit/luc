/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/

package gocli

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/abtransitionit/luc/pkg/util"
)

// # Purpose
//
// - Build and deploy luc for the current platform
//
// # Parameters
//
// - srcProjectFolderPath: the absolute folder path to the local GIT project
// - dstBinaryFolderPath:  the absolute folder path to the output file

func GoBuild(srcProjectFolderPath string, dstBinaryFilePath string) (string, error) {

	// get property
	osType, err := util.GetPropertyLocal("ostype")
	if err != nil {
		return "", err
	}
	osArch, err := util.GetPropertyLocal("osarch")
	if err != nil {
		return "", err
	}

	return GoBuildXPtf(srcProjectFolderPath, dstBinaryFilePath, osType, osArch)

}
func GoBuildXPtf(srcProjectFolderPath string, dstBinaryFilePath string, osType string, osArch string) (string, error) {

	// check source path
	if srcProjectFolderPath == "" {
		return "", fmt.Errorf("❌ Error: source project folder path not provided")
	}
	if !filepath.IsAbs(srcProjectFolderPath) {
		return "", fmt.Errorf("❌ Error: source project folder path must be absolute: %s", srcProjectFolderPath)
	}
	if stat, err := os.Stat(srcProjectFolderPath); err != nil || !stat.IsDir() {
		return "", fmt.Errorf("❌ Error: source project folder does not exist or is not a directory: %s", srcProjectFolderPath)
	}

	// check destination path
	if dstBinaryFilePath == "" {
		return "", fmt.Errorf("❌ Error: destination binary file path not provided")
	}
	if !filepath.IsAbs(dstBinaryFilePath) {
		return "", fmt.Errorf("❌ Error: destination binary file path must be absolute: %s", dstBinaryFilePath)
	}
	dstBinaryFolderPath := filepath.Dir(dstBinaryFilePath)
	if stat, err := os.Stat(dstBinaryFolderPath); err != nil || !stat.IsDir() {
		return "", fmt.Errorf("❌ Error: destination folder does not exist: %s", dstBinaryFolderPath)
	}

	// Clean up old builds (optional: improve later as per TODO)
	cli := fmt.Sprintf(`rm -rf %s &> /dev/null`, dstBinaryFilePath)
	if _, err := util.RunCLILocal(cli); err != nil {
		return "", fmt.Errorf("❌ Cleanup failed: %v", err)
	}

	// build command
	cli = fmt.Sprintf(`
		cd %s && 
		GOOS=%s GOARCH=%s go build -o %s && 
		cd -`,
		srcProjectFolderPath,
		osType, osArch,
		dstBinaryFilePath,
	)
	// error
	if _, err := util.RunCLILocal(cli); err != nil {
		return "", fmt.Errorf("❌ Error: Build failed: %v", err)
	}

	// success
	return dstBinaryFilePath, nil
}

// sudo mv %s /usr/local/bin/luc &&
// GOOS=linux GOARCH=amd64 go build -o %s
