/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package upgrade

import (
	"strings"

	"github.com/abtransitionit/luc/pkg/logx"
)

const RunPipelineDescription = "upgrade OS repositories and packages to version latest"

func RunPipeline(vmList string) (string, error) {
	logx.L.Debug(RunPipelineDescription)
	// define var
	vms := strings.Fields(vmList) // convert ListAsString to []string (ie. go slice)
	nbWorker := len(vms)          // as many workers as VMs
	// Define the pipeline channels
	ch01 := make(chan PipelineData)
	ch02 := make(chan PipelineData)
	ch03 := make(chan PipelineData)
	chOutLast := ch03

	// aync stage
	go source(ch01, vms) // define instances to send to the pipeline
	go rUpgrade(ch01, ch02, nbWorker)
	go remoteReboot(ch02, ch03, nbWorker)

	// final sequential step. collects all instances in the pipeline and build a sumary
	err := lastStep(chOutLast)
	if err != nil {
		return "", err
	}
	// success
	return "", nil
}
