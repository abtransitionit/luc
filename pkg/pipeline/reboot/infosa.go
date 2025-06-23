/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package reboot

import (
	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/abtransitionit/luc/pkg/util"
)

func infoAfter(in <-chan PipelineData, out chan<- PipelineData) {
	defer close(out)
	for data := range in {
		if data.Err != nil {
			out <- data
			logx.L.Debugf("❌ Previous error detected")
			continue
		}

		// step 2: get property
		osKernelVersion, err := util.OsPropertyGet("oskversion")
		if err != nil {
			data.Err = err
			logx.L.Debugf("❌ Error detected")
		}
		data.OskernelVersionAfter = osKernelVersion

		// send
		out <- data
	}
}
