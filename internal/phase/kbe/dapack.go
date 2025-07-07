/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package kbe

import (
	"strings"

	"github.com/abtransitionit/luc/internal/config"
	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/abtransitionit/luc/pkg/pipeline/dnfapt/packagex"
)

const DaPackDescription = "provision Dnfapt package(s)"

func daPack(arg ...string) (string, error) {
	logx.L.Info(DaPackDescription)

	// Build StringList of packages
	listPackage := ""
	for key := range config.KbeDnfaptCliConfigMap {
		listPackage += key + " "
	}
	listPackage = strings.TrimSpace(listPackage)

	// launch this pipeline
	_, err := packagex.RunPipeline(config.KbeListNode, listPackage)
	if err != nil {
		logx.L.Debugf("%s", err)
		return "", err
	}

	// success
	return "", nil
}
