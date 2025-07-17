/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package kind

import (
	"github.com/abtransitionit/luc/internal/config"
	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/abtransitionit/luc/pkg/pipeline/dnfapt/packagex"
	"github.com/abtransitionit/luc/pkg/util"
)

const DaPackDescription = "provision OS dnfapt package(s) on VM(s)."

func daPack(arg ...string) (string, error) {
	logx.L.Info(DaPackDescription)

	// get all Map:key as []string
	listPackage := util.GetMapKeys(config.KindDaCliConfigMap)

	// launch this pipeline
	_, err := packagex.RunPipeline(config.KindVm, listPackage)
	if err != nil {
		logx.L.Debugf("%s", err)
		return "", err
	}

	// success
	return "", nil
}

// listPackage := ""
// for key := range config.KindDaCliConfigMap {
// 	listPackage += key + " "
// }
// listPackage = strings.TrimSpace(listPackage)
