/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package reboot

import (
	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/abtransitionit/luc/pkg/util"
)

func needReboot(in <-chan PipelineData, out chan<- PipelineData) {
	defer close(out)
	for data := range in {
		if data.Err != nil {
			out <- data
			logx.L.Debugf("❌ Previous error detected")
			continue
		}

		// step 2: get property
		needReboot, err := util.GetLocalProperty("reboot")
		if err != nil {
			data.Err = err
			logx.L.Debugf("❌ Error detected")
		}
		data.NeedReboot = needReboot

		// send
		out <- data
	}
}
