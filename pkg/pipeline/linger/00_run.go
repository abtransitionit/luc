/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package linger

import (
	"strings"

	"github.com/abtransitionit/luc/pkg/logx"
)

const RunPipelineDescription = "Allow non 11.1 root user to run OS services"

func RunPipeline(vmList string) (string, error) {
	logx.L.Debug(RunPipelineDescription)

	// define var
	vms := strings.Fields(vmList) // convert ListAsString to []string (ie. go slice)

	// Define the pipeline channels
	ch01 := make(chan PipelineData)
	ch02 := make(chan PipelineData)
	chOutLast := ch02

	// aync stage (i.e running concurrently/in parallel)
	go source(ch01, vms)        // define instances to send to the pipeline
	go enableLinger(ch01, ch02) // define instances to send to the pipeline

	// final sequential step. collects all instances in the pipeline and build a sumary
	err := lastStep(chOutLast)
	if err != nil {
		return "", err
	}
	// success
	return "", nil
}
