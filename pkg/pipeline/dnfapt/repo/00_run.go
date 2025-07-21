/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package repo

import (
	"strings"

	"github.com/abtransitionit/luc/pkg/logx"
)

const RunPipelineDescription = "provision OS dnfapt repositories on VMs."

func RunPipeline(vmList string, repositories []string) (string, error) {
	logx.L.Debug(RunPipelineDescription)

	// define var
	vms := strings.Fields(vmList) // convert ListAsString to []string (ie. go slice)
	// nbVm := len(vms)

	// Define the pipeline channels
	ch01 := make(chan PipelineData)
	// ch02 := make(chan PipelineData)
	// ch03 := make(chan PipelineData)
	chOutLast := ch01

	// stage running async/concurrently/in parallel
	go source(ch01, vms, repositories) // define data instances to send to the pipeline

	// final sequential step. collects all instances in the pipeline and build a sumary
	err := lastStep(chOutLast)
	if err != nil {
		return "", err
	}
	// success
	return "", nil
}

// packages := strings.Fields(packageList) // convert ListAsString to []string (ie. go slice)
