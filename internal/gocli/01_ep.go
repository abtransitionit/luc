/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package gocli

import (
	"fmt"

	"github.com/abtransitionit/luc/pkg/logx"
)

const EpDescription = "produces a CLI name and sends it into a channel."

func ep(arg ...string) (string, error) {
	// check argmuents
	if len(arg) == 0 {
		logx.L.Error("no CLI name provided")
		return "", fmt.Errorf("no CLI name provided")
	}
	// Get the CLI name
	cliName := arg[0]
	// cliName := "toto"
	logx.L.Infof("Received CLI name: %s", cliName)

	// Define the pipeline channels
	chCliName := make(chan PipelineData)
	chGenericUrl := make(chan PipelineData)
	chSpecificUrl := make(chan PipelineData)
	chArtifactName := make(chan PipelineData)
	chArtifactPath := make(chan PipelineData)
	chCurlUrl := make(chan PipelineData)
	chGuessFileType := make(chan PipelineData)
	chSaveFile := make(chan PipelineData)

	// Start each pipeline step
	CliName(chCliName, cliName)
	GenericUrl(chCliName, chGenericUrl)
	SpecificUrl(chGenericUrl, chSpecificUrl)
	ArtifactName(chSpecificUrl, chArtifactName)
	ArtifactPath(chArtifactName, chArtifactPath)
	CurlUrl(chArtifactPath, chCurlUrl)
	GuessFileType(chCurlUrl, chGuessFileType)
	SaveFile(chGuessFileType, chSaveFile)

	// Read final result
	for data := range chSaveFile {
		if data.Err != nil {
			logx.L.Debugf("❌ Error: %s", data.Err)
		}
	}
	return "", fmt.Errorf("nothing received")

}
