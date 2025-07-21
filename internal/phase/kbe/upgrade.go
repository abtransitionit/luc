/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package kbe

import (
	"github.com/abtransitionit/luc/internal/config"
	"github.com/abtransitionit/luc/pkg/logx"
	upgradep "github.com/abtransitionit/luc/pkg/pipeline/dnfapt/upgrade"
)

const UpgradeDescription = "provision OS nodes with latest dnfapt packages and repositories."

func upgrade(arg ...string) (string, error) {
	logx.L.Info(UpgradeDescription)
	_, err := upgradep.RunPipeline(config.KbeListNode)
	if err != nil {
		return "", err
	}
	return "", nil
}
