/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package kbe

import (
	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/abtransitionit/luc/pkg/util"
)

const CheckDescription = "check KIND clusters."

func check(arg ...string) (string, error) {
	logx.L.Info(CheckDescription)

	// check the cluster
	logx.L.Info("checking the cluster")

	// list clusters
	cli := "kind get clusters"
	output, err := util.RunCLILocal(cli)
	if err != nil {
		logx.L.Debugf("❌ Error detected")
		return "", err
	}
	logx.L.Infof("list of clusters : %s", output)

	// list nodes of default cluster
	cli = "kind get nodes"
	output, err = util.RunCLILocal(cli)
	if err != nil {
		logx.L.Debugf("❌ Error detected")
		return "", err
	}

	logx.L.Infof("list of nodes for default cluster : %s", output)

	// success
	return "", nil
}
