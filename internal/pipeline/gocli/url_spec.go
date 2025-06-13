/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package gocli

import (
	"os/exec"
	"runtime"
	"strings"

	"github.com/abtransitionit/luc/pkg/logx"
)

func SpecificUrl(in <-chan PipelineData, out chan<- PipelineData) {
	go func() {
		// close channel
		defer close(out)

		for data := range in {
			// propagate error if any
			if data.Err != nil {
				out <- data
				continue
			}

			// replace placeholders
			url := data.GenericUrl
			url = strings.ReplaceAll(url, "$NAME", data.Config.Name)
			url = strings.ReplaceAll(url, "$TAG", data.Config.Tag)
			url = strings.ReplaceAll(url, "$OS", runtime.GOOS)
			url = strings.ReplaceAll(url, "$ARCH", runtime.GOARCH)
			url = strings.ReplaceAll(url, "$UNAME", getUnameM())

			// step 2: define property
			data.SpecificUrl = url

			// log information
			logx.L.Infof("Specific URL: '%s'", data.SpecificUrl)

			// step 3: send pipeline var to next pipeline step
			out <- data
		}
	}()
}

func getUnameM() string {
	out, err := exec.Command("uname", "-m").Output()
	if err != nil {
		return "unknown"
	}
	return strings.TrimSpace(string(out))
}

// logx.L.Infow("Specific URL generated", "cli", data.Config.Name, "specificUrl", data.SpecificUrl)
// logx.L.Infow("Specific URL generated: '%s'", data.SpecificUrl)
