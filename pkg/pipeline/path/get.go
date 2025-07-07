/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package path

import (
	"fmt"

	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/abtransitionit/luc/pkg/util"
)

func getPath(in <-chan PipelineData, out chan<- PipelineData) {
	defer close(out)

	for data := range in {
		// get instance property
		vm := data.HostName
		tmpPath := data.TmpFilePath

		if data.Err != nil {
			out <- data
			logx.L.Debugf("❌ Previous error detected")
			continue
		}

		// Update envar PATH
		var pathExtend string
		var err error
		logx.L.Debugf("[%s] creating envar PATH with tree path", vm)
		cli := fmt.Sprintf(`luc util getprop pathext '%s' `, data.Path)
		if pathExtend, err = util.RunCLIRemote(vm, cli); err != nil {
			logx.L.Debugf("[%s] ❌ Error detected 1", vm)
			data.Err = err
			out <- data
			continue
		}

		// save path to file
		logx.L.Debugf("[%s] persisting envar PATH to file", vm)
		cli = fmt.Sprintf(`luc util strfile '%s' %s false --force`, pathExtend, tmpPath)
		if _, err = util.RunCLIRemote(vm, cli); err != nil {
			logx.L.Debugf("[%s] ❌ Error detected 2", vm)
			data.Err = err
			out <- data
			continue
		}
		// set instance property
		data.Path = pathExtend
		// success
		logx.L.Debugf("[%s] saved envar PATH to file", vm)
		out <- data
	}
}
