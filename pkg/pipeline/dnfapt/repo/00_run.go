/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package repo

import (
	"strings"

	"github.com/abtransitionit/luc/pkg/config"
	"github.com/abtransitionit/luc/pkg/logx"
)

const RunPipelineDescription = "provision 1..n OS dnfapt repository"

func RunPipeline(vmList string, repoMap config.CustomDnfaptRepoConfigMap) (string, error) {
	logx.L.Debug(RunPipelineDescription)

	// define var
	vms := strings.Fields(vmList) // convert ListAsString to []string (ie. go slice)
	nbWorker := len(vms)
	// nbWorker := len(vms)

	// Define the pipeline channels
	ch01 := make(chan PipelineData)
	ch02 := make(chan PipelineData)
	ch03 := make(chan PipelineData)
	ch04 := make(chan PipelineData)
	ch05 := make(chan PipelineData)
	ch06 := make(chan PipelineData)
	chOutLast := ch06

	// stage running async/concurrently/in parallel
	// go source(ch01, vms, repositories) // define data instances to send to the pipeline
	go source(ch01, vms, repoMap)        // define data instances to send to the pipeline
	go setUrl(ch01, ch02)                // define specific URLs
	go setPath(ch02, ch03)               // define specific Path
	go saveRepoFile(ch03, ch04)          // define and save the repo file on the FS
	go saveGpgFile(ch04, ch05, nbWorker) // define and save the repo file on the FS
	go updateRepo(ch05, ch06)            // define and save the repo file on the FS
	// go setPath(ch02, ch03)        // define repo and gpg OS filepath
	// go addRepo(ch02, ch03, nbWorker) // add repos

	// final sequential step. collects all instances in the pipeline and build a sumary
	err := lastStep(chOutLast)
	if err != nil {
		return "", err
	}
	// success
	return "", nil
}
