/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package oservice

import (
	"fmt"
	"strings"

	"github.com/abtransitionit/luc/pkg/config"
	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/abtransitionit/luc/pkg/util"
)

func source(out chan<- PipelineData, vms []string, osServiceMap config.OsServiceConfigMap) {
	defer close(out)

	// define var
	nbVm := len(vms)
	nbService := len(osServiceMap)

	// log
	logx.L.Debugf("defining instances to be pipelined")
	logx.L.Debugf("Vms        to provision:  %d : %s", nbVm, vms)
	logx.L.Debugf("Services(s) to install per VM: %d : %s", nbVm, nbService)

	// loop over each CLI
	for _, item := range osServiceMap {
		// create an instance per item
		data := PipelineData{}

		// Fetch the shared public config for this CLI
		osServiceConfig := item
		// loop over each VM
		for _, vm := range vms {
			vm = strings.TrimSpace(vm)
			if vm == "" {
				continue
			}

			// get some OS property
			osFamily, err := util.GetPropertyRemote(vm, "osfamily")
			if err != nil {
				data.Err = fmt.Errorf("❌ Error: %v, %s", err, osFamily)
				logx.L.Debugf("[%s] ❌ Error detected 1", vm)
			}
			osType, err := util.GetPropertyRemote(vm, "ostype")
			if err != nil {
				data.Err = fmt.Errorf("❌ Error: %v, %s", err, osType)
				logx.L.Debugf("[%s] ❌ Error detected 1", vm)
			}

			// avoid creating instance for Os type not manage
			if strings.ToLower(strings.TrimSpace(osType)) != "linux" {
				logx.L.Debugf("[%s] [%s] ⚠️ Os type not managed", vm, osType)
				continue
			}

			// avoid creating instance for services not manage by this OS family
			if isExcluded, err := ServiceIsExcluded(osFamily, data.Config.Name); err != nil {
				logx.L.Debugf("❌ Error :%v", err)
			} else if isExcluded {
				logx.L.Debugf("[%s] [%s] Service excluded, skipping...", vm, osServiceConfig.Name)
				continue
			}

			// define instance property - 1 per VmxService
			data.HostName = vm
			data.Config = osServiceConfig
			data.OsFamily = osFamily

			// log information
			logx.L.Debugf("[%s] [%s] Loaded service config. Sending instance to the pipeline", item.Name, vm)
			// sen this instance to the channel
			out <- data
		}

	} // for

}
