/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package gocli

import (
	"strings"

	"github.com/abtransitionit/luc/pkg/config"
	"github.com/abtransitionit/luc/pkg/logx"
)

const RunPipelineDescription = "install GO CLIs on VMs."

func RunPipeline(vmList string, cliMap config.CustomCLIConfigMap) (string, error) {
	logx.L.Debug(RunPipelineDescription)

	// define var
	vms := strings.Fields(vmList) // convert ListAsString to slice ([]string)
	// nbVm := len(vms)
	nbWorker := len(vms)

	// // Count and log the number of CLI args
	// logx.L.Debugf("Received %d CLI(s) to provisioned", len(cliMap))

	// Define the pipeline channels
	ch01 := make(chan PipelineData)
	ch02 := make(chan PipelineData)
	ch03 := make(chan PipelineData)
	ch04 := make(chan PipelineData)
	ch05 := make(chan PipelineData)
	ch06 := make(chan PipelineData)
	// ch07 := make(chan PipelineData)
	chOutLast := ch06

	// log
	logx.L.Debugf("Will use a many workers as VM: %d ", nbWorker)

	// aync stage (i.e running concurrently/in parallel)
	go source(ch01, vms, cliMap) // define instances to send to the pipeline
	go setUrlSpec(ch01, ch02)
	go setArtifact(ch02, ch03)
	go getArtifact(ch03, ch04, nbWorker) // get artifact
	go unTgz(ch04, ch05, nbWorker)
	go Move(ch05, ch06, nbWorker) // move to final destination

	// final sequential step. collects all instances in the pipeline and build a sumary
	err := lastStep(chOutLast, vms)
	if err != nil {
		return "", err
	}
	// success
	return "", nil
}
