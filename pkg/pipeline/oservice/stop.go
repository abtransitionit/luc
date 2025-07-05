/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package oservice

import (
	"fmt"

	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/abtransitionit/luc/pkg/util"
)

func stopService(in <-chan PipelineData, out chan<- PipelineData) {
	defer close(out)

	for data := range in {
		if data.Err != nil {
			out <- data
			logx.L.Debugf("❌ Previous error detected")
			continue
		}
		// get instance property
		vm := data.HostName

		// remote start service
		logx.L.Debugf("[%s] [%s] stoping service", vm, data.Config.Name)
		cli := fmt.Sprintf(`luc util oservice stop %s --local --force`, data.Config.SName)
		_, err := util.RunCLIRemote(vm, cli)

		// error
		if err != nil {
			logx.L.Debugf("[%s][%s] ❌ Error detected 1", vm, data.Config.Name)
			data.Err = err
			out <- data
			continue
		}

		// success
		logx.L.Debugf("[%s] [%s] stoped service", vm, data.Config.Name)
		out <- data
	}
}
