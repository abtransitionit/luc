/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package cpluc

import (
	"github.com/abtransitionit/luc/pkg/logx"
)

const RunPipelineDescription = "build and provision LUC."

func RunPipeline(vmList string) (string, error) {
	logx.L.Debug(RunPipelineDescription)

	// Define the pipeline channels
	ch01 := make(chan PipelineData)
	ch02 := make(chan PipelineData)
	chOutLast := ch02

	// sequential step
	var dtpip PipelineData
	dtpip.localOutput = "/tmp/luc"
	dtpip.localExePath = "/usr/local/bin/luc"
	dtpip.localOutXptf = "/tmp/luc-linux"
	// build and deploy locally
	buildLuc(dtpip)

	// aync stage
	go source(ch01, dtpip, vmList) // define instances to send to the pipeline
	go scpLuc(ch01, ch02, vmList)  // scp luc from local to remote

	// final sequential step. collects all instances in the pipeline and build a sumary
	err := lastStep(chOutLast)
	if err != nil {
		return "", err
	}
	// success
	return "", nil
}
