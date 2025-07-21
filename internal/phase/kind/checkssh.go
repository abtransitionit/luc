/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package kind

import (
	"github.com/abtransitionit/luc/internal/config"
	"github.com/abtransitionit/luc/pkg/util"
	"github.com/abtransitionit/luc/test"
)

func checkSsh(arg ...string) (string, error) {
	// // method
	// util.CheckSshV1(config.KindVm)

	// method
	for _, vm := range util.GetSlicefromStringWithSpace(config.KbeListNode) {
		test.CheckVmIsSshReachable(vm)
	}

	// success
	return "", nil
}
