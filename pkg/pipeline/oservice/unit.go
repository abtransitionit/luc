/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package oservice

import (
	"fmt"

	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/abtransitionit/luc/pkg/util"
)

func createUnit(in <-chan PipelineData, out chan<- PipelineData) {
	defer close(out)

	for data := range in {
		vm := data.HostName
		if data.Err != nil {
			out <- data
			logx.L.Debugf("❌ Previous error detected")
			continue
		}

		// remote create file:service
		stringContent := data.Config.Content
		filePath := data.Config.Path
		logx.L.Debugf("[%s] [%s] creating service file", vm, data.Config.Name)
		cli := fmt.Sprintf(`luc do ServiceCreateUnitFile '%s' %s`, stringContent, filePath)
		outp, err := util.RunCLIRemote(vm, cli)

		// error
		if err != nil {
			data.Err = fmt.Errorf("%v, %s", err, outp)
			logx.L.Debugf("❌ [%s][%s] Error detected 1", data.Config.Name, vm)
			out <- data
			continue
		}

		// success
		logx.L.Debugf("[%s] [%s] created service file", vm, data.Config.Name)
		out <- data
	}
}
