/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package gocli

import (
	"fmt"

	configinternal "github.com/abtransitionit/luc/internal/config"
	"github.com/abtransitionit/luc/pkg/config"
	"github.com/abtransitionit/luc/pkg/logx"
)

// # Purpose
//
// - This stage create an instance of the structure to be pipelined
// - 1 instance of the structure per item in the cliMap (e.g 9 cliMap => 1 structure => 9 instances)
// - This stage will send (out chan<-) each instance into the channel
func source(out chan<- PipelineData, cliMap map[string]configinternal.CustomCLIConfig) {
	// close channel when this code ended
	// closing it make it available for next stage
	// because it is defined outside
	defer close(out)

	// log information
	logx.L.Debugf("defining data to be pipelined")

	// loop over all items in the list
	for _, item := range cliMap {
		// create a new instance per item
		data := PipelineData{}

		// Fetch the config for this CLI
		CliConfig, ok := config.GetCLIConfigMap(item.Name)
		if !ok {
			// add this property (if error) to the instance structure
			data.Err = fmt.Errorf("❌ CLI not found in config map")
			logx.L.Debugf("❌ Error detected")
		} else {
			// add this property to the pipelined data
			data.Config = CliConfig
			data.Version = item.Version
			data.DstFolder = item.DstFolder
			// log information
			logx.L.Infof("[%s] Loaded CLI config", item.Name)
		}

		// send the instance to the channel (for next stage or final step)
		out <- data
	}

}
