/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package rc

import (
	"fmt"
	"strings"

	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/abtransitionit/luc/pkg/util"
)

func source(out chan<- PipelineData, vms []string, rcFilePath string, stringLine string) {
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
			data.Err = fmt.Errorf("%v, %s", err, osFamily)
			logx.L.Debugf("[%s] ❌ Error detected 1", vm)
		}

		osUser, err := util.GetPropertyRemote(vm, "osuser")
		if err != nil {
			data.Err = fmt.Errorf("%v, %s", err, osUser)
			logx.L.Debugf("[%s] ❌ Error detected 1", vm)
		}

		// define instance property - 1 per VmxService
		data.HostName = vm
		data.OsFamily = osFamily
		data.osUser = osUser
		data.RcFilePath = rcFilePath
		data.Line = stringLine
		// log and send
		logx.L.Debugf("[%s] defined instance. Sending it to the pipeline", vm)
		out <- data
	}

}
