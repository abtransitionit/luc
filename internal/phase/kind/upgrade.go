/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package kind

import (
	"github.com/abtransitionit/luc/internal/config"
	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/abtransitionit/luc/pkg/pipeline/dnfapt/rupgrade"
)

const UpgradeDescription = "provision OS nodes with latest dnfapt packages and repositories."

func upgrade(arg ...string) (string, error) {
	logx.L.Info(UpgradeDescription)
	_, err := rupgrade.RunPipeline(config.KindVm)
	if err != nil {
		return "", err
	}
	// on SUCCESS
	return "", nil
}
