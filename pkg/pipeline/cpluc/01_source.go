/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package cpluc

import (
	"strings"

	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/abtransitionit/luc/pkg/util"
)

// # Purpose
//
// - This stage creates an instance per item in the vmList
// - e.g 3 items in the vmList => 3 instances of the structure PipelineData
// - It sends (out chan<-) each instance into the output channel
func source(out chan<- PipelineData, dtpip PipelineData, vmList string) {
	// close channel when this code ended
	// closing it make it available for next stage, because it is defined outside
	defer close(out)
	vms := strings.Fields(vmList) // convert ListAsString to slice

	logx.L.Debugf("defining instances to be pipelined")
	logx.L.Debugf("Vms        to provision.    : %d : %s", len(vms), vms)

	for _, vm := range vms {

		vm = strings.TrimSpace(vm)

		if vm == "" {
			continue
		}

		// avoid creating instance for non SSH reachable VMs
		result, err := util.GetPropertyLocal("sshreachability", vm)
		if err != nil {
			logx.L.Debugf("⚠️ %v : %s : %s", err, result, "skipping data instance for it")
			continue
		} else if strings.ToLower(strings.TrimSpace(result)) == "false" {
			logx.L.Debugf("⚠️ [%s] remote vm is not reachable, skipping data instance for it", vm)
			continue
		}

		// define one per item
		data := PipelineData{}

		// set this instance properties
		data.VmName = vm
		data.localOutput = dtpip.localOutput
		data.localExePath = dtpip.localExePath
		data.localOutXptf = dtpip.localOutXptf
		data.remoteTmpPath = "/var/tmp/luc"
		data.remoteExePath = "/usr/local/bin/luc"

		// log information
		logx.L.Debugf("[%s] sending instance to the pipeline", vm)
		// sen this instance to the channel
		out <- data

	} // for
}
