/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package linger

import (
	"strings"

	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/abtransitionit/luc/pkg/util"
)

func source(out chan<- PipelineData, vms []string) {
	defer close(out)

	// define var
	nbVm := len(vms)

	// log
	logx.L.Debugf("defining instances to be pipelined")
	logx.L.Debugf("Vms        to provision:  %d : %s", nbVm, vms)

	// loop over each CLI
	for _, item := range vms {
		// create an instance per item
		data := PipelineData{}

		// get property
		vm := strings.TrimSpace(item)
		if vm == "" {
			continue
		}

		osFamily, err := util.GetPropertyRemote(vm, "osfamily")
		if err != nil {
			data.Err = err
			logx.L.Debugf("[%s] ❌ Error detected 1", vm)
		}

		osType, err := util.GetPropertyRemote(vm, "ostype")
		if err != nil {
			data.Err = err
			logx.L.Debugf("[%s] ❌ Error detected 1", vm)
		}
		osUser, err := util.GetPropertyRemote(vm, "osuser")
		if err != nil {
			data.Err = err
			logx.L.Debugf("[%s] ❌ Error detected 1", vm)
		}
		UserLinger, err := util.GetPropertyRemote(vm, "userlinger", osUser)
		if err != nil {
			data.Err = err
			logx.L.Debugf("[%s] ❌ Error detected 1", vm)
		}

		// avoid creating instance for Os type not manage
		if strings.ToLower(strings.TrimSpace(osType)) != "linux" {
			logx.L.Debugf("[%s] [%s] ⚠️ Os type not managed", vm, osType)
			continue
		}

		// define instance property - 1 per VmxService
		data.HostName = vm
		data.OsFamily = osFamily
		data.osUser = osUser
		data.LingerStatus = UserLinger

		// log and send
		logx.L.Debugf("[%s] sending instance to the pipeline", vm)
		out <- data
	}

}
