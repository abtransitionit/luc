/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package kind

import (
	"fmt"
	"strings"

	"github.com/abtransitionit/luc/internal/config"
	"github.com/abtransitionit/luc/pkg/logx"
	rcp "github.com/abtransitionit/luc/pkg/pipeline/rc"
	"github.com/abtransitionit/luc/pkg/util"
)

const RcDescription = "Add a line to non-root user RC file."

func rc(arg ...string) (string, error) {
	logx.L.Info(RcDescription)

	// define var
	RcFilePath := "$HOME/.bashrc"
	customRcFilePath := "$HOME/.profile.luc "
	line := fmt.Sprintf("source %s", customRcFilePath)

	// create a custom RC user file (on remote) - not using a pipeline
	vms := strings.Fields(config.KindVm)
	for _, vm := range vms {
		logx.L.Debugf("[%s] creating file on remote: %s", vm, customRcFilePath)
		cli := fmt.Sprintf("luc do TouchFile %s", customRcFilePath)
		outp, err := util.RunCLIRemote(vm, cli)
		if err != nil {
			logx.L.Debugf("%v: %s", err, outp)
			return "", err
		}
		logx.L.Debugf("[%s] created file on remote: %s", vm, customRcFilePath)
	} // vm loop

	// add one line to user RC file (on remote) - using a pipeline
	_, err := rcp.RunPipeline(config.KindVm, RcFilePath, line)
	if err != nil {
		logx.L.Debugf("%s", err)
		return "", err
	}

	// success
	return "", nil
}
