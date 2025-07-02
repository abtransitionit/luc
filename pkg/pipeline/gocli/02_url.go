/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package gocli

import (
	"strings"

	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/abtransitionit/luc/pkg/util"
)

func setUrlSpec(in <-chan PipelineData, out chan<- PipelineData) {
	defer close(out)

	for data := range in {
		if data.Err != nil {
			out <- data
			logx.L.Debugf("❌ Previous error detected")
			continue
		}
		// Get some properties
		uname, err := util.GetRemoteProperty("uname", data.HostName)
		if err != nil {
			data.Err = err
			logx.L.Debugf("[%s] ❌ Error detected 1", data.HostName)
		}

		osArch, err := util.GetRemoteProperty("osarch", data.HostName)
		if err != nil {
			data.Err = err
			logx.L.Debugf("[%s] ❌ Error detected 2", data.HostName)
		}

		osType, err := util.GetRemoteProperty("ostype", data.HostName)
		if err != nil {
			data.Err = err
			logx.L.Debugf("[%s] ❌ Error detected 3", data.HostName)
		}

		// define var from instance property
		url := data.GenericUrl

		// replace placeholders
		url = strings.ReplaceAll(url, "$NAME", data.Config.Name)
		url = strings.ReplaceAll(url, "$TAG", data.Version)
		url = strings.ReplaceAll(url, "$OS", osType)
		url = strings.ReplaceAll(url, "$ARCH", osArch)
		url = strings.ReplaceAll(url, "$UNAME", uname)

		// set instance property
		data.HostUrl = url

		// log
		logx.L.Debugf("[%s] setted Url Spec", data.Config.Name)

		// send
		out <- data
	}
}
