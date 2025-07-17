/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package rc

import (
	"fmt"

	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/abtransitionit/luc/pkg/util"
)

func addLine(in <-chan PipelineData, out chan<- PipelineData) {
	defer close(out)

	for data := range in {

		// check error
		if data.Err != nil {
			out <- data
			logx.L.Debugf("❌ Previous error detected")
			continue
		}

		// get instance property
		vm := data.HostName
		user := data.osUser

		// TODO: weird: explicit the user. remote enable linger for current user (aka. sudo user)
		logx.L.Debugf("[%s] [%s] adding line to user RC file", vm, user)
		cli := fmt.Sprintf(`luc do AddLineToFile %s '%s'`, data.RcFilePath, data.Line)
		if outp, err := util.RunCLIRemote(vm, cli); err != nil {
			data.Err = fmt.Errorf("%v, %s", err, outp)
			logx.L.Debugf("❌ [%s] [%s] Error detected 1", vm, user)
			out <- data
			continue
		}

		// success
		logx.L.Debugf("[%s] [%s] addded line to user RC file", vm, user)
		out <- data
	}
}

// // Enable lingering for the user
// logx.L.Debugf("enabling lingering for user %s", osUser)
// if err := util.EnableUserService(osUser); err != nil {
// 	logx.L.Debugf("❌ Error detected 6")
// 	return "", err
// }
// // logx.L.Infof("lingering enabled for user %-5s", osUser)
