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

	// // Count and log the number of CLI args
	// logx.L.Debugf("Received %d CLI(s) to provisioned", len(cliMap))

	// Define the pipeline channels
	ch01 := make(chan PipelineData)
	// chOutGenericUrl := make(chan PipelineData)
	// chOutSpecificUrl := make(chan PipelineData)
	// chOutArtifactName := make(chan PipelineData)
	// chOutArtifactPath := make(chan PipelineData)
	// chOutArtifactGet := make(chan PipelineData)
	// chOutFileGuessType := make(chan PipelineData)
	// chOutFileSave := make(chan PipelineData)
	// chOutFileUnTgz := make(chan PipelineData)
	// chOutFileMove := make(chan PipelineData)
	// chOutBuildPath := make(chan PipelineData)
	chOutLast := ch01

	// aync stage (i.e running concurrently/in parallel)
	go source(ch01, vms, cliMap) // define instances to send to the pipeline
	// go GenericUrl(chOutSource, chOutGenericUrl)           // set property
	// go SpecificUrl(chOutGenericUrl, chOutSpecificUrl)     // set property
	// go ArtifactName(chOutSpecificUrl, chOutArtifactName)  // set property
	// go ArtifactPath(chOutArtifactName, chOutArtifactPath) // set property
	// go ArtifactGet(chOutArtifactPath, chOutArtifactGet)   // get artifact
	// go FileGuessType(chOutArtifactGet, chOutFileGuessType)
	// go FileSave(chOutFileGuessType, chOutFileSave)
	// go FileUntgz(chOutFileSave, chOutFileUnTgz)
	// go FileMove(chOutFileUnTgz, chOutFileMove)  // move file to final destination
	// go BuildPath(chOutFileMove, chOutBuildPath) // build $PATH

	// final sequential step. collects all instances in the pipeline and build a sumary
	err := lastStep(chOutLast)
	if err != nil {
		return "", err
	}
	// success
	return "", nil
}
