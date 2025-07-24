/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package repo

import (
	"fmt"
	"strings"

	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/abtransitionit/luc/pkg/util"
)

func updateRepo(in <-chan PipelineData, out chan<- PipelineData) {
	defer close(out)

	for data := range in {

		vm := data.HostName
		repoName := data.Config.Name

		if data.Err != nil {
			out <- data
			logx.L.Debugf("❌ Previous error detected")
			continue
		}

		// define cli based on OS:family
		var cli string
		osFamily := strings.TrimSpace(strings.ToLower(data.OsFamily))

		switch osFamily {
		case "rhel", "fedora":
			cli = `sudo dnf update -qq -y > /dev/null`
		case "debian":
			cli = `sudo apt update -qq -y > /dev/null`
		default:
			data.Err = fmt.Errorf("unknown OS family [%s]", osFamily)
			logx.L.Debugf("[%s] [%s] ❌ error detected 1", repoName, vm)
			out <- data
			continue
		}

		// play cli
		logx.L.Debugf("[%s] [%s] updating dnfapt repo", vm, repoName)
		cliOut, err := util.RunCLIRemote(data.HostName, cli)

		// error
		if err != nil {
			data.Err = fmt.Errorf("%v, %s", err, cliOut)
			logx.L.Debugf("[%s] [%s] ❌ Error detected 2", vm, repoName)
			out <- data
			continue
		}

		// success
		logx.L.Debugf("[%s] [%s] updated dnfapt repo", vm, repoName)
		out <- data
	}
}
