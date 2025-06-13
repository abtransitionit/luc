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

		// get config for this CLI
		SingleCliConfig, ok := config.GetCLIConfig(cliName)

		// declaradef variable
		data := PipelineData{}

		// do the job
		if !ok {
			data.Err = fmt.Errorf("CLI not found in CLI config map: %s", cliName)
			logx.L.Debugf("Error detetected")
		} else {
			// define property
			data.Config = SingleCliConfig
			// log information
			logx.L.Infof("Loaded %s CLI config", cliName)
		}

		// send data to next step
		out <- data
	}()
}

// logx.L.Infow("Loaded CLI config", "cli", cliName, "url", SingleCliConfig.Url)
// logx.L.Infof("URl is %s", SingleCliConfig.Url)
// logx.L.Errorw("CLI config not found", "cli", cliName)
