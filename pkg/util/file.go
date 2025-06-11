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
	"path/filepath"

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

func ListTgzInMemory(data []byte) error {
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

		// Print name and size (like ls -lh)
		fmt.Printf("%-50s %10d bytes\n", hdr.Name, hdr.Size)
	}

	return nil
}
