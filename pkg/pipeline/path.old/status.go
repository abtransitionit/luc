/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package oservice

import (
	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/abtransitionit/luc/pkg/util"
)

func statusService(in <-chan PipelineData, out chan<- PipelineData) {
	defer close(out)

	for data := range in {
		if data.Err != nil {
			out <- data
			logx.L.Debugf("❌ Previous error detected")
			continue
		}

		// get OS service properties
		logx.L.Debugf("[%s] [%s] geting service status", data.HostName, data.Config.Name)
		serviceInfos, err := util.GetPropertyRemote(data.HostName, "serviceinfos", data.Config.Name)

		// error
		if err != nil {
			logx.L.Debugf("[%s][%s] ❌ Error detected 1", data.Config.Name, data.HostName)
			data.Err = err
			out <- data
			continue
		}

		// success
		data.ServiceInfos = serviceInfos
		logx.L.Debugf("[%s] [%s] got service status", data.HostName, data.Config.Name)
		out <- data
	}
}
