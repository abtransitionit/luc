/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package update

import (
	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/abtransitionit/luc/pkg/util"
)

func infoAfter(in <-chan PipelineData, out chan<- PipelineData) {
	logx.L.Info("Enter infoAfter")
	defer close(out) // close channel when done
	// loop over each item of type PipelineData in the channel
	for data := range in {
		// Step 1: propagate error if any
		if data.Err != nil {
			logx.L.Debugf("❌ Previous error detected %v", data.Err)
			out <- data
			continue
		}

		// step 2: get property
		osKernelVersion, err := util.OsPropertyGet("oskversion")
		if err != nil {
			data.Err = err
			logx.L.Debugf("❌ Error detected")
		}
		data.OskernelVersionAfter = osKernelVersion

		// step 3: send pipeline var to next pipeline step
		out <- data
	}
}

// logx.L.Infow("Specific URL generated", "cli", data.Config.Name, "specificUrl", data.SpecificUrl)
