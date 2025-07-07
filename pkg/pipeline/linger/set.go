/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package linger

import (
	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/abtransitionit/luc/pkg/util"
)

func enableLinger(in <-chan PipelineData, out chan<- PipelineData) {
	defer close(out)

	for data := range in {
		if data.Err != nil {
			out <- data
			logx.L.Debugf("❌ Previous error detected")
			continue
		}
		// get instance property
		vm := data.HostName
		user := data.osUser

		// TODO: weird: explicit the user. remote enable linger for current user (aka. sudo user)
		logx.L.Debugf("[%s] [%s] enabling linger for user ", vm, user)
		cli := `luc util oservice linger --local --force`
		if _, err := util.RunCLIRemote(vm, cli); err != nil {
			logx.L.Debugf("[%s][%s] ❌ Error detected 1", vm, user)
			data.Err = err
			out <- data
			continue
		}

		// get property
		lingerStatus, err := util.GetPropertyRemote(vm, "userlinger", user)
		if err != nil {
			logx.L.Debugf("[%s][%s] ❌ Error detected 2", vm, user)
			data.Err = err
			out <- data
			continue
		}

		// set instance property
		data.LingerStatus = lingerStatus
		// success
		logx.L.Debugf("[%s] [%s] enabled linger for user", vm, user)
		out <- data
	}
}
