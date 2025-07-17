/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package kbe

import (
	"github.com/abtransitionit/luc/internal/config"
	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/abtransitionit/luc/pkg/pipeline/cpluc"
)

const CpLucDescription = "provision LUC CLI"

func cpLuc(arg ...string) (string, error) {
	logx.L.Info(CpLucDescription)
	_, err := cpluc.RunPipeline(config.KbeListNode)
	if err != nil {
		logx.L.Debugf("%s", err)
		return "", err
	}
	return "", nil
}
