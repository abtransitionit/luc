/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package kbe

import (
	"github.com/abtransitionit/luc/internal/config"
	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/abtransitionit/luc/pkg/pipeline/dnfapt/repo"
	"github.com/abtransitionit/luc/pkg/util"
)

const DaRepoDescription = "provision dnfapt repositories"

func daRepo(arg ...string) (string, error) {
	logx.L.Info(DaRepoDescription)

	// get all Map:key as []string
	listRepository := util.GetMapKeys(config.KbeDnfaptRepoConfigMap)

	// launch this pipeline
	_, err := repo.RunPipeline(config.KbeListNode, listRepository)
	if err != nil {
		logx.L.Debugf("%s", err)
		return "", err
	}

	// success
	return "", nil
}
