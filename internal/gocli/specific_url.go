/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package gocli

import (
	"runtime"
	"strings"

	"github.com/abtransitionit/luc/pkg/logx"
)

func SpecificUrl(in <-chan PipelineData, out chan<- PipelineData) {
	go func() {
		// close channel
		defer close(out)

		// get config for this CLI - Did something gets wrong earlier
		for data := range in {
			if data.Err != nil {
				// send data to next step
				out <- data
				// Keep reading data from channel
				continue
			}
			// placeholders are replaced
			url := data.GenericUrl
			url = strings.ReplaceAll(url, "$NAME", data.Config.Name)
			url = strings.ReplaceAll(url, "$TAG", data.Config.Tag)
			url = strings.ReplaceAll(url, "$OS", runtime.GOOS)
			url = strings.ReplaceAll(url, "$ARCH", runtime.GOARCH)

			data.SpecificUrl = url

			// logx.L.Infow("Specific URL generated: '%s'", data.SpecificUrl)
			logx.L.Infof("Specific URL generated: '%s'", data.SpecificUrl)

			// send data to next step
			out <- data
		}
	}()
}

// logx.L.Infow("Specific URL generated", "cli", data.Config.Name, "specificUrl", data.SpecificUrl)
