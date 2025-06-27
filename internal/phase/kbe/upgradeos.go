/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package kbe

import (
	"github.com/abtransitionit/luc/internal/config"
	"github.com/abtransitionit/luc/pkg/logx"
	rupgrade "github.com/abtransitionit/luc/pkg/pipeline/dnfapt/rupgrade"
)

const UpgradeOsDescription = "provision OS nodes with latest dnfapt packages and repositories."

func upgradeOs(arg ...string) (string, error) {
	logx.L.Info(UpgradeOsDescription)
	// Launch the pipeline attach to this phase
	err := rupgrade.RunPipeline(config.KbeListNode)
	if err != nil {
		return "", err
	}
	// on SUCCESS
	return "", nil
}
