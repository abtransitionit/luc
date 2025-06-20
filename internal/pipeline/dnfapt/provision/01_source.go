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
// - 1 instance of the structure per item in the cliNameList (e.g 9 cli => 9 structures)
// - This stage will send (out chan<-) each instance into the channel
func source(out chan<- PipelineData, packageNameList ...string) {
	// close channel when this code ended
	// closing it make it available for next stage
	// because it is defined outside
	defer close(out)
	data := PipelineData{}

	// log information
	logx.L.Debugf("defining data to be pipelined")

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
	data.OsFamily = osFamily
	data.HostType = hostType

	// log information
	logx.L.Debugf("pipelined data defined")

	out <- data
}
