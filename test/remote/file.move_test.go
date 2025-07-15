/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/

package remote

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/abtransitionit/luc/pkg/util"
	"github.com/stretchr/testify/assert"
)

const vm = "o1u" // example remote VM alias; adjust as needed

// runRemoteCommand runs a command on a remote VM in test mode
func runRemoteCommand(t *testing.T, cmd string) (string, error) {
	out, err := util.RunCLIRemote(vm, cmd)
	if err != nil {
		t.Logf("yoyo %s", err)
	}
	return out, err
}

// Run MvFile function remotely
// cli = fmt.Sprintf("mv %s %s", srcFile, dstFile)
// _, err = runRemoteCommand(t, cli)
// assert.NoError(t, err, "error while moving file remotely")

// _, err = runRemoteCommand(t, cli)
// assert.NoError(t, err, "creating remote source file") // No error should occur while while doing action remotely

func TestMvFile_Nominal(t *testing.T) {
	// Define test inputs
	tmpDir := "/tmp/testmvfile" // remote folder name auto
	srcContent := "test content"
	srcFile := filepath.Join(tmpDir, "source.txt")
	dstFile := filepath.Join(tmpDir, "destination.txt")
	dstFilePermission := "0644"
	FileIsRoot := "false"

	// Create test folder
	cli := fmt.Sprintf("mkdir -p %s", tmpDir)
	_, err := util.RunCLIRemote(vm, cli)
	assert.NoError(t, err, "creating test folder on remote") // No error should occur while while doing action remotely

	// Create source file on remote
	cli = fmt.Sprintf("echo '%s' > %s", srcContent, srcFile)
	_, err = util.RunCLIRemote(vm, cli)
	assert.NoError(t, err, "creating test src file on remote") // No error should occur while doing action remotely

	// run the code under test
	cli = fmt.Sprintf("luc do MoveFile %s %s %s %s", srcFile, dstFile, dstFilePermission, FileIsRoot)
	_, err = util.RunCLIRemote(vm, cli)
	assert.NoError(t, err, "moving file from remote src to remote dst") // No error should occur while doing action

	// Check dst file exists
	cli = fmt.Sprintf("test -f %s", dstFile)
	_, err = util.RunCLIRemote(vm, cli)
	assert.NoError(t, err, "checking dst file exists on remote")

	// Check src file not exists
	cli = fmt.Sprintf("test ! -f %s", srcFile)
	_, err = util.RunCLIRemote(vm, cli)
	assert.NoError(t, err, "checking src file no more exists on remote") // No error should occur while doing action
}

// // Clean after test
// cli = fmt.Sprintf("rm -rf %s", tmpDir)
// _, err = util.RunCLILocal(cli)
// assert.NoError(t, err, "checking test folder is deleted") // No error should occur while doing action

func TestMvFile_SourceFileMissing(t *testing.T) {
	// Define test inputs
	tmpDir := "/tmp/testmvfile"                         // remote folder name
	srcFile := filepath.Join(tmpDir, "nonexistent.txt") // source path
	dstFile := filepath.Join(tmpDir, "destination.txt") // destination path
	dstFilePermission := "0644"
	FileIsRoot := "false"

	// Create test directory
	cli := fmt.Sprintf("mkdir -p %s", tmpDir)
	_, err := util.RunCLIRemote(vm, cli)
	assert.NoError(t, err, "creating test directory on remote") // No error should occur while doing action

	// run the code under test
	cli = fmt.Sprintf("luc do MoveFile %s %s %s %s", srcFile, dstFile, dstFilePermission, FileIsRoot)
	_, err = util.RunCLIRemote(vm, cli)
	assert.Error(t, err, "moving file from remote src to remote dst") // No error should occur while doing action
	// assert.Equal(t, "false", out)

	// clean up
	cli = fmt.Sprintf("rm -rf %s", tmpDir)
	_, err = util.RunCLIRemote(vm, cli)
	assert.NoError(t, err, "cleaning up test directory on remote") // No error should occur while doing action
}

func TestMvFile_SourcePathNotAbsolute(t *testing.T) {
	// Define test inputs
	tmpDir := "/tmp/testmvfile"                         // remote folder name
	srcFile := "relative_source.txt"                    // relative path intentionally
	dstFile := filepath.Join(tmpDir, "destination.txt") // destination path
	dstFilePermission := "0644"
	FileIsRoot := "false"

	// Create test directory
	cli := fmt.Sprintf("mkdir -p %s", tmpDir)
	_, err := util.RunCLIRemote(vm, cli)
	assert.NoError(t, err, "creating test directory on remote") // No error should occur while doing action

	// run the code under test
	cli = fmt.Sprintf("luc do MoveFile %s %s %s %s", srcFile, dstFile, dstFilePermission, FileIsRoot)
	_, err = util.RunCLIRemote(vm, cli)
	assert.Error(t, err, "moving file from remote src to remote dst") // No error should occur while doing action
	// assert.Equal(t, "false", out)

	// clean up
	cli = fmt.Sprintf("rm -rf %s", tmpDir)
	_, err = util.RunCLIRemote(vm, cli)
	assert.NoError(t, err, "cleaning up test directory on remote") // No error should occur while doing action
}

