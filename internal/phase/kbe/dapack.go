/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package kbe

import (
	"github.com/abtransitionit/luc/internal/config"
	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/abtransitionit/luc/pkg/pipeline/dnfapt/packagex"
	"github.com/abtransitionit/luc/pkg/util"
)

const DaPackDescription = "provision Dnfapt package(s)"

func daPack(arg ...string) (string, error) {
	logx.L.Info(DaPackDescription)

	// get all Map:key as []string
	listPackage := util.GetMapKeys(config.KbeDnfaptCliConfigMap)

	// launch this pipeline
	_, err := packagex.RunPipeline(config.KbeListNode, listPackage)
	if err != nil {
		logx.L.Debugf("%s", err)
		return "", err
	}

	// success
	return "", nil
}
