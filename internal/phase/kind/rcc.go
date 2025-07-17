/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package kind

import (
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/abtransitionit/luc/internal/config"
	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/abtransitionit/luc/pkg/util"
)

const RccDescription = "Add lines to non-root user custom RC file."

func rcc(arg ...string) (string, error) {
	logx.L.Info(RccDescription)

	// define var
	customRcFilePath := "$HOME/.profile.luc "
	pathTmpFile := config.KindPathFile

	// get persisted envar PATH (on remote) - not using a pipeline
	vms := strings.Fields(config.KindVm)
	for _, vm := range vms {
		logx.L.Infof("[%s] gettting persisted path as base64 encoded string from %s.", vm, pathTmpFile)
		cli := fmt.Sprintf("luc do GetStringFromFile  %s false", pathTmpFile)
		pathBase64, err := util.RunCLIRemote(vm, cli)
		if err != nil {
			logx.L.Debugf("%v: %s", err, pathBase64)
			return "", err
		}
		// logx.L.Infof("[%s] base64 encoded persisted path: %s", vm, outp)
		// decode string
		pathString, err := base64.StdEncoding.DecodeString(strings.TrimSpace(outp))
		if err != nil {
			logx.L.Debugf("%v", err)
			return "", err
		}
		// log
		logx.L.Infof("[%s] got persisted path: %s", vm, string(pathString))

		logx.L.Infof("adding line to rc file %s.", customRcFilePath)

	} // vm loop

	// do
	logx.L.Infof("adding line to rc file %s.", customRcFilePath)

	// success
	logx.L.Infof("added line to rc file")
	return "", nil
}
