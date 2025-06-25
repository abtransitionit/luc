/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package kbe

import (
	"github.com/abtransitionit/luc/internal/config"
	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/abtransitionit/luc/pkg/pipeline/rfilecopy"
)

const CpLucDescription = "provision needed go CLI"

func cpLuc(arg ...string) (string, error) {
	logx.L.Info(CpLucDescription)
	// Launch the pipeline attach to this phase
	err := rfilecopy.RunPipeline(config.KbeListNode, "/tmp/luc-linux", "/tmp/luc")
	if err != nil {
		return "", err
	}
	// on SUCCESS
	return "", nil
}
