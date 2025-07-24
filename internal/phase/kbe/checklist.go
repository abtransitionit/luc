/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package kbe

import (
	"github.com/abtransitionit/luc/internal/config"
	"github.com/abtransitionit/luc/pkg/util"
	"github.com/abtransitionit/luc/test"
)

func checkList(arg ...string) (string, error) {
	// // method
	// util.CheckSshV1(config.KindVm)

	// method
	for _, vm := range util.GetSlicefromStringWithSpace(config.KbeListNode) {
		test.CheckVmIsSshReachable(vm)
	}
	for _, vm := range util.GetSlicefromStringWithSpace(config.KbeListNode) {
		test.CheckCliExistsOnremote(vm, "gpg")

	}
	for _, vm := range util.GetSlicefromStringWithSpace(config.KbeListNode) {
		test.CheckCliExistsOnremote(vm, "curl")

	}

	// success
	return "", nil
}
