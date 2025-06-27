/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package reboot

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
	osFamily, err := util.GetLocalProperty("osfamily")
	if err != nil {
		data.Err = err
		logx.L.Debugf("❌ Error detected")
	}

	osDistro, err := util.GetLocalProperty("osdistro")
	if err != nil {
		data.Err = err
		logx.L.Debugf("❌ Error detected")
	}

	hostType, err := util.GetLocalProperty("host")
	if err != nil {
		data.Err = err
		logx.L.Debugf("❌ Error detected")
	}

	osVersion, err := util.GetLocalProperty("osversion")
	if err != nil {
		data.Err = err
		logx.L.Debugf("❌ Error detected")
	}

	// set this instance properties
	data.OsFamily = osFamily
	data.OsDistro = osDistro
	data.HostType = hostType
	data.OsVersion = osVersion

	// log information
	// logx.L.Debugf("[%s] defined data instances to be pipelined", vm)
	// send the instance to the channel (for next stage or final step)
	out <- data
}
