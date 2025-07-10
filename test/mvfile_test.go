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

func TestMvFile(t *testing.T) {
	// create temporary folder
	tmpDir := t.TempDir() // creates a temp directory that will be cleaned up automatically
	// create temporary file inside that folder
	tmpFile, err := os.CreateTemp(tmpDir, "mytempfile-*.txt") // creates a temp file inside that dir
	if err != nil {
		t.Fatal(err)
	}
	defer tmpFile.Close()
	// log
	t.Logf("tmpDir: %s", tmpDir)
	t.Logf("tmpDir: %s", tmpFile.Name())

	// Define function parameters
	srcDir := t.TempDir() // create a uniq folder and clean it up after the test
	dstDir := t.TempDir()

	srcFile := filepath.Join(srcDir, "test.txt")
	dstFile := filepath.Join(dstDir, "test.txt")

	content := []byte("hello world")
	err = os.WriteFile(srcFile, content, 0644) // create the source file
	assert.NoError(t, err)                     // check no error happened while creating the file.

	// Action: test the function
	result, err := util.MvFile(srcFile, dstFile, 0644, false)
	assert.NoError(t, err)             // check no error happened while moving the file.
	assert.Contains(t, result, "true") // check Real (aka. result) Vs Expected (aka. "true")

	// Check: destination file exists
	data, err := os.ReadFile(dstFile)
	assert.NoError(t, err)         // check no error happened while reading the file
	assert.Equal(t, content, data) // check Expected (aka. content) Vs Real (aka. data)

	// Check: source file no longer exists
	_, err = os.Stat(srcFile)
	assert.True(t, os.IsNotExist(err)) // check the source file is gone
}
