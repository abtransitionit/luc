/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package gocli

import (
	"errors"
	"fmt"
	"path"
	"time"

	"github.com/abtransitionit/luc/pkg/config"
	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/abtransitionit/luc/pkg/util"
)

const CurlDescription = "curl the artifact from a CLI in the ConfigMap into memory."

// # Purpose
//
// curl a file in a temporary folder
//
// # Parameters
//
// - arg[0] : the file to curl
//
// Returns:
//   - bool
//   - error
//
// Possible returns:
//
//   - (true, nil)   if the file is successfully downloaded
//   - (false, nil)  if the file is not successfully downloaded
//   - (true, error) if the file fails to download
func curl(arg ...string) (string, error) {

	// check arguments
	if len(arg) == 0 {
		logx.L.Debugf("❌ No argument provided", arg)
		return "", nil
	}

	// Get impacted CLI
	cliName := arg[0]

	// Print message
	logx.L.Infof("'%s' for CLI '%s'", CurlDescription, cliName)

	// Does config exist ?
	cli, ok := config.GetCLIConfigMap(cliName)
	if !ok {
		msg := "CLI not found in the CliConfigMap"
		logx.L.Debugf("❌ %s. impacted CLI: %s", msg, cliName)
		return "", errors.New(msg)
	}

	// is the URL Curlable ?
	yes, err := cli.UrlType.IsCurlable()
	if err != nil || !yes {
		msg := "the URL type of the CLI is not curlable"
		logx.L.Debugf("❌ %s", msg)
		return "", errors.New(msg)
	}

	// Get the URL
	cliUrl, _ := config.GetCliSpecificUrl(logx.L, cliName) // OS and Arch auto-detected
	artefactName := path.Base(cliUrl)

	// Print message
	logx.L.Infof("The artefact of CLI '%s' is CURLable and its URL is '%s'", cliName, cliUrl)

	// curl the artifact into memory
	fileInMemory, err := util.GetPublicFile(logx.L, cliUrl)
	if err != nil {
		return "", err
	}

	// Guess the curl content
	isGzip, _ := util.IsGzippedMemoryContent(fileInMemory)
	isExe, _ := util.IsMemoryContentAnExe(fileInMemory)

	switch {
	case isGzip:
		logx.L.Debugf("the curled content in memory is guessed as a gzipped file and will be extracted to '%s'", artefactName)

	case isExe:
		msg := "the curled content in memory is guessed as an executable"
		logx.L.Debugf("'%s' and will be extracted to %s", msg, artefactName)

	default:
		msg := "the curled content in memory is not guessed as a gzipped file, nor an executable"
		logx.L.Debug(msg)
		return "", errors.New(msg)
	}

	// ls contents
	_ = util.ListTgzContentInMemory(fileInMemory)

	// save memory content into OS file
	uniqueFsPath := fmt.Sprintf("/tmp/%s_%d", artefactName, time.Now().UnixNano())
	_, err = util.SaveToFile(logx.L, uniqueFsPath, fileInMemory)
	if err != nil {
		return "", err
	}

	// handle applogic SUCCESS
	dcp(uniqueFsPath)
	return "", nil
}
