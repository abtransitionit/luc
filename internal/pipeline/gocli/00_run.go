/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package gocli

import (
	"github.com/abtransitionit/luc/pkg/logx"
)

const RunPipelineDescription = "provision GO CLIs."

func RunPipeline(cli ...string) (string, error) {
	logx.L.Debug(RunPipelineDescription)

	// Count and log the number of CLI args
	argCount := len(cli)
	logx.L.Debugf("Received %d CLI(s) to provisioned:  %v", argCount, cli)

	// Define the pipeline channels
	chOutSource := make(chan PipelineData)
	chOutGenericUrl := make(chan PipelineData)
	chOutSpecificUrl := make(chan PipelineData)
	chOutArtifactName := make(chan PipelineData)
	chOutArtifactPath := make(chan PipelineData)
	chOutArtifact := make(chan PipelineData)
	chOutFileGuessType := make(chan PipelineData)
	chOutFileSave := make(chan PipelineData)
	chOutLast := chOutFileSave
	// chOutLast := make(chan PipelineData)

	// Start each pipeline stage concurently
	go source(chOutSource, cli...)                        // boostrap the Data
	go GenericUrl(chOutSource, chOutGenericUrl)           // set property
	go SpecificUrl(chOutGenericUrl, chOutSpecificUrl)     // set property
	go ArtifactName(chOutSpecificUrl, chOutArtifactName)  // set property
	go ArtifactPath(chOutArtifactName, chOutArtifactPath) // set property
	go GetArtifact(chOutArtifactPath, chOutArtifact)      // get artifact
	go FileGuessType(chOutArtifact, chOutFileGuessType)   // get artifact
	go FileSave(chOutFileGuessType, chOutFileSave)        // get artifact

	// This is the not a stage but the last foreground step reading all instance in the pipeline
	err := lastStep(chOutLast)
	if err != nil {
		return "", err
	}
	// on SUCCESS
	// time.Sleep(10 * time.Second)
	return "", nil
}
