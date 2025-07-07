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

		if data.Err != nil {
			out <- data
			logx.L.Debugf("❌ Previous error detected")
			continue
		}

		// Update envar PATH
		var pathExtend string
		var err error
		logx.L.Debugf("[%s] updating envar PATH with tree path", vm)
		logx.L.Debugf("data.Path is : %s", data.Path)
		cli := fmt.Sprintf(`luc util getprop pathext '%s' `, data.Path)
		if pathExtend, err = util.RunCLIRemote(vm, cli); err != nil {
			logx.L.Debugf("[%s] ❌ Error detected 1", vm)
			data.Err = err
			out <- data
			continue
		}

		// set instance property
		data.Path = pathExtend
		// success
		logx.L.Debugf("[%s] updated envvar PATH", vm)
		out <- data
	}
}
