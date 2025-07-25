/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/

package util

import (
	"github.com/abtransitionit/luc/pkg/util"
)

// # Purpose
//
// - List SSH reachable OVH VMs
//
// # Usage
//
// logx.L.Infof("list OVH Vm: %s", util.ListOvhVm())
func ListOvhVm() []string {
	// var
	var (
		listVmPrefix       = []string{"o1", "o2", "o3", "o4", "o5"}
		listOs             = []string{"a", "r", "f", "u", "d"}
		listVmSshReachable = []string{}
	)
	// define potentially vm names
	listVmName, _ := util.CartesianProduct(listVmPrefix, listOs)

	// loop over vm names
	for _, vmName := range listVmName {
		// check VM is ssh reachable
		IsReachable, _ := util.IsSshConfiguredVmSshReachable(vmName)
		if IsReachable {
			listVmSshReachable = append(listVmSshReachable, vmName)
		}
	}

	return listVmSshReachable
}
