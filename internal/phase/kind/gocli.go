/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package kind

import (
	"github.com/abtransitionit/luc/internal/config"
	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/abtransitionit/luc/pkg/pipeline/gocli"
)

const GoCliDescription = "provision Go CLI"

func goCli(arg ...string) (string, error) {
	logx.L.Info(CpLucDescription)
	_, err := gocli.RunPipeline(config.KindVm, config.KindGoCliConfigMap)
	if err != nil {
		logx.L.Debugf("%s", err)
		return "", err
	}
	return "", nil
}
