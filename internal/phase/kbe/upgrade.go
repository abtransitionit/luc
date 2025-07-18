/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package kbe

import (
	"github.com/abtransitionit/luc/internal/config"
	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/abtransitionit/luc/pkg/pipeline/dnfapt/rupgrade"
)

const UpgradeDescription = "provision OS nodes with latest dnfapt packages and repositories."

func upgrade(arg ...string) (string, error) {
	logx.L.Info(UpgradeDescription)
	// launch this pipeline
	_, err := rupgrade.RunPipeline(config.KbeListNode)
	if err != nil {
		return "", err
	}
	// on SUCCESS
	return "", nil
}
