/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package oservice

import (
	"fmt"

	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/abtransitionit/luc/pkg/util"
)

func startService(in <-chan PipelineData, out chan<- PipelineData) {
	defer close(out)

	for data := range in {
		if data.Err != nil {
			out <- data
			logx.L.Debugf("❌ Previous error detected")
			continue
		}

		// remote start service
		logx.L.Debugf("[%s] [%s] start service", data.HostName, data.Config.Name)
		cli := fmt.Sprintf(`luc util oservice status %s --local --force`, data.Config.SName)
		_, err := util.RunCLIRemote(cli, data.HostName)

		// error
		if err != nil {
			logx.L.Debugf("[%s][%s] ❌ Error detected 1", data.Config.Name, data.HostName)
			data.Err = err
			out <- data
			continue
		}

		// success
		logx.L.Debugf("[%s] [%s] started service file", data.HostName, data.Config.Name)
		out <- data
	}
}
