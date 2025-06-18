/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package gocliold

import (
	"github.com/abtransitionit/luc/pkg/logx"
)

func GenericUrl(in <-chan PipelineData, out chan<- PipelineData) {
	go func() {
		defer close(out) // close channel when done
		// loop over each item of type PipelineData in the channel
		for data := range in {
			// propagate error if any
			if data.Err != nil {
				// send data to next step
				out <- data
				// Keep reading data from channel
				continue
			}
			// step 2: define property
			data.GenericUrl = data.Config.Url
			// log information
			logx.L.Infof("Generic URL: '%s'", data.GenericUrl)

			// step 3: send pipeline var to next pipeline step
			out <- data
		}
	}()
}

// logx.L.Infow("Generic URL passed through", "cli", data.Config.Name, "url", data.GenericUrl)
