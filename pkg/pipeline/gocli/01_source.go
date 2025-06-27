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
// - This stage creates a PipelineData instance for each CLI in the cliMap
// - e.g 9 CLI in the cliMap => 9 instances of the structure PipelineData
// - It sends (out chan<-) each instance into the output channel
func source(out chan<- PipelineData, cliMap map[string]configinternal.CustomCLIConfig) {
	// close channel when this code ended
	// closing it make it available for next stage
	// because it is defined outside
	defer close(out)

	// log information
	logx.L.Debugf("defining data to be pipelined")

	// loop over all items in the list
	for _, item := range cliMap {
		// create a new instance per item (CLI)
		data := PipelineData{}

		// Fetch the config for this CLI
		CliConfig, ok := config.GetCLIConfig(item.Name)
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
		// log information
		// logx.L.Debugf("[%s] defined data instances to be pipelined", vm)
		// send the instance to the channel (for next stage or final step)
		out <- data
	}

}
