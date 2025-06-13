/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package update

import (
	"github.com/abtransitionit/luc/pkg/logx"
)

// Move File or folder to final destination
func EndPipeline(in <-chan PipelineData, out chan<- PipelineData) {
	go func() {
		// close channel
		defer close(out)

		for data := range in {
			// Step 1: propagate error if any
			if data.Err != nil {
				out <- data
				continue
			}

			// Step 2: Print a msg
			logx.L.Infof("End of piepeline")

			// Step 3: send result to next pipeline step
			out <- data
		}
	}()
}
