/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package repo

import (
	"fmt"
	"strings"

	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/abtransitionit/luc/pkg/util"
)

// # Purpose
//
// - This stage create an instance of the structure to be pipelined
// - 1 instance of the structure per item in the cliNameList (e.g 9 cli => 9 structures)
// - This stage will send (out chan<-) each instance into the channel
func source(out chan<- PipelineData, vms []string, repositorises []string) {
	// close channel when this code ended
	// closing it make it available for next stage, because it is defined outside
	defer close(out)

	logx.L.Debugf("defining instances to be pipelined")
	logx.L.Debugf("Vms        to provision.       : %d : %s", len(vms), vms)
	logx.L.Debugf("repository(s) to install per VM: %d : %s", len(repositorises), repositorises)

	for _, vm := range vms {
		vm = strings.TrimSpace(vm)

		if vm == "" {
			continue
		}
		// define one per item
		data := PipelineData{}

		// get some OS property
		osFamily, err := util.GetPropertyRemote(vm, "osfamily")
		if err != nil {
			data.Err = fmt.Errorf("❌ Error: %v, %s", err, osFamily)
			logx.L.Debugf("[%s] ❌ Error detected 1", vm)
		}

		osType, err := util.GetPropertyRemote(vm, "ostype")
		if err != nil {
			data.Err = fmt.Errorf("❌ Error: %v, %s", err, osType)
			logx.L.Debugf("[%s] ❌ Error detected 1", vm)
		}

		osDistro, err := util.GetPropertyRemote(vm, "osdistro")
		if err != nil {
			data.Err = fmt.Errorf("❌ Error: %v, %s", err, osDistro)
			logx.L.Debugf("[%s] ❌ Error detected 2", vm)
		}

		hostType, err := util.GetPropertyRemote(vm, "host")
		if err != nil {
			data.Err = fmt.Errorf("❌ Error: %v, %s", err, hostType)
			logx.L.Debugf("[%s] ❌ Error detected 3", vm)
		}

		osVersion, err := util.GetPropertyRemote(vm, "osversion")
		if err != nil {
			data.Err = fmt.Errorf("❌ Error: %v, %s", err, osVersion)
			logx.L.Debugf("[%s] ❌ Error detected 4", vm)
		}

		kernelVersion, err := util.GetPropertyRemote(vm, "oskversion")
		if err != nil {
			data.Err = fmt.Errorf("❌ Error: %v, %s", err, kernelVersion)
			logx.L.Debugf("[%s] ❌ Error detected 5", vm)
		}

		// avoid creating instance for Os:type not manage
		if strings.ToLower(strings.TrimSpace(osType)) != "linux" {
			logx.L.Debugf("[%s] [%s] ⚠️ Os type not managed", vm, osType)
			continue
		}

		// avoid creating instance for Os:family not managed
		family := strings.ToLower(strings.TrimSpace(osFamily))
		if family != "debian" && family != "rhel" && family != "fedora" {
			logx.L.Debugf("[%s] [%s] ⚠️ OS family not managed", vm, osFamily)
			continue
		}

		// set this instance properties
		data.HostName = vm
		data.OsFamily = osFamily
		data.OsDistro = osDistro
		data.HostType = hostType
		data.OsVersion = osVersion
		data.RepositoryList = repositorises
		data.OskernelVersionBefore = kernelVersion

		// log information
		logx.L.Debugf("[%s] sending instance to the pipeline", vm)
		// sen this instance to the channel
		out <- data

	} // for
}
