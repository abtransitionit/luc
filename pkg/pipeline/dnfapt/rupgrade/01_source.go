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
func source(out chan<- PipelineData, vms []string) {
	// close channel when this code ended
	// closing it make it available for next stage, because it is defined outside
	defer close(out)

	logx.L.Debugf("defining instances to be pipelined : %d VM(s) : %s", len(vms), vms)
	for _, vm := range vms {
		vm = strings.TrimSpace(vm)
		if vm == "" {
			continue
		}
		// define one per item
		data := PipelineData{}

		// get some OS property
		osFamily, err := util.GetRemoteProperty("osfamily", vm)
		if err != nil {
			data.Err = err
			logx.L.Debugf("[%s] ❌ Error detected 1", vm)
		}

		osDistro, err := util.GetRemoteProperty("osdistro", vm)
		if err != nil {
			data.Err = err
			logx.L.Debugf("[%s] ❌ Error detected 2", vm)
		}

		hostType, err := util.GetRemoteProperty("host", vm)
		if err != nil {
			data.Err = err
			logx.L.Debugf("[%s] ❌ Error detected 3", vm)
		}

		osVersion, err := util.GetRemoteProperty("osversion", vm)
		if err != nil {
			data.Err = err
			logx.L.Debugf("[%s] ❌ Error detected 4", vm)
		}

		kernelVersion, err := util.GetRemoteProperty("oskversion", vm)
		if err != nil {
			data.Err = err
			logx.L.Debugf("[%s] ❌ Error detected 5", vm)
		}

		// set this instance properties
		data.HostName = vm
		data.OsFamily = osFamily
		data.OsDistro = osDistro
		data.HostType = hostType
		data.OsVersion = osVersion
		data.OskernelVersionBefore = kernelVersion

		// log information
		logx.L.Debugf("[%s] send instance to the pipeline", vm)
		// sen this instance to the channel
		out <- data

	} // for
}
