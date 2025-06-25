/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package rupgrade

import (
	"strings"

	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/abtransitionit/luc/pkg/util"
)

// # Purpose
//
// - This stage create an instance of the structure to be pipelined
// - 1 instance of the structure per item in the cliNameList (e.g 9 cli => 9 structures)
// - This stage will send (out chan<-) each instance into the channel
func source(out chan<- PipelineData, vmList string) {
	// close channel when this code ended
	// closing it make it available for next stage
	// because it is defined outside
	defer close(out)
	vms := strings.Fields(vmList) // convert ListAsString to slice

	logx.L.Debugf("defining data instances to be pipelined")
	for _, vm := range vms {
		vm = strings.TrimSpace(vm)
		if vm == "" {
			continue
		}
		// define one per VM
		data := PipelineData{}

		// get some OS property
		osFamily, err := util.GetRemoteProperty("osfamily", vm)
		if err != nil {
			data.Err = err
			logx.L.Debugf("❌ Error detected", err)
		}

		osDistro, err := util.GetRemoteProperty("osdistro", vm)
		if err != nil {
			data.Err = err
			logx.L.Debugf("❌ Error detected")
		}

		hostType, err := util.GetRemoteProperty("host", vm)
		if err != nil {
			data.Err = err
			logx.L.Debugf("❌ Error detected")
		}

		osVersion, err := util.GetRemoteProperty("osversion", vm)
		if err != nil {
			data.Err = err
			logx.L.Debugf("❌ Error detected")
		}

		kernelVersion, err := util.GetRemoteProperty("oskversion", vm)
		if err != nil {
			data.Err = err
			logx.L.Debugf("❌ Error detected")
		}

		// set this instance properties
		data.HostName = vm
		data.OsFamily = osFamily
		data.OsDistro = osDistro
		data.HostType = hostType
		data.OsVersion = osVersion
		data.OskernelVersionBefore = kernelVersion

		// log information
		logx.L.Debugf("[%s] defined data instances to be pipelined", vm)

		out <- data

	} // for

}
