/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package gocli

import (
	"fmt"

	"github.com/abtransitionit/luc/pkg/config"
	"github.com/abtransitionit/luc/pkg/logx"
)

// First/Source step of the pipeline that define the data to be pipelined
func CliName(out chan<- PipelineData, cliName string) {
	// define goroutine
	go func() {
		// close channel
		defer close(out)

		// log information
		logx.L.Debugf("defining data to be pipelined")

		// get config for this CLI
		SingleCliConfig, ok := config.GetCLIConfigMap(cliName)

		// declare the var that will be pipelined (i.e shared/process by all process of the pipeline)
		data := PipelineData{}

		// step1: define the source data
		if !ok {
			data.Err = fmt.Errorf("CLI not found in CLI config map: %s", cliName)
			logx.L.Debugf("Error detetected")
		} else {
			// define property
			data.Config = SingleCliConfig
			// log information
			logx.L.Infof("Loaded %s CLI config", cliName)
		}

		// Log information
		logx.L.Debugf("✅ pipelined data defined")

		// step 2: send pipeline var to next pipeline step
		out <- data
	}()
}

// logx.L.Infow("Loaded CLI config", "cli", cliName, "url", SingleCliConfig.Url)
// logx.L.Infof("URl is %s", SingleCliConfig.Url)
// logx.L.Errorw("CLI config not found", "cli", cliName)
