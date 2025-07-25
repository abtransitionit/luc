/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package path

import (
	"fmt"
	"strings"

	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/abtransitionit/luc/pkg/util"
)

func source(out chan<- PipelineData, vms []string, pathFile string) {
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

		PathTree, err := util.GetPropertyRemote(vm, "pathtree", "/usr/local/bin")
		if err != nil {
			data.Err = fmt.Errorf("%v, %s", err, PathTree)
			logx.L.Debugf("[%s] ❌ Error detected 1", vm)
		}

		// define instance property - 1 per VmxService
		data.HostName = vm
		data.Path = PathTree
		data.TmpFilePath = pathFile

		// log and send
		logx.L.Debugf("[%s] sending instance to the pipeline", vm)
		out <- data
	}

}
