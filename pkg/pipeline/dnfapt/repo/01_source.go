/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package repo

import (
	"fmt"
	"strings"

	"github.com/abtransitionit/luc/pkg/config"
	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/abtransitionit/luc/pkg/util"
)

// # Purpose
//
// - This stage create an instance of the structure to be pipelined
// - 1 instance of the structure per item in the cliNameList (e.g 9 cli => 9 structures)
// - This stage will send (out chan<-) each instance into the channel
func source(out chan<- PipelineData, vms []string, repoMap config.CustomDnfaptRepoConfigMap) {
	// close channel when this code ended
	// closing it make it available for next stage, because it is defined outside
	defer close(out)

	// define var
	nbVm := len(vms)
	nbRepo := len(repoMap)

	// log
	logx.L.Debugf("defining instances to be pipelined")
	logx.L.Debugf("VM to provision        : %2d : %s", nbVm, vms)
	logx.L.Debugf("Repo to install per VM : %2d : %s", nbRepo, util.GetMapKeys(repoMap))

	// loop over each repository
	for _, item := range repoMap {
		repoName := item.Name

		// create an instance per item
		data := PipelineData{}

		// Fetch the shared public config for this repository
		repoConfig, ok := config.GetDnfapteRepoConfig(repoName)
		if !ok {
			data.Err = fmt.Errorf("[%s] ❌ repository not found in config map", repoName)
			logx.L.Debugf("[%s] ❌ Error detected", repoName)
			out <- data
			continue
		}

		// loop over each VM
		for _, vm := range vms {

			vm = strings.TrimSpace(vm)

			if vm == "" {
				continue
			}

			// avoid creating instance for non SSH reachable VMs
			result, err := util.GetPropertyLocal("sshreachability", vm)
			if err != nil {
				logx.L.Debugf("⚠️ %v : %s : %s", err, result, "skipping data instance for it")
				continue
			} else if strings.ToLower(strings.TrimSpace(result)) == "false" {
				logx.L.Debugf("⚠️ [%s] remote vm is not reachable, skipping data instance for it", vm)
				continue
			}

			// get some properties
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

			// define instance property - 1 per VMxCLI
			data.HostName = vm
			data.Config = repoConfig
			data.CName = item.CName
			data.Version = item.Version
			data.GenericUrlRepo = repoConfig.UrlRepo
			data.GenericUrlGpg = repoConfig.UrlGpg
			data.HostName = vm
			data.OsFamily = osFamily
			data.OsDistro = osDistro
			data.HostType = hostType
			data.OsVersion = osVersion
			data.RepositoryList = util.GetMapKeys(repoMap)
			data.OskernelVersionBefore = kernelVersion

			// success
			logx.L.Debugf("[%s] [%s] Loaded repo config. Sending instance to the pipeline", vm, repoName)
			out <- data
		}

	} // for
}
