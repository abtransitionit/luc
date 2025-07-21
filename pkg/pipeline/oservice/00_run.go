/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package oservice

import (
	"strings"

	"github.com/abtransitionit/luc/pkg/config"
	"github.com/abtransitionit/luc/pkg/logx"
)

const RunPipelineDescription = "Create 1..n Linux OS service unit file(s) on 1..n VMS."

func RunPipeline(vmList string, osServiceMap config.OsServiceConfigMap) (string, error) {
	logx.L.Debug(RunPipelineDescription)

	// define var
	vms := strings.Fields(vmList) // convert ListAsString to []string (ie. go slice)

	// Define the pipeline channels
	ch01 := make(chan PipelineData)
	ch02 := make(chan PipelineData)
	ch03 := make(chan PipelineData)
	ch04 := make(chan PipelineData)
	chOutLast := ch04

	// stage running async/concurrently/in parallel
	go source(ch01, vms, osServiceMap) // define instances to send to the pipeline
	go createUnit(ch01, ch02)
	go startService(ch02, ch03)
	go statusService(ch03, ch04)

	// final sequential step. collects all instances in the pipeline and build a sumary
	err := lastStep(chOutLast)
	if err != nil {
		return "", err
	}
	// success
	return "", nil
}
