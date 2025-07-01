/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package gocli

import (
	"fmt"

	"github.com/abtransitionit/luc/pkg/config"
	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/abtransitionit/luc/pkg/util"
)

// # Purpose
//
// - This stage create an instance of the structure to be pipelined
// - 1 instance of the structure per item in the cliMap (e.g 9 cli => 9 instances)
// - This stage will send (out chan<-) each instance into the channel
func source(out chan<- PipelineData, vms []string, cliMap config.CustomCLIConfigMap) {
	// close channel when this code ended
	// closing it make it available for next stage, because it is defined outside
	defer close(out)

	// define var
	nbVm := len(vms)
	nbCli := len(cliMap)
	nbWorker := nbVm

	// log
	logx.L.Debugf("defining instances to be pipelined")
	logx.L.Debugf("Vms        to provision:  %d : %s", nbVm, vms)
	logx.L.Debugf("CLI(s) to install per VM: %d : %s", nbCli, util.GetMapKeys(cliMap))
	logx.L.Debugf("Will use a many workers as VM: %d ", nbWorker)

	// loop over all items in the list
	for _, item := range cliMap {
		// create an instance per item (CLI)
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

// for _, item := range cliMap {
// 	// create an instance per item (CLI)
// 	data := PipelineData{}

// 	// Fetch the config for this CLI
// 	CliConfig, ok := config.GetCLIConfig(item.Name)
// 	if !ok {
// 		// add this property (if error) to the instance structure
// 		data.Err = fmt.Errorf("❌ CLI not found in config map")
// 		logx.L.Debugf("❌ Error detected")
// 	} else {
// 		// add this property to the pipelined data
// 		data.Config = CliConfig
// 		data.Version = item.Version
// 		data.DstFolder = item.DstFolder
// 		// log information
// 		logx.L.Infof("[%s] Loaded CLI config", item.Name)
// 	}
// 	// log information
// 	// logx.L.Debugf("[%s] defined data instances to be pipelined", vm)
// 	// send the instance to the channel (for next stage or final step)
// 	out <- data
// }
