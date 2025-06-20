/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package provision

import (
	"github.com/abtransitionit/luc/pkg/logx"
)

const RunPipelineDescription = "provision Dnfapt CLI(s)."

func RunPipeline(packageNameList ...string) (string, error) {
	logx.L.Debug(RunPipelineDescription)

	// Count and log the number of CLI args
	argCount := len(packageNameList)
	logx.L.Debugf("Received %d CLI(s) to provisioned:  %v", argCount, packageNameList)

	// Define the pipeline channels
	chOutSource := make(chan PipelineData)
	chOutProvision := make(chan PipelineData)
	chOutLast := chOutProvision

	// Start each pipeline stage concurently
	go source(chOutSource, packageNameList...) // boostrap the Data
	go provision(chOutSource, chOutProvision)  // provision

	// This is not a stage but the last foreground step reading all instance in the pipeline
	err := lastStep(chOutLast)
	if err != nil {
		return "", err
	}
	// on SUCCESS
	// time.Sleep(10 * time.Second)
	return "", nil
}
