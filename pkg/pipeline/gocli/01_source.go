/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package gocli

import (
	"fmt"
	"strings"

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

	// log
	logx.L.Debugf("defining instances to be pipelined")
	logx.L.Debugf("VM to provision      : %2d : %s", nbVm, vms)
	logx.L.Debugf("CLI to install per VM: %2d : %s", nbCli, util.GetMapKeys(cliMap))

	// loop over each CLI
	for _, item := range cliMap {
		cliName := item.Name
		// create an instance per item
		data := PipelineData{}

		// Fetch the shared public config for this CLI
		cliConfig, ok := config.GetCLIConfig(cliName)
		if !ok {
			data.Err = fmt.Errorf("[%s] ❌ CLI not found in config map", cliName)
			logx.L.Debugf("[%s] ❌ Error detected", cliName)
			out <- data
			continue
		}

		// loop over each VM
		for _, vm := range vms {
			vm = strings.TrimSpace(vm)
			if vm == "" {
				continue
			}
			// define instance property - 1 per VMxCLI
			data.HostName = vm
			data.Config = cliConfig
			data.Version = item.Version
			data.DstFolder = item.DstFolder
			data.GenericUrl = cliConfig.Url
			// log information
			logx.L.Debugf("[%s] [%s] Loaded CLI config. Sending instance to the pipeline", cliName, vm)
			// sen this instance to the channel
			out <- data
		}

	} // for

}
