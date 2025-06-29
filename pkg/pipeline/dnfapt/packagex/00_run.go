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
	vms := strings.Fields(vmList) // convert ListAsString to slice ([]string)
	// nbVm := len(vms)
	packages := strings.Fields(packageList) // convert ListAsString to slice ([]string)
	// Define the pipeline channels
	ch01 := make(chan PipelineData)
	chOutLast := ch01

	// aync stage
	go source(ch01, vms, packages) // define instances to send to the pipeline

	// final sequential step. collects all instances in the pipeline and build a sumary
	err := lastStep(chOutLast)
	if err != nil {
		return "", err
	}
	// success
	return "", nil
}
