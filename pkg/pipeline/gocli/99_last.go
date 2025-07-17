/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package gocli

import (
	"fmt"

	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/abtransitionit/luc/pkg/util"
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
func lastStep(in <-chan PipelineData, vms []string) error {
	// common loop
	for data := range in {
		if data.Err != nil {
			logx.L.Debugf("❌ [%s] Pipeline error : %v", data.Config.Name, data.Err)
			continue
		}
		logx.L.Infof("[%s] [%S] Received Pipeline Data", data.Config.Name, data.HostName)
		fmt.Println(data.String())
	}

	// specific actions: for each vm display the PATH to all CLI(s) installed
	for _, vm := range vms {
		// play code
		path, err := util.GetSubdirRemote("/usr/local/bin", vm)
		// error
		if err != nil {
			return err
		}
		// success
		logx.L.Infof("[%s] export PATH=%s:$PATH", vm, path)
	}
	return nil
}
