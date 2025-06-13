/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/

package util

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/abtransitionit/luc/pkg/errorx"
	"go.uber.org/zap"
)

// FolderExists checks if a folder exists and is accessible.
// Returns:
//   - (true, nil)  if the folder exists
//   - (false, nil) if the folder doesn't exist (normal case)
//   - (false, error) for permission issues or other system errors
func FolderExists(path string) (bool, error) {
	info, err := os.Stat(path)
	if err == nil {
		return info.IsDir(), nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	// handle system FAILURE
	return errorx.BoolError("check folder exists", path, err)
}

// FileExists checks if a file exists and is accessible.
// Returns:
//   - (true, nil)  if the file exists
//   - (false, nil) if the file doesn't exist (normal case)
//   - (false, error) for permission issues or other system errors
func FileExists(path string) (bool, error) {
	info, err := os.Stat(path)
	if err == nil {
		return !info.IsDir(), nil // return  true nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	// handle system FAILURE
	return errorx.BoolError("check file exists", path, err)
}

// writes data in memory to a file at the specified path.
//
// Parameters:
//   - path : string - Filesystem path where data should be written
//   - data : []byte - Byte slice containing data to write
//
// Returns:
//   - string : Success information (file path + size) ("" on failure)
//   - error  : Failure details (nil on success)
//
// Possible returns:
//   - ("/path/to/file (128 bytes)", nil) : On success
//   - ("", error)												: On failure
//     -- fs.ErrPermission  : On permission error
//     -- os.ErrNotExist    : If parent dir missing
//     -- syscall.ENOSPC    : If disk full
//     -- Filesystem errors
//     -- Write failures
//
// Returns:
//   - error - nil on success, or an error describing the failure:
//
// Example:
//
//	_, err = util.SaveToFile(logx.L, "/tmp/toto", fileInMemory)
//	if err != nil {
//		return
//	}
//
// Notes:
//   - Existing files will be overwritten
//   - The complete file is written atomically (all-or-nothing)
//   - Parent directories must exist (does not create directories)
//   - Uses 0644 file permissions by default
func SaveToFile(log *zap.SugaredLogger, path string, data []byte) (string, error) {
	// manage argument
	if path == "" {
		msg := fmt.Sprintf("path is empty (%s)", path)
		log.Debugf("❌ %s", msg)
		return errorx.StringError("path is empty", "", errors.New(""))
	}
	if data == nil {
		msg := "memory data to save is nil"
		log.Debugf("❌ %s", msg)
		return errorx.StringError("Save empty data memory", "", errors.New(""))
	}
	// prerequisit: check it is an absolute path
	absPath, err := filepath.Abs(path)
	if err != nil {
		msg := fmt.Sprintf("get absolute path (%s)", path)
		log.Debugf("❌ %s", msg)
		return errorx.StringError("get absolute path", path, err)
	}
	// create a file
	file, err := os.Create(absPath)
	if err != nil {
		msg := fmt.Sprintf("create file (%s)", absPath)
		log.Debugf("❌ %s", msg)
		return errorx.StringError("create file", absPath, err)
	}
	defer file.Close()

	// copy content to file
	bytesWritten, err := file.Write(data)
	if err != nil {
		msg := fmt.Sprintf("write to file (%s)", absPath)
		log.Debugf("❌ %s", msg)
		return errorx.StringError("write to file", absPath, err)
	}

	// handle applogic SUCCESS
	kbWritten := float64(bytesWritten) / 1024.0
	log.Infof("✅ file (%s) created succesfully from memory data", absPath)
	return fmt.Sprintf("%s (%.0f KB)", absPath, kbWritten), nil
}

// # Purpose
//
//   - Lists the content of tgz file that is in memory.
//   - Size is in kB
//
// # Example Usage
//
// _ := util.ListTgzContentInMemory(fileInMemory)
//
// # Usage
//
//	err = util.ListTgzContentInMemory(fileInMemory)
//	if err != nil {
//		return "", err
//	}
func ListTgzContentInMemory(data []byte) error {
	gzReader, err := gzip.NewReader(bytes.NewReader(data))
	if err != nil {
		return err
	}
	defer gzReader.Close()

	tarReader := tar.NewReader(gzReader)

	for {
		hdr, err := tarReader.Next()
		if err == io.EOF {
			break // done
		}
		if err != nil {
			return err
		}

		fmt.Printf("%-50s %10d kB\n", hdr.Name, (hdr.Size+1023)/1024)
	}

	return nil
}

// func IsGzippedMemoryContent(data []byte) (bool, error) {
// 	if len(data) < 2 {
// 		return false, fmt.Errorf("maybe not a gzipped file")
// 	}
// 	if data[0] != 0x1F || data[1] != 0x8B {
// 		return false, fmt.Errorf("Surely a gzipped file")
// 	}
// 	return true, nil
// }

// check if data culred in memory is a valid gzip tar
func IsGzippedMemoryContent(data []byte) (bool, error) {
	if len(data) < 2 {
		return false, fmt.Errorf("data too short to determine gzip header")
	}
	if data[0] == 0x1F && data[1] == 0x8B {
		return true, nil
	}
	return false, nil // not gzipped, but not an error
}

func IsMemoryContentAnExe(data []byte) (bool, error) {
	if len(data) < 4 {
		return false, fmt.Errorf("data too short to determine executable")
	}

	// Check ELF (Linux)
	if data[0] == 0x7F && data[1] == 'E' && data[2] == 'L' && data[3] == 'F' {
		return true, nil
	}

	// Check Windows PE (MZ)
	if data[0] == 'M' && data[1] == 'Z' {
		return true, nil
	}

	// Check Mach-O (multiple signatures)
	machO1 := []byte{0xCF, 0xFA, 0xED, 0xFE}
	machO2 := []byte{0xFE, 0xED, 0xFA, 0xCF}
	machO3 := []byte{0xCA, 0xFE, 0xBA, 0xBE}

	if bytes.HasPrefix(data, machO1) || bytes.HasPrefix(data, machO2) || bytes.HasPrefix(data, machO3) {
		return true, nil
	}

	return false, fmt.Errorf("unknown executable format")
}

func UntargzFile(srcTgzPath string, destFolder string) error {
	// Check destination is an absolute path
	if !strings.HasPrefix(destFolder, "/") {
		return errors.New("destination directory must be an absolute path starting with '/'")
	}
	// Check srcTgzPath is an absolute path
	if !strings.HasPrefix(srcTgzPath, "/") {
		return errors.New("srcTgzPath must be an absolute path starting with '/'")
	}

	// Check destination directory exists
	if err := os.MkdirAll(destFolder, 0755); err != nil {
		return err
	}

	// Prepare the command: tar -C destFolder -xzf tgzPath
	cmd := exec.Command("tar", "-C", destFolder, "-xzf", srcTgzPath)

	// Run the command and wait for it to finish
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

// # Purpose
//
// - Moves a single file from srcPath to dstPath
//
// # Parameters
//   - srcPath: absolute path to the source file.
//   - dstPath: absolute path to the destination file location (can include a rename).
//
// # Returns
//   - success: true if the move operation succeeded, false otherwise.
//   - err: a detailed error if any of the validation or move steps fail.
//
// # Example:
//
//	success, err := MvFile("/tmp/foo.txt", "/var/log/foo-renamed.txt")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	if success {
//	    fmt.Println("File moved successfully.")
//	}
//
// # Notes
//
// The helper function helperMvSudo is used when sudo is required
func MvFile(srcPath, dstPath string, permission os.FileMode, isSudo bool) (bool, error) {
	// check srcPath is absolute
	if !filepath.IsAbs(srcPath) {
		return false, errors.New("source path must be absolute")
	}

	// check dstPath is absolute
	if !filepath.IsAbs(dstPath) {
		return false, errors.New("destination path must be absolute")
	}

	// check source file exists and is a regular file
	srcInfo, err := os.Stat(srcPath)
	if err != nil {
		return false, fmt.Errorf("source file error: %w", err)
	}
	if !srcInfo.Mode().IsRegular() {
		return false, errors.New("source is not a regular file")
	}

	// Set permissions on source file before moving
	if err := os.Chmod(srcPath, permission); err != nil {
		return false, fmt.Errorf("failed to set source file permissions: %w", err)
	}

	// check the parent directory of dstPath exists and is a directory
	dstDir := filepath.Dir(dstPath)
	dstInfo, err := os.Stat(dstDir)
	if err != nil {
		return false, fmt.Errorf("destination directory does not exist: %w", err)
	}
	if !dstInfo.IsDir() {
		return false, fmt.Errorf("destination parent path is not a directory: %s", dstDir)
	}

	// Perform the move as sudo
	if isSudo {
		if err := helperMvSudo(srcPath, dstPath); err != nil {
			return false, err
		}
		return true, nil
	}
	// Perform the move as normal user
	if err := os.Rename(srcPath, dstPath); err != nil {
		return false, fmt.Errorf("failed to move file: %w", err)
	}

	return true, nil
}

// Helper function to move a file as sudo user from srcPath to dstPath
// Returns an error if the command fails or produces output.
//
// Note:
//   - Requires sudo permissions configured appropriately.
//   - May prompt for password unless passwordless sudo is set up.
func helperMvSudo(srcPath, dstPath string) error {
	cmd := exec.Command("sudo", "mv", srcPath, dstPath)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("sudo mv failed: %v, output: %s", err, string(output))
	}
	return nil
}

// Helper function to remove a file or directory as sudo user at targetPath.
// Returns an error if the command fails or produces output.
//
// Note:
//   - Requires sudo permissions configured appropriately.
//   - May prompt for password unless passwordless sudo is set up.
func helperRmSudo(targetPath string) error {
	cmd := exec.Command("sudo", "rm", "-rf", targetPath)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("sudo rm -rf failed: %v, output: %s", err, string(output))
	}
	return nil
}

// # Purpose
//
// Moves a single folder from srcPath to dstPath.
//
// # Parameters
//   - srcPath: absolute path to the source directory.
//   - dstPath: absolute path to the destination directory location (can include a rename).
//   - permission: permission bits to apply to the source directory before the move.
//   - forceOverwrite: whether to overwrite the destination if it already exists.
//   - isSudo: whether to perform the move using elevated privileges.
//
// # Returns
//   - success: true if the move operation succeeded, false otherwise.
//   - err: a detailed error if any of the validation or move steps fail.
//
// # Example:
//
//	success, err := MvFolder("/tmp/mydata", "/opt/data/archive", 0755, false)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	if success {
//	    fmt.Println("Folder moved successfully.")
//	}
//
// # Notes
//
// The helper function helperMvSudo is used when sudo is required
func MvFolder(srcPath, dstPath string, permission os.FileMode, forceOverwrite bool, isSudo bool) (bool, error) {
	// check srcPath is absolute
	if !filepath.IsAbs(srcPath) {
		return false, errors.New("source path must be absolute")
	}

	// check dstPath is absolute
	if !filepath.IsAbs(dstPath) {
		return false, errors.New("destination path must be absolute")
	}

	// check source folder exists and is a directory
	srcInfo, err := os.Stat(srcPath)
	if err != nil {
		return false, fmt.Errorf("source folder error: %w", err)
	}
	if !srcInfo.IsDir() {
		return false, errors.New("source is not a directory")
	}

	// Set permissions on source folder before moving
	if err := os.Chmod(srcPath, permission); err != nil {
		return false, fmt.Errorf("failed to set source folder permissions: %w", err)
	}

	// check the parent directory of dstPath exists and is a directory
	dstDir := filepath.Dir(dstPath)
	dstInfo, err := os.Stat(dstDir)
	if err != nil {
		return false, fmt.Errorf("destination directory does not exist: %w", err)
	}
	if !dstInfo.IsDir() {
		return false, fmt.Errorf("destination parent path is not a directory: %s", dstDir)
	}

	// Handle overwrite if destination exists
	if _, err := os.Stat(dstPath); err == nil {
		if !forceOverwrite {
			return false, fmt.Errorf("destination path already exists: %s", dstPath)
		}
		if isSudo {
			if err := helperRmSudo(dstPath); err != nil {
				return false, fmt.Errorf("failed to remove existing destination with sudo: %w", err)
			}
		} else {
			if err := os.RemoveAll(dstPath); err != nil {
				return false, fmt.Errorf("failed to remove existing destination: %w", err)
			}
		}
	}
	// Perform the move as sudo
	if isSudo {
		if err := helperMvSudo(srcPath, dstPath); err != nil {
			return false, err
		}
		return true, nil
	}

	// Perform the move as normal user
	if err := os.Rename(srcPath, dstPath); err != nil {
		return false, fmt.Errorf("failed to move folder: %w", err)
	}

	return true, nil
}

// returns a string to add to $PATH
//
// # Parameters
//
//   - base: the root directory to start the recursive search.
//
// # Returns
//
//   - string: a colon-separated list of all directories including the base.
//   - error: any error encountered during directory traversal.
//
// # Example
//
//	pathStr, err := BuildPathFromSubdirs("/usr/local/bin")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println("export PATH=" + pathStr + ":$PATH")
func BuildPath(base string) (string, error) {
	var paths []string

	err := filepath.WalkDir(base, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			paths = append(paths, path)
		}
		return nil
	})
	if err != nil {
		return "", err
	}

	return strings.Join(paths, ":"), nil
}
