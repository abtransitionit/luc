/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package update

import (
	"github.com/abtransitionit/luc/pkg/logx"
)

// First/Source step of the pipeline that define the data to be pipelined
func Source(out chan<- PipelineData, cliName string) {
	// define goroutine
	go func() {
		// close channel
		defer close(out)

		// declaradef variable
		data := PipelineData{}

		// declare the var that will be pipelined (i.e shared/process by all process of the pipeline)
		// TODO

		// log information
		logx.L.Debugf("defining data to be pipelined")
		logx.L.Debugf("pipelined data defined")

		// Step 2: send data to next step
		out <- data
	}()
}

// logx.L.Infow("Loaded CLI config", "cli", cliName, "url", SingleCliConfig.Url)
// logx.L.Infof("URl is %s", SingleCliConfig.Url)
// logx.L.Errorw("CLI config not found", "cli", cliName)
