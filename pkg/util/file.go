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

	"github.com/abtransitionit/luc/pkg/config"
	"github.com/abtransitionit/luc/pkg/errorx"
)

// # Purpose
//
//   - checks if a regular file exists and is accessible.
//
// # Returns
//
//   - (true, nil)  file exists and is accessible
//   - (false, error) file not exists, or permission issue, special file, or other system errors
//
// # Note
//
//   - TODO: should check os.IsNotExist(err)
func CheckFileExists(path string) (string, error) {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return "false", fmt.Errorf("❌ Error: file does not exist: %s", path)
		}
		return "false", fmt.Errorf("❌ Error: file inaccessible: %v", err)
	}
	return "true", nil
}

// # Purpose
//
//   - deletes a file
//
// # Returns
//
//   - (true, nil)  if the file is successfully deleted
//   - (false, error) for permission issues or other system errors
func DeleteFile(path string) (string, error) {
	err := os.Remove(path)
	if err != nil {
		return "", fmt.Errorf("❌ Error: could not delete file %s: %w", path, err)
	}

	return fmt.Sprintf("✅ deleted file: %s", path), nil
}

// # Purpose
//
//   - creates an empty regular file
//
// # Returns
//
//   - (true, nil)  if the file is successfully created
//   - (false, error) for permission issues or other system errors
func TouchFile(path string) (string, error) {
	file, err := os.OpenFile(path, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		return "", fmt.Errorf("❌ Error: could not touch file %s: %w", path, err)
	}
	defer file.Close()

	return fmt.Sprintf("✅ touched file: %s", path), nil
}

// # Purpose
//
//   - checks if a folder exists and is accessible.
//
// # Returns
//
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

