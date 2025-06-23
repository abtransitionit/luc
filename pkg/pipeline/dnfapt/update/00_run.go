/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package update

import (
	"github.com/abtransitionit/luc/pkg/logx"
)

const RunPipelineDescription = "Pipeline: update OS package and repositories to version latest."

func RunPipeline() (string, error) {
	logx.L.Debug(RunPipelineDescription)

	// Define the pipeline channels
	chOutSource := make(chan PipelineData)
	chOutBefore := make(chan PipelineData)
	chOutUpdate := make(chan PipelineData)
	chOutAfter := make(chan PipelineData)
	chOutLast := chOutAfter

	// Start each pipeline stage concurently
	go source(chOutSource)                  // boostrap the Data
	go infoBefore(chOutSource, chOutBefore) // set property
	go update(chOutBefore, chOutUpdate)     // update the OS
	go infoAfter(chOutUpdate, chOutAfter)   // set property

	// This is the not a stage but the last foreground process waiting for the last stage data
	err := lastStep(chOutLast)
	if err != nil {
		return "", err
	}
	// on SUCCESS
	// time.Sleep(10 * time.Second)
	return "", nil
}
