/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package provision

import (
	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/abtransitionit/luc/pkg/util"
)

// # Purpose
//
// - This stage create an instance of the structure to be pipelined
// - 1 instance of the structure per item in the listName (e.g 9 cli => 9 structures)
// - This stage will send (out chan<-) each instance into the channel
func source(out chan<- PipelineData, listName ...string) {
	// close channel when this code ended
	// closing it make it available for next stage
	// because it is defined outside
	defer close(out)

	// log information
	logx.L.Debugf("defining data to be pipelined")

	// loop over all items in the LIST
	for _, item := range listName {
		// create a new instance per item
		data := PipelineData{}

		// get some OS property
		osFamily, err := util.OsPropertyGet("osfamily")
		if err != nil {
			data.Err = err
			logx.L.Debugf("❌ Error detected")
		}

		hostType, err := util.OsPropertyGet("host")
		if err != nil {
			data.Err = err
			logx.L.Debugf("❌ Error detected")
		}

		// set this instance properties
		data.Name = item
		data.OsFamily = osFamily
		data.HostType = hostType

		// send the instance to the channel (for next stage or final step)
		out <- data
	}

	// log information
	logx.L.Debugf("pipelined data defined")

	// out <- data
}
