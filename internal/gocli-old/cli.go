/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package gocli

import (
	"fmt"

	"github.com/abtransitionit/luc/pkg/logx"
)

const CliDescription = "produces a CLI name and sends it into a channel."

func cli(arg ...string) (string, error) {
	// check argmuents
	if len(arg) == 0 {
		logx.L.Error("no CLI name provided")
		return "", fmt.Errorf("no CLI name provided")
	}

	cliName := arg[0]
	// logx.L.Info("Received CLI name", zap.String("cli", cliName))
	logx.L.Infof("Received CLI name: %s", cliName)

	// Step 1: Create a channel to send the CLI name
	cliNameChan := make(chan string)

	// define a goroutine
	go func() {
		defer close(cliNameChan)
		// logx.L.Debug("Sending CLI name to channel", zap.String("cli", cliName))
		logx.L.Debugf("Sending CLI name to channel")
		cliNameChan <- cliName
	}()

	// Step 2 will go here (GenericUrl)

	// TEMP: Just read from the channel to confirm it's working
	for name := range cliNameChan {
		// logx.L.Info("Read from channel", zap.String("cli", name))
		logx.L.Info("Read from channel: %s", name)
		return fmt.Sprintf("Received CLI name: %s", name), nil
	}

	return "", fmt.Errorf("nothing received")

}
