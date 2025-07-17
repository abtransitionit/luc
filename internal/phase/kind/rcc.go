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
	rcp "github.com/abtransitionit/luc/pkg/pipeline/rc"
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

		// get persisted PATH as a base64 string
		logx.L.Infof("[%s] gettting persisted path as base64 encoded string from %s.", vm, pathTmpFile)
		cli := fmt.Sprintf("luc do GetStringFromFile  %s false", pathTmpFile) // get persisted PATH as base64 string
		pathBase64, err := util.RunCLIRemote(vm, cli)
		if err != nil {
			logx.L.Debugf("❌ %v: %s", err, pathBase64)
			return "", err
		}

		// decode this base64 string to a txt string
		pathString, err := base64.StdEncoding.DecodeString(strings.TrimSpace(pathBase64))
		if err != nil {
			logx.L.Debugf("❌ %v", err)
			return "", err
		}

		// log
		logx.L.Infof("[%s] got persisted path: %s", vm, pathString)
		logx.L.Infof("adding line (PATH definition) to rc custom file %s.", customRcFilePath)

		// add line - using a pipeline
		lineToAdd := fmt.Sprintf("export PATH=%s", string(pathString))
		_, err = rcp.RunPipeline(vm, customRcFilePath, lineToAdd)
		if err != nil {
			logx.L.Debugf("❌ %s", err)
			return "", err
		}

		logx.L.Infof("added line (PATH definition) to rc custom file")

		// _, err := rc.RunPipeline(config.KindVm, RcFilePath, line)
		// if err != nil {
		// 	logx.L.Debugf("%s", err)
		// 	return "", err
		// }

		// logx.L.Infof("added line to rc custom file")
	} // vm loop

	// success
	return "", nil
}
