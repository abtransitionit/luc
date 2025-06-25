/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package rfilecopy

import (
	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/abtransitionit/luc/pkg/util"
)

// # Purpose
//
// - This step checks that the VM is SSH-configured and reachable.
// - If not, it sets `data.Err` and passes it along.
func checkVM(in <-chan PipelineData, out chan<- PipelineData) {
	defer close(out)

	// loop over each key of the instance structure PipelineData
	for data := range in {
		// if this instance property exists
		if data.Err != nil {
			// send the instance instance into the channel (for next stage/step)
			out <- data
			// log information
			logx.L.Debugf("❌ Previous error detected")
			// read another instance from the channel
			continue
		}

		// set this instance property if Vm is not SSH-reachable
		ok, err := util.IsSshConfiguredVmSshReachable(data.Node)
		if err != nil || !ok {
			// data.Err = fmt.Errorf("VM %s not SSH reachable or misconfigured: %v", data.Node, err)
			data.Err = err
			logx.L.Debugf("❌ Previous error detected")
		}

		out <- data
	}
}
