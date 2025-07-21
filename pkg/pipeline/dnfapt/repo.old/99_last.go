/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package packagex

import (
	"fmt"

	"github.com/abtransitionit/luc/pkg/logx"
)

// # Purpose
//
// Last stage reads (in <-chan) data (of type PipelineData) from the channel
//
// # Notes
//
// - No need to close the channel. Only senders should close the channel
func lastStep(in <-chan PipelineData) error {
	// loop over each instance of PipelineData in the channel
	for data := range in {
		// if an error exits
		if data.Err != nil {
			logx.L.Debugf("Pipeline error: %v", data.Err)
			continue
		}
		// if no error exits : log information - one per structure
		logx.L.Infof("[%s] Received instance", data.HostName)
		fmt.Println(data.String())
	}
	return nil
}
