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
	// ch03 := make(chan PipelineData)
	chOutLast := ch02

	var dtpip PipelineData
	dtpip.localExePath = "/usr/local/bin/luc"
	dtpip.localOutput = "/tmp/luc"
	dtpip.localOutXptf = "/tmp/luc-linux"

	buildLuc(dtpip)                // build and deploy locally
	go source(ch01, dtpip, vmList) // define instances to send to the pipeline
	go scpLuc(ch01, ch02, vmList)  // scp luc from local to remote
	// go mvLucRemote(ch01, ch03, vmList) // mv remotely from tmp to final path

	// This is not a stage but the last foreground step reading all instance in the pipeline
	err := lastStep(chOutLast)
	if err != nil {
		return "", err
	}
	// on SUCCESS
	// time.Sleep(10 * time.Second)
	return "", nil
}
