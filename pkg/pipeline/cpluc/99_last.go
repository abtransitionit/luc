/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package cpluc

import (
	"fmt"

	"github.com/abtransitionit/luc/pkg/logx"
)

// # Purpose
//
// - This is the last step after all other stages
// - This process is not a goroutine, it is a standard function
// - It reads (in <-chan) each instance in the channel and process them (often says OK or error)
//
// # Notes
//
// - No need to close the channel. Only senders should close the channel
func lastStep(in <-chan PipelineData) error {
	// loop over each instance of PipelineData in the channel
	for data := range in {
		// if an error exits
		if data.Err != nil {
			logx.L.Debugf("❌ [%s] Pipeline error : %v", data.VmName, data.Err)
			continue
		}
		// if no error exits : log information - one per structure
		logx.L.Infof("[%s] Received Pipeline Data", data.VmName)
		fmt.Println(data.String())
	}
	return nil
}
