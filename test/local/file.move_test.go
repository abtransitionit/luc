/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/

package local

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/abtransitionit/luc/pkg/util"
	"github.com/stretchr/testify/assert"
)

func TestMvFile_Nominal(t *testing.T) {
	// Create temporary source file and destination path
	tmpDir := t.TempDir()
	srcFilePath := filepath.Join(tmpDir, "source.txt")
	dstFilePath := filepath.Join(tmpDir, "destination.txt")

	// Create the source file
	err := os.WriteFile(srcFilePath, []byte("test content"), 0644)
	assert.NoError(t, err, "should be able to create source file")

	// Run the function under test (non-root mode)
	msg, err := util.MvFile(srcFilePath, dstFilePath, 0644, false)

	// Assertions
	assert.NoError(t, err)        // No error should occur when moving the file
	assert.Equal(t, "true", msg)  // Expected return value is "true"
	_, err = os.Stat(dstFilePath) // Check that destination file exists
	assert.NoError(t, err, "expected destination file to exist")
	_, err = os.Stat(srcFilePath) // Check that source file no longer exists
	assert.Error(t, err, "expected source file to be gone")
}

func TestMvFile_SourceFileMissing(t *testing.T) {
	// Create temporary destination path
	tmpDir := t.TempDir()
	dstFilePath := filepath.Join(tmpDir, "destination.txt")

	// Provide non-existing source file path
	srcFilePath := filepath.Join(tmpDir, "nonexistent.txt")

	// Run the function under test
	msg, err := util.MvFile(srcFilePath, dstFilePath, 0644, false)

	// Assertions
	assert.Error(t, err)          // Should fail because source file does not exist
	assert.Equal(t, "false", msg) // Expected return is "false"
}

func TestMvFile_SourcePathNotAbsolute(t *testing.T) {
	// Prepare source and destination paths (source is relative path)
	tmpDir := t.TempDir()
	srcFilePath := "relative_source.txt"
	dstFilePath := filepath.Join(tmpDir, "destination.txt")

	// Run the function under test
	msg, err := util.MvFile(srcFilePath, dstFilePath, 0644, false)

	// Assertions
	assert.Error(t, err)          // Should fail because source path is not absolute
	assert.Equal(t, "false", msg) // Expected return is "false"
}

func TestMvFile_DestinationPathNotAbsolute(t *testing.T) {
	// Create temporary source file
	tmpDir := t.TempDir()
	srcFilePath := filepath.Join(tmpDir, "source.txt")
	dstFilePath := "relative_destination.txt" // invalid destination

	// Create the source file
	err := os.WriteFile(srcFilePath, []byte("test content"), 0644)
	assert.NoError(t, err, "should be able to create source file")

	// Run the function under test
	msg, err := util.MvFile(srcFilePath, dstFilePath, 0644, false)

	// Assertions
	assert.Error(t, err)          // Should fail because destination path is not absolute
	assert.Equal(t, "false", msg) // Expected return is "false"
}

func TestMvFile_DestinationDirMissing(t *testing.T) {
	// Create temporary source file
	tmpDir := t.TempDir()
	srcFilePath := filepath.Join(tmpDir, "source.txt")
	dstFilePath := filepath.Join(tmpDir, "nonexistent_dir", "destination.txt")

	// Create the source file
	err := os.WriteFile(srcFilePath, []byte("test content"), 0644)
	assert.NoError(t, err, "should be able to create source file")

	// Run the function under test
	msg, err := util.MvFile(srcFilePath, dstFilePath, 0644, false)

	// Assertions
	assert.Error(t, err)          // Should fail because destination directory does not exist
	assert.Equal(t, "false", msg) // Expected return is "false"
}

func TestMvFile_SourceNotARegularFile(t *testing.T) {
	// Create a directory instead of a regular file
	tmpDir := t.TempDir()
	srcDirPath := filepath.Join(tmpDir, "sourcedir")
	err := os.Mkdir(srcDirPath, 0755)
	assert.NoError(t, err, "should be able to create source directory")

	dstFilePath := filepath.Join(tmpDir, "destination.txt")

	// Run the function under test
	msg, err := util.MvFile(srcDirPath, dstFilePath, 0644, false)

	// Assertions
	assert.Error(t, err)          // Should fail because source is not a regular file
	assert.Equal(t, "false", msg) // Expected return is "false"
}
