/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package update

import (
	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/abtransitionit/luc/pkg/util/dnfapt"
)

func update(in <-chan PipelineData, out chan<- PipelineData) {
	defer close(out)
	for data := range in {
		if data.Err != nil {
			out <- data
			logx.L.Debugf("❌ Previous error detected")
			continue
		}

		// step 2:Update the OS
		_, err := dnfapt.UpdateOs()
		if err != nil {
			data.Err = err
			logx.L.Debugf("Error detected")
			out <- data
			continue
		}

		// send
		out <- data
	}
}
