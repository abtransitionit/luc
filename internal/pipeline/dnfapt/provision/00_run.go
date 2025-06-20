/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package provision

import (
	"github.com/abtransitionit/luc/pkg/logx"
)

const RunPipelineDescription = "provision GO CLI(s)."

func RunPipeline(packageNameList ...string) (string, error) {
	logx.L.Debug(RunPipelineDescription)

	// Count and log the number of CLI args
	argCount := len(packageNameList)
	logx.L.Debugf("Received %d CLI(s) to provisioned:  %v", argCount, packageNameList)

	// Define the pipeline channels
	chOutSource := make(chan PipelineData)
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
	chOutLast := chOutSource
	// chOutLast := make(chan PipelineData)

	// Start each pipeline stage concurently
	go source(chOutSource, packageNameList...) // boostrap the Data
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

	// This is the not a stage but the last foreground step reading all instance in the pipeline
	err := lastStep(chOutLast)
	if err != nil {
		return "", err
	}
	// on SUCCESS
	// time.Sleep(10 * time.Second)
	return "", nil
}
