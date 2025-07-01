/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package gocli

import (
	"os/exec"
	"strings"

	"github.com/abtransitionit/luc/pkg/logx"
)

func setUrlSpec(in <-chan PipelineData, out chan<- PipelineData) {
	defer close(out)

	for data := range in {
		if data.Err != nil {
			out <- data
			logx.L.Debugf("❌ Previous error detected")
			continue
		}

		// define var from instance property
		url := data.GenericUrl
		// replace placeholders
		url = strings.ReplaceAll(url, "$NAME", data.Config.Name)
		url = strings.ReplaceAll(url, "$TAG", data.Version)
		// url = strings.ReplaceAll(url, "$OS", runtime.GOOS)
		// url = strings.ReplaceAll(url, "$ARCH", runtime.GOARCH)
		// url = strings.ReplaceAll(url, "$UNAME", getUnameM())

		// // set this instance property
		data.AppUrl = url

		// log information
		logx.L.Debugf("[%s] setted Url Spec", data.Config.Name)
		// send
		out <- data
	}
}

func getUnameM() string {
	out, err := exec.Command("uname", "-m").Output()
	if err != nil {
		return "unknown"
	}
	return strings.TrimSpace(string(out))
}

// logx.L.Infow("Specific URL generated", "cli", data.Config.Name, "AppUrl", data.AppUrl)
// logx.L.Infow("Specific URL generated: '%s'", data.AppUrl)
