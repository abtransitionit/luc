/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package gocli

import (
	"fmt"

	"github.com/abtransitionit/luc/pkg/config"
	"github.com/abtransitionit/luc/pkg/logx"
)

// # Purpose
//
// - This stage create an instance of the structure  to be pipelined
// - 1 instance of the structure per item in the cliNameList (e.g 9 cli => 9 structures)
// - This stage will send (out chan<-) each instance into the channel
func source(out chan<- PipelineData, cliNameList ...string) {
	// close channel when this code ended
	// closing it make it available for next stage
	// because it is defined outside
	defer close(out)

	// log information
	logx.L.Debugf("defining data to be pipelined")

	// loop over all CLI
	for _, item := range cliNameList {
		// create a new instance per CLI
		data := PipelineData{}

		// Fetch the config for this CLI
		CliConfig, ok := config.GetCLIConfigMap(item)
		if !ok {
			// add this property (if error) to the instance structure
			data.Err = fmt.Errorf("❌ CLI not found in config map")
			logx.L.Debugf("❌ Error detected")
		} else {
			// add this property to the pipelined data
			data.Config = CliConfig
			// log information
			logx.L.Infof("[%s] Loaded CLI config", item)
		}

		// send the instance to the channel (for next stage or final step)
		out <- data
	}

}
