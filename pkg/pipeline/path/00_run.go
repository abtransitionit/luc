/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package path

import (
	"strings"

	"github.com/abtransitionit/luc/pkg/logx"
)

const RunPipelineDescription = "Create Temporary file with envar PATH."

func RunPipeline(vmList string, pathFile string) (string, error) {
	logx.L.Debug(RunPipelineDescription)

	// define var
	vms := strings.Fields(vmList) // convert ListAsString to slice ([]string)

	// Define the pipeline channels
	ch01 := make(chan PipelineData)
	ch02 := make(chan PipelineData)
	chOutLast := ch02

	// stage running async/concurrently/in parallel
	go source(ch01, vms, pathFile) // define instances to send to the pipeline
	go getPath(ch01, ch02)         // define instances to send to the pipeline

	// final sequential step. collects all instances in the pipeline and build a sumary
	err := lastStep(chOutLast)
	if err != nil {
		return "", err
	}
	// success
	return "", nil
}
