/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package kbe

import (
	"os"

	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/abtransitionit/luc/pkg/util"
)

const CreateDescription = "create KIND cluster."

func create(arg ...string) (string, error) {
	logx.L.Info(CreateDescription)

	// set env var
	logx.L.Info("setting env var")
	err := os.Setenv("CNI_PATH", "/usr/local/bin/cni")
	if err != nil {
		panic(err)
	}

	// create the cluster
	logx.L.Info("creating the cluster")
	cli := "kind create cluster"
	output, err := util.RunCLILocal(cli)
	if err != nil {
		logx.L.Debugf("❌ Error detected")
		return "", err
	}

	// success
	return output, nil
}
