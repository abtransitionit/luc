/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package update

import (
	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/abtransitionit/luc/pkg/util"
)

func infoBefore(in <-chan PipelineData, out chan<- PipelineData) {
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

		// step 2: get property
		osKernelVersion, err := util.GetLocalProperty("oskversion")
		if err != nil {
			data.Err = err
			logx.L.Debugf("❌ Error detected")
		}
		data.OskernelVersionBefore = osKernelVersion

		// send the instance to the channel (for next stage/step)
		out <- data
	}
}