// # Purpose
//
// writes existing data in memory to a file at the specified path.
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
func SaveToFile(data []byte, path string) (string, error) {
	// manage argument
	if path == "" {
		// msg := fmt.Sprintf("path is empty (%s)", path)
		// log.Debugf("❌ %s", msg)
		return errorx.StringError("path is empty", "", errors.New(""))
	}
	if data == nil {
		// msg := "memory data to save is nil"
		// log.Debugf("❌ %s", msg)
		return errorx.StringError("Save empty data memory", "", errors.New(""))
	}
	// prerequisit: check it is an absolute path
	absPath, err := filepath.Abs(path)
	if err != nil {
		// msg := fmt.Sprintf("get absolute path (%s)", path)
		// log.Debugf("❌ %s", msg)
		return errorx.StringError("get absolute path", path, err)
	}
	// create a file
	file, err := os.Create(absPath)
	if err != nil {
		// msg := fmt.Sprintf("create file (%s)", absPath)
		// log.Debugf("❌ %s", msg)
		return errorx.StringError("create file", absPath, err)
	}
	defer file.Close()

	// copy content to file
	bytesWritten, err := file.Write(data)
	if err != nil {
		// msg := fmt.Sprintf("write to file (%s)", absPath)
		// log.Debugf("❌ %s", msg)
		return errorx.StringError("write to file", absPath, err)
	}

	// handle applogic SUCCESS
	kbWritten := float64(bytesWritten) / 1024.0
	// log.Infof("✅ file (%s) created succesfully from memory data", absPath)
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

// # Purpose
//
// func IsGzippedMemoryContent(data []byte) (bool, error) {
// 	if len(data) < 2 {
// 		return false, fmt.Errorf("maybe not a gzipped file")
// 	}
// 	if data[0] != 0x1F || data[1] != 0x8B {
// 		return false, fmt.Errorf("Surely a gzipped file")
// 	}
// 	return true, nil
// }

// # Purpose
//
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

// # Purpose
//
// check if data culred in memory is an exe
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

// # Purpose
//
// decompresses a tgz file into a destination folder
func UnTgz(srcTgzPath string, destFolder string) error {
	// check arg
	if srcTgzPath == "" {
		return fmt.Errorf("❌ Error: Source file is empty")
	}
	if destFolder == "" {
		return fmt.Errorf("❌ Error: Destination folder is empty")
	}

	// Check absolute path
	if !strings.HasPrefix(srcTgzPath, "/") {
		return fmt.Errorf("❌ Error: srcTgzPath must be an absolute path starting with '/'")
	}
	if !strings.HasPrefix(destFolder, "/") {
		return fmt.Errorf("❌ Error: destination directory must be an absolute path starting with '/'")
	}

	// Create dest folder if not exist
	if err := os.MkdirAll(destFolder, 0755); err != nil {
		return err
	}

	// check folder exists
	if _, err := os.Stat(destFolder); os.IsNotExist(err) {
		return fmt.Errorf("❌ Error: Destination folder does not exist")
	}

	// play cli
	cli := fmt.Sprintf("tar -C %s -xzf %s", destFolder, srcTgzPath)
	_, err := RunCLILocal(cli)
	if err != nil {
		return fmt.Errorf("❌ Error: %s : %s", err, cli)
	}
	return nil
}

// # Purpose
//
// - Moves a single file from srcPath to dstPath
//
// # Parameters
//   - srcPath: absolute path to the source file.
//   - dstPath: absolute dst folder or absolute filepath (if rename).
//
// # Returns
//   - success: true if the move operation succeeded, false otherwise.
//   - err: a detailed error if any of the validation or move steps fail.
//
// # Example:
//
//	success, err := MvFile2("/tmp/foo.txt", "/var/log/foo-renamed.txt")
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
func MvFile(srcFilePath, dstFilePath string, permission os.FileMode, pathIsRoot bool) (string, error) {

	// check source file path
	if srcFilePath == "" {
		return "false", fmt.Errorf("source path not provided")
	}
	if !filepath.IsAbs(srcFilePath) {
		return "false", fmt.Errorf("source path must be absolute: %s", srcFilePath)
	}
	srcInfo, err := os.Stat(srcFilePath)
	if err != nil {
		return "false", fmt.Errorf("source file does not exist: %w", err)
	}
	if !srcInfo.Mode().IsRegular() {
		return "false", fmt.Errorf("source path is not a regular file: %s", srcFilePath)
	}

	// check destination file path
	if dstFilePath == "" {
		return "false", fmt.Errorf("destination file path not provided")
	}
	if !filepath.IsAbs(dstFilePath) {
		return "false", fmt.Errorf("destination file path must be absolute: %s", dstFilePath)
	}
	dstDirPath := filepath.Dir(dstFilePath)
	dstDirInfo, err := os.Stat(dstDirPath)
	if err != nil {
		return "false", fmt.Errorf("destination directory does not exist: %w", err)
	}
	if !dstDirInfo.IsDir() {
		return "false", fmt.Errorf("destination parent path is not a directory: %s", dstDirPath)
	}

	// Perform actions as root user
	if pathIsRoot {
		// move
		cli := fmt.Sprintf(`sudo mv "%s" "%s"`, srcFilePath, dstFilePath)
		if _, err := RunCLILocal(cli); err != nil {
			return "false", err
		}
		// set permissions
		cli = fmt.Sprintf(`sudo chmod "%#o" "%s"`, permission, dstFilePath)
		if _, err := RunCLILocal(cli); err != nil {
			return "false", err
		}
		// success as root
		return "true", nil
	}

	// Perform actions as non-root user

	// move
	if err := os.Rename(srcFilePath, dstFilePath); err != nil {
		return "false", fmt.Errorf("failed to move file: %w", err)
	}
	// Set permissions
	if err := os.Chmod(dstFilePath, permission); err != nil {
		return "false", fmt.Errorf("failed to set source file permissions: %w", err)
	}
	// success as non-root
	return "true", nil
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

// # Purpose
//
// Helper function to remove a file or directory as sudo user.
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
//   - pathIsRoot: whether to perform the move using elevated privileges.
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
func MvFolder(srcPath, dstPath string, permission os.FileMode, forceOverwrite bool, pathIsRoot bool) (bool, error) {
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
		if pathIsRoot {
			// Perform the rm as sudo
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
	if pathIsRoot {
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

func GetFileType(absPath string) (config.UrlType, error) {
	// check arg
	if absPath == "" {
		return "", fmt.Errorf("❌ Error : no file provided")
	}

	// check file exists
	_, err := CheckFileExists(absPath)
	if err != nil {
		return "", err
	}
	// get file type
	cli := fmt.Sprintf("file --brief %s", absPath)
	output, err := RunCLILocal(cli)
	if err != nil {
		return "", err
	}

	// parse
	fileInfo := strings.ToLower(output)

	// UseCase
	switch {
	case strings.Contains(fileInfo, "executable"):
		return config.UrlExe, nil
	case strings.Contains(fileInfo, "gzip compressed"):
		return config.UrlTgz, nil
	case strings.Contains(fileInfo, "git"):
		return config.UrlGit, nil
	case strings.Contains(fileInfo, "go source"):
		return config.UrlGo, nil
	default:
		// do nothing
	}
	return "", nil
}
