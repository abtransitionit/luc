/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package update

import (
	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/abtransitionit/luc/pkg/util"
)

// # Purpose
//
// Source stage will send (out chan<-) data of type PipelineData to the channel
func source(out chan<- PipelineData) {
	defer close(out)
	data := PipelineData{}

	// log information
	logx.L.Debugf("defining data to be pipelined")

	osFamily, err := util.OsPropertyGet("osfamily")
	if err != nil {
		data.Err = err
		logx.L.Debugf("❌ Error detected")
	}

	osDistro, err := util.OsPropertyGet("osdistro")
	if err != nil {
		data.Err = err
		logx.L.Debugf("❌ Error detected")
	}

	hostType, err := util.OsPropertyGet("host")
	if err != nil {
		data.Err = err
		logx.L.Debugf("❌ Error detected")
	}

	osVersion, err := util.OsPropertyGet("osversion")
	if err != nil {
		data.Err = err
		logx.L.Debugf("❌ Error detected")
	}

	data.OsFamily = osFamily
	data.OsDistro = osDistro
	data.HostType = hostType
	data.OsVersion = osVersion

	// log information
	logx.L.Debugf("pipelined data defined")

	out <- data
}