func TestMvFile_DestinationPathNotAbsolute(t *testing.T) {
	// Define test inputs
	tmpDir := "/tmp/testmvfile"                    // remote folder name
	srcFile := filepath.Join(tmpDir, "source.txt") // source path
	dstFile := "relative_destination.txt"          // relative destination path intentionally
	dstFilePermission := "0644"
	FileIsRoot := "false"
	srcContent := "test content"

	// Create test directory
	cli := fmt.Sprintf("mkdir -p %s", tmpDir)
	_, err := util.RunCLIRemote(vm, cli)
	assert.NoError(t, err, "creating test directory on remote") // No error should occur while doing action

	// Create source file on remote
	cli = fmt.Sprintf("echo '%s' > %s", srcContent, srcFile)
	_, err = util.RunCLIRemote(vm, cli)
	assert.NoError(t, err, "creating test src file on remote") // No error should occur while doing action

	// run the code under test
	cli = fmt.Sprintf("luc do MoveFile %s %s %s %s", srcFile, dstFile, dstFilePermission, FileIsRoot)
	_, err = util.RunCLIRemote(vm, cli)
	assert.Error(t, err, "moving file from remote src to remote dst") // No error should occur while doing action
	// assert.Equal(t, "false", out)

	// clean up
	cli = fmt.Sprintf("rm -rf %s", tmpDir)
	_, err = util.RunCLIRemote(vm, cli)
	assert.NoError(t, err, "cleaning up test directory on remote") // No error should occur while doing action
}

func TestMvFile_DestinationDirMissing(t *testing.T) {
	// Define test inputs
	tmpDir := "/tmp/testmvfile"                                            // remote folder name
	srcFile := filepath.Join(tmpDir, "source.txt")                         // source path
	dstFile := filepath.Join(tmpDir, "nonexistent_dir", "destination.txt") // destination path with missing dir
	dstFilePermission := "0644"
	FileIsRoot := "false"
	srcContent := "test content"

	// Create test directory
	cli := fmt.Sprintf("mkdir -p %s", tmpDir)
	_, err := util.RunCLIRemote(vm, cli)
	assert.NoError(t, err, "creating test directory on remote") // No error should occur while doing action

	// Create source file on remote
	cli = fmt.Sprintf("echo '%s' > %s", srcContent, srcFile)
	_, err = util.RunCLIRemote(vm, cli)
	assert.NoError(t, err, "creating test src file on remote") // No error should occur while doing action

	// run the code under test
	cli = fmt.Sprintf("luc do MoveFile %s %s %s %s", srcFile, dstFile, dstFilePermission, FileIsRoot)
	_, err = util.RunCLIRemote(vm, cli)
	assert.Error(t, err, "moving file from remote src to remote dst") // No error should occur while doing action
	// assert.Equal(t, "false", out)

	// clean up
	cli = fmt.Sprintf("rm -rf %s", tmpDir)
	_, err = util.RunCLIRemote(vm, cli)
	assert.NoError(t, err, "cleaning up test directory on remote") // No error should occur while doing action
}

func TestMvFile_SourceNotARegularFile(t *testing.T) {
	// Define test inputs
	tmpDir := "/tmp/testmvfile"                         // remote folder name
	srcDir := filepath.Join(tmpDir, "sourcedir")        // source directory path
	dstFile := filepath.Join(tmpDir, "destination.txt") // destination path
	dstFilePermission := "0644"
	FileIsRoot := "false"

	// Create test directory structure
	cli := fmt.Sprintf("mkdir -p %s", tmpDir)
	_, err := util.RunCLIRemote(vm, cli)
	assert.NoError(t, err, "creating test directory on remote") // No error should occur while doing action

	// Create source directory on remote
	cli = fmt.Sprintf("mkdir -p %s", srcDir)
	_, err = util.RunCLIRemote(vm, cli)
	assert.NoError(t, err, "creating source directory on remote") // No error should occur while doing action

	// run the code under test
	cli = fmt.Sprintf("luc do MoveFile %s %s %s %s", srcDir, dstFile, dstFilePermission, FileIsRoot)
	_, err = util.RunCLIRemote(vm, cli)
	assert.Error(t, err, "moving file from remote src to remote dst") // No error should occur while doing action
	// assert.Equal(t, "false", out)

	// clean up
	cli = fmt.Sprintf("rm -rf %s", tmpDir)
	_, err = util.RunCLIRemote(vm, cli)
	assert.NoError(t, err, "cleaning up test directory on remote") // No error should occur while doing action
}
