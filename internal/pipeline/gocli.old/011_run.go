/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package gocliold

import (
	"fmt"

	"github.com/abtransitionit/luc/pkg/logx"
)

const Ep11Description = "produces a CLI name and sends it into a channel."

func RunPipeline3(cliName string) (string, error) {
	// check argmuents
	if cliName == "" {
		logx.L.Error("no CLI name provided")
		return "", fmt.Errorf("no CLI name provided")
	}

	// Define the pipeline channels
	chCliName := make(chan PipelineData)
	chGenericUrl := make(chan PipelineData)
	chSpecificUrl := make(chan PipelineData)
	chArtifactName := make(chan PipelineData)
	chArtifactPath := make(chan PipelineData)
	chGetArtifact := make(chan PipelineData)
	chGuessFileType := make(chan PipelineData)
	chSaveFile := make(chan PipelineData)
	chUnTgzFile := make(chan PipelineData)
	chMvFof := make(chan PipelineData)
	chUpdatePath := make(chan PipelineData)
	chEndPipeline := make(chan PipelineData)

	// Start each pipeline step
	CliName(chCliName, cliName)
	GenericUrl(chCliName, chGenericUrl)
	SpecificUrl(chGenericUrl, chSpecificUrl)
	ArtifactName(chSpecificUrl, chArtifactName)
	ArtifactPath(chArtifactName, chArtifactPath)
	GetArtifact(chArtifactPath, chGetArtifact)
	GuessFileType(chGetArtifact, chGuessFileType)
	SaveFile(chGuessFileType, chSaveFile)
	UnTgzFile(chSaveFile, chUnTgzFile)
	MvFof(chUnTgzFile, chMvFof)
	UpdatePath(chMvFof, chUpdatePath)
	EndPipeline(chUpdatePath, chEndPipeline)

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
