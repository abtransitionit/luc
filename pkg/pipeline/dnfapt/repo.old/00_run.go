/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package packagex

import (
	"strings"

	"github.com/abtransitionit/luc/pkg/logx"
)

const RunPipelineDescription = "install OS dnfapt packages on VMs."

func RunPipeline(vmList string, packageList string) (string, error) {
	logx.L.Debug(RunPipelineDescription)
	// define var
	vms := strings.Fields(vmList) // convert ListAsString to []string (ie. go slice)
	nbVm := len(vms)
	packages := strings.Fields(packageList) // convert ListAsString to []string (ie. go slice)
	// Define the pipeline channels
	ch01 := make(chan PipelineData)
	ch02 := make(chan PipelineData)
	ch03 := make(chan PipelineData)
	chOutLast := ch03

	// aync stage
	go source(ch01, vms, packages) // define instances to send to the pipeline
	go remoteInstall(ch01, ch02, nbVm, vms, packages)
	go remoteReboot(ch02, ch03, nbVm)

	// final sequential step. collects all instances in the pipeline and build a sumary
	err := lastStep(chOutLast)
	if err != nil {
		return "", err
	}
	// success
	return "", nil
}
