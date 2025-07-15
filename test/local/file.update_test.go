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

func TestAddLineToFile_Nominal(t *testing.T) {
	// create inputs for the test
	tmpDir := t.TempDir()
	testFilePath := filepath.Join(tmpDir, "testfile.txt")
	initialContent := "first line"

	// Create the file with initial content
	err := os.WriteFile(testFilePath, []byte(initialContent+"\n"), 0644)
	assert.NoError(t, err, "should be able to create temp file") // No error should occur while creating the file

	// run the function under test with a new line
	out, err := util.AddLineToFile(testFilePath, "new line")

	// Assertions
	assert.NoError(t, err)                                       // No error should occur while adding the line
	assert.Equal(t, "✅ added line to file : "+testFilePath, out) // compare expected return vs actual return

	// Check if line has been added
	content, err := os.ReadFile(testFilePath)
	assert.NoError(t, err, "should be able to read file after adding line") // No error should occur while reading the file
	assert.Contains(t, string(content), "new line")                         // File content should contain the new line
}

func TestAddLineToFile_LineAlreadyExists(t *testing.T) {
	// create inputs for the test
	tmpDir := t.TempDir()
	testFilePath := filepath.Join(tmpDir, "testfile.txt")
	lineToAdd := "existing line"

	// Create the file with the line already present
	err := os.WriteFile(testFilePath, []byte(lineToAdd+"\n"), 0644)
	assert.NoError(t, err, "should be able to create temp file") // No error should occur while creating the file

	// run the function under test with the same line
	out, err := util.AddLineToFile(testFilePath, lineToAdd)

	// Assertions
	assert.NoError(t, err)                                                             // No error should occur if line already exists
	assert.Equal(t, "✅ done nothing, line already exists in file: "+testFilePath, out) // compare expected return vs actual return
}

func TestAddLineToFile_FileNotExist(t *testing.T) {
	// use a non-existing file path
	nonExistentFile := "/nonexistent/path.txt"

	// run the function under test
	out, err := util.AddLineToFile(nonExistentFile, "some line")

	// Assertions
	assert.Error(t, err) // an error should occur if file does not exist
	assert.Empty(t, out) // compare expected return vs actual return (The returned message should be empty)
}

func TestAddLineToFile_EmptyFilePath(t *testing.T) {
	// run the function under test with empty path
	out, err := util.AddLineToFile("", "some line")

	// Assertions
	assert.Error(t, err) // an error should occur for empty file path
	assert.Empty(t, out) // compare expected return vs actual return (The returned message should be empty)
}

func TestAddLineToFile_EmptyLine(t *testing.T) {
	// create inputs for the test
	tmpDir := t.TempDir()
	testFilePath := filepath.Join(tmpDir, "testfile.txt")

	// Create an empty file
	err := os.WriteFile(testFilePath, []byte(""), 0644)
	assert.NoError(t, err, "should be able to create temp file") // No error should occur while creating the file

	// run the function under test with empty line
	out, err := util.AddLineToFile(testFilePath, "")

	// Assertions
	assert.Error(t, err) // an error should occur for empty line
	assert.Empty(t, out) // compare expected return vs actual return (The returned message should be empty)
}
