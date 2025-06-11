/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package gocli

import (
	"github.com/abtransitionit/luc/pkg/logx"
)

func GenericUrl(in <-chan PipelineData, out chan<- PipelineData) {
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
			// get this property
			data.GenericUrl = data.Config.Url
			logx.L.Infof("Generic URL : '%s'", data.GenericUrl)

			// send data to next step
			out <- data
		}
	}()
}

// logx.L.Infow("Generic URL passed through", "cli", data.Config.Name, "url", data.GenericUrl)
