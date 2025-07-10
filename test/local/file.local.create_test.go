/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/

package test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/abtransitionit/luc/pkg/util"
	"github.com/stretchr/testify/assert"
)

func TestTouchFile_Nominal(t *testing.T) {
	// create inputs for the test
	tmpDir := t.TempDir()
	testFilePath := filepath.Join(tmpDir, "testfile.txt")

	// run the function under test
	msg, err := util.TouchFile(testFilePath)

	// Assertions
	assert.NoError(t, err)                                // No error should occur while touching the file
	assert.Equal(t, "✅ touched file: "+testFilePath, msg) // compare expected return vs actual return

	// Check if file exists
	_, err = os.Stat(testFilePath)
	assert.NoError(t, err, "expected file to exist") // check no error happened
}

func TestTouchFile_Error(t *testing.T) {
	// use an intentionally invalid path (to simulate folder not exist or permission issues)
	invalidPath := "/invalid_dir/testfile.txt"

	// run the function under test
	msg, err := util.TouchFile(invalidPath)

	// Assertions
	assert.Error(t, err) // an error should occur (touching an inexistent file)
	assert.Empty(t, msg) // compare expected return vs actual return (The returned message should be empty)
}

func TestDeleteFile_Nominal(t *testing.T) {
	// create inputs for the test
	tmpDir := t.TempDir()
	testFilePath := filepath.Join(tmpDir, "testfile.txt")

	// Create the file to ensure it exists before deletion
	err := os.WriteFile(testFilePath, []byte("test content"), 0644)
	assert.NoError(t, err) // No error should occur while creating the file

	// run the function under test
	msg, err := util.DeleteFile(testFilePath)

	// Assertions
	assert.NoError(t, err)                                // No error should occur while deleting the file
	assert.Equal(t, "✅ deleted file: "+testFilePath, msg) // compare expected return vs actual return

	// Check if file has been deleted
	_, err = os.Stat(testFilePath)
	assert.Error(t, err, "expected file to be deleted") // File should no longer exist
}

func TestDeleteFile_Error(t *testing.T) {
	// use an intentionally invalid path (to simulate folder not exist or permission issues)
	invalidPath := "/invalid_dir/nonexistentfile.txt"

	// run the function under test
	msg, err := util.DeleteFile(invalidPath)

	// Assertions
	assert.Error(t, err) // an error should occur (deleting an inexistent file)
	assert.Empty(t, msg) // compare expected return vs actual return (The returned message should be empty)
}

func TestCheckFileExists_Nominal(t *testing.T) {
	// create inputs for the test
	tmpDir := t.TempDir()
	testFilePath := filepath.Join(tmpDir, "testfile.txt")
	err := os.WriteFile(testFilePath, []byte("hello"), 0644)     // create the file
	assert.NoError(t, err, "should be able to create temp file") // No error should occur while creating the file

	// Run the function under test
	result, err := util.CheckFileExists(testFilePath)

	// Assertions
	assert.NoError(t, err)          // No error should occur for an existing file
	assert.Equal(t, "true", result) // compare expected return vs actual return
}

func TestCheckFileExists_FileNotExist(t *testing.T) {
	// Use a non-existing file path
	nonExistentPath := "/nonexistent/file.txt"

	// Run the function under test
	result, err := util.CheckFileExists(nonExistentPath)

	// Assertions
	assert.Error(t, err)             // an error should occur for missing file
	assert.Equal(t, "false", result) // compare expected return vs actual return
}

func TestCheckFileExists_EmptyPath(t *testing.T) {
	// Run the function under test with empty path
	result, err := util.CheckFileExists("")

	// Assertions
	assert.Error(t, err)             // an error should occur for invalid path
	assert.Equal(t, "false", result) // compare expected return vs actual return
}

// func TestHostnameResolution(t *testing.T) {
// 	cmd := exec.Command("sh", "-c", "ssh -G o1u")
// 	out, err := cmd.CombinedOutput()
// 	if err != nil {
// 		t.Fatalf("ssh -G o1u failed: %v\nOutput: %s", err, out)
// 	}
// 	t.Logf("ssh -G o1u output:\n%s", out)

// 	addrs, err := net.LookupHost("o1u")
// 	t.Logf("DNS lookup results: %v, error: %v", addrs, err)
// }
