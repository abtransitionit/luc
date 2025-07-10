/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/

package test

import (
	"fmt"
	"testing"

	"github.com/abtransitionit/luc/pkg/util"
	"github.com/stretchr/testify/assert"
)

func TestRemoteTouchFile_Nominal(t *testing.T) {
	// Define var
	vm := "o1u"
	vm = "o2a"
	remoteFilePath := "/tmp/testfile_remote.txt"

	// remote run the function under test
	cli := fmt.Sprintf("luc do TouchFile %s --remote %s", remoteFilePath, vm)
	t.Logf("⚠️ cli: %s", cli)
	msg, err := util.RunCLIRemote(vm, cli)

	// Assertions
	assert.NoError(t, err)                                  // No error should occur while touching the file remotely
	assert.Equal(t, "✅ touched file: "+remoteFilePath, msg) // check expected return vs actual return

	// check file exists on the remote
	cli = fmt.Sprintf("test -f %s && echo true || echo false", remoteFilePath)
	existsMsg, err := util.RunCLIRemote(vm, cli)

	assert.NoError(t, err)             // No error should occur during file existence check
	assert.Equal(t, "true", existsMsg) // check expected return vs actual return (The remote file should exist)
}
