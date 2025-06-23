/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package reboot

import (
	"strconv"

	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/abtransitionit/luc/pkg/util"
)

func reboot(in <-chan PipelineData, out chan<- PipelineData) {
	logx.L.Debugf("reboot")
	defer close(out)
	for data := range in {

		if data.Err != nil {
			out <- data
			logx.L.Debugf("❌ Previous error detected")
			continue
		}

		// map string to bool
		a := data.NeedReboot
		boolVal, err := strconv.ParseBool(a)
		if err != nil {
			data.Err = err
			out <- data
			logx.L.Debugf("❌ Previous error detected")
			continue
		}

		// check reboot is needed
		if !boolVal {
			logx.L.Debugf("No reboot needed")
			out <- data
			continue
		}

		// Here: reboot is needed
		logx.L.Debugf("Rebooting now")
		_, err = util.Reboot()
		if err != nil {
			data.Err = err
			logx.L.Debugf("❌ Error detected")
			out <- data
			continue
		}

		// send
		out <- data
	}
}
