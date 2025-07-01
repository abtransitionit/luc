/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package gocli

import (
	"fmt"

	"github.com/abtransitionit/luc/pkg/logx"
)

// # Purpose
//
// - This stage (like all other stage) get each instance in the piepeline, one by one
// - It processes the instance, and send the result to the next pipeline stage
func GenericUrl(in <-chan PipelineData, out chan<- PipelineData) {
	defer close(out)
	// loop over each key of the instance structure PipelineData
	for data := range in {
		// if this instance property exists
		if data.Err != nil {
			// send the instance instance into the channel (for next stage/step)
			out <- data
			// log information
			logx.L.Debugf("❌ Previous error detected")
			// read another instance from the channel
			continue
		}
		// set this instance property
		if data.Config.Url == "" {
			data.Err = fmt.Errorf("❌ git URL is empty")
			logx.L.Debugf("❌Error detected 4")
			out <- data
			continue
		}

		data.GenericUrl = data.Config.Url

		// send the instance to the channel (for next stage/step)
		out <- data
	}
}
