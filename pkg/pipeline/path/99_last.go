/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package path

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
	// common loop
	for data := range in {
		vm := data.HostName
		if data.Err != nil {
			logx.L.Debugf("❌ [%s] Pipeline error : %v", vm, data.Err)
			continue
		}
		logx.L.Infof("[%s] Received Pipeline Data", vm)
		fmt.Println(data.String())
	}

	return nil
}
