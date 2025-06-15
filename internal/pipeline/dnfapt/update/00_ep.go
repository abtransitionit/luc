/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package update

import (
	"github.com/abtransitionit/luc/pkg/logx"
)

const RunPipelineDescription = "update OS package and repositories to version latest."

func RunPipeline() (string, error) {
	logx.L.Debug(RunPipelineDescription)
	// Define the pipeline channels
	chSource := make(chan PipelineData)
	// chGenericUrl := make(chan PipelineData)
	// chSpecificUrl := make(chan PipelineData)
	// chArtifactName := make(chan PipelineData)
	// chArtifactPath := make(chan PipelineData)
	// chGetArtifact := make(chan PipelineData)
	// chGuessFileType := make(chan PipelineData)
	// chSaveFile := make(chan PipelineData)
	// chUnTgzFile := make(chan PipelineData)
	// chMvFof := make(chan PipelineData)
	// chUpdatePath := make(chan PipelineData)
	chEndPipeline := make(chan PipelineData)

	// Start each pipeline step
	Source(chSource, "dnf")
	// GenericUrl(chCliName, chGenericUrl)
	// SpecificUrl(chGenericUrl, chSpecificUrl)
	// ArtifactName(chSpecificUrl, chArtifactName)
	// ArtifactPath(chArtifactName, chArtifactPath)
	// GetArtifact(chArtifactPath, chGetArtifact)
	// GuessFileType(chGetArtifact, chGuessFileType)
	// SaveFile(chGuessFileType, chSaveFile)
	// UnTgzFile(chSaveFile, chUnTgzFile)
	// MvFof(chUnTgzFile, chMvFof)
	// UpdatePath(chMvFof, chUpdatePath)
	// EndPipeline(chUpdatePath, chEndPipeline)
	EndPipeline(chSource, chEndPipeline)

	// Read final result
	for data := range chEndPipeline {
		if data.Err != nil {
			logx.L.Debugf("❌ Error: %s", data.Err)
			return "", data.Err
		}
	}
	// return "", fmt.Errorf("nothing received")
	return "", nil

}
