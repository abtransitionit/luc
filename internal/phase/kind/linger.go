/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package kind

import (
	"github.com/abtransitionit/luc/internal/config"
	"github.com/abtransitionit/luc/pkg/logx"
	lingerP "github.com/abtransitionit/luc/pkg/pipeline/linger"
)

const LingerDescription = "Allow non root user to run OS services."

func linger(arg ...string) (string, error) {
	_, err := lingerP.RunPipeline(config.KindVm)
	if err != nil {
		logx.L.Debugf("%s", err)
		return "", err
	}
	return "", nil
}
