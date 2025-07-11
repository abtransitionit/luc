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

func TestMvFile_Nominal(t *testing.T) {
	// Define test inputs
	tmpDir := "/tmp/testmvfile" // remote folder name
	srcFile := filepath.Join(tmpDir, "source.txt")
	dstFile := filepath.Join(tmpDir, "destination.txt")
	dstFilePermission := "0644"
	FileIsRoot := "false"

	// remove previous test artifacts - TODO use Temporary folders
	cli := fmt.Sprintf("rm -rf %s && mkdir -p %s", tmpDir, tmpDir)
	runRemoteCommand(t, cli)

	// Create source file on remote
	content := "test content"
	cli = fmt.Sprintf("echo '%s' > %s", content, srcFile)
	_, err := runRemoteCommand(t, cli)
	assert.NoError(t, err, "creating remote source file") // No error should occur while creating the file remotely

	// run the code under test
	// MvFile(srcFilePath, dstFilePath, 0644, false)
	cli = fmt.Sprintf("luc do MoveFile %s %s %s %s", srcFile, dstFile, dstFilePermission, FileIsRoot)
	_, err = util.RunCLILocal(cli)

	// Run MvFile function remotely
	// cli = fmt.Sprintf("mv %s %s", srcFile, dstFile)
	// _, err = runRemoteCommand(t, cli)
	assert.NoError(t, err, "error while moving file remotely")

	// Check destination exists
	_, err = runRemoteCommand(t, fmt.Sprintf("test -f %s", dstFile))
	assert.NoError(t, err, "checking remote file existence")

	// Check source no longer exists
	_, err = runRemoteCommand(t, fmt.Sprintf("test ! -f %s", srcFile))
	assert.NoError(t, err, "checking file is removed")

	// Clean after test
	runRemoteCommand(t, fmt.Sprintf("rm -rf %s", tmpDir))
}

func TestMvFile_SourceFileMissing(t *testing.T) {
	// Define test inputs
	tmpDir := "/tmp/testmvfile"                         // a name
	srcFile := filepath.Join(tmpDir, "nonexistent.txt") // a path name
	dstFile := filepath.Join(tmpDir, "destination.txt") // a path name

	// clean remote before test
	cli := fmt.Sprintf("rm -rf %s && mkdir -p %s", tmpDir, tmpDir)
	runRemoteCommand(t, cli)

	// run the code under test
	msg, err := util.MvFile(srcFile, dstFile, 0644, false)
	assert.Error(t, err)
	assert.Equal(t, "false", msg)

	runRemoteCommand(t, fmt.Sprintf("rm -rf %s", tmpDir))
}

func TestMvFile_SourcePathNotAbsolute(t *testing.T) {
	tmpDir := "/tmp/testmvfile"
	srcFile := "relative_source.txt" // relative path intentionally
	dstFile := filepath.Join(tmpDir, "destination.txt")

	runRemoteCommand(t, fmt.Sprintf("rm -rf %s && mkdir -p %s", tmpDir, tmpDir))

	msg, err := util.MvFile(srcFile, dstFile, 0644, false)
	assert.Error(t, err)
	assert.Equal(t, "false", msg)

	runRemoteCommand(t, fmt.Sprintf("rm -rf %s", tmpDir))
}

func TestMvFile_DestinationPathNotAbsolute(t *testing.T) {
	tmpDir := "/tmp/testmvfile"
	srcFile := filepath.Join(tmpDir, "source.txt")
	dstFile := "relative_destination.txt" // relative dest path

	runRemoteCommand(t, fmt.Sprintf("rm -rf %s && mkdir -p %s", tmpDir, tmpDir))

	_, err := runRemoteCommand(t, fmt.Sprintf("echo 'test content' > %s", srcFile))
	assert.NoError(t, err, "create remote source file")

	msg, err := util.MvFile(srcFile, dstFile, 0644, false)
	assert.Error(t, err)
	assert.Equal(t, "false", msg)

	runRemoteCommand(t, fmt.Sprintf("rm -rf %s", tmpDir))
}

func TestMvFile_DestinationDirMissing(t *testing.T) {
	tmpDir := "/tmp/testmvfile"
	srcFile := filepath.Join(tmpDir, "source.txt")
	dstFile := filepath.Join(tmpDir, "nonexistent_dir", "destination.txt")

	runRemoteCommand(t, fmt.Sprintf("rm -rf %s && mkdir -p %s", tmpDir, tmpDir))

	_, err := runRemoteCommand(t, fmt.Sprintf("echo 'test content' > %s", srcFile))
	assert.NoError(t, err, "create remote source file")

	msg, err := util.MvFile(srcFile, dstFile, 0644, false)
	assert.Error(t, err)
	assert.Equal(t, "false", msg)

	runRemoteCommand(t, fmt.Sprintf("rm -rf %s", tmpDir))
}

func TestMvFile_SourceNotARegularFile(t *testing.T) {
	tmpDir := "/tmp/testmvfile"
	srcDir := filepath.Join(tmpDir, "sourcedir")
	dstFile := filepath.Join(tmpDir, "destination.txt")

	runRemoteCommand(t, fmt.Sprintf("rm -rf %s && mkdir -p %s", tmpDir, tmpDir))
	runRemoteCommand(t, fmt.Sprintf("mkdir -p %s", srcDir))

	msg, err := util.MvFile(srcDir, dstFile, 0644, false)
	assert.Error(t, err)
	assert.Equal(t, "false", msg)

	runRemoteCommand(t, fmt.Sprintf("rm -rf %s", tmpDir))
}
