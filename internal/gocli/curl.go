/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package gocli

import (
	"errors"

	"github.com/abtransitionit/luc/pkg/config"
	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/abtransitionit/luc/pkg/util"
)

const CurlDescription = "curl the artifact from a CLI in the ConfigMap."

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
	cli, ok := config.GetCLIConfig(cliName)
	if !ok {
		msg := "CLI not found in the CliConfigMap"
		logx.L.Debugf("❌ %s. impacted CLI: %s", msg, cliName)
		return "", errors.New(msg)
	}

	// is the URL Curlable ?
	yes, err := cli.UrlType.IsCurlable()
	if err != nil || !yes {
		msg := "the URL type of the CLI is not curlable"
		logx.L.Debugf("❌ %s. impacted CLI: %s", msg, cliName)
		return "", errors.New(msg)
	}

	// Get the URL
	cliUrl, _ := config.GetCliUrl(logx.L, cliName, "linux", "amd64") // OS and Arch auto-detected

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
		logx.L.Debug("the curled content in memory is guessed as a gzipped file")

	case isExe:
		logx.L.Debug("the curled content in memory is guessed as an executable")
		return "", errors.New("the curled content in memory is not a gzipped file")

	default:
		logx.L.Debug("the curled content in memory is not guessed as a gzipped file, nor an executable")
		logx.L.Debugf("the curled content in memory is not guessed as an executable")
		return "", errors.New("the curled content in memory is not a gzipped file")
	}

	yes, err = util.IsGzippedMemoryContent(fileInMemory)
	if err != nil || !yes {
		logx.L.Debugf("the curled content in memory is not guessed as a gzipped file")

		ok, err := util.IsMemoryContentAnExe(fileInMemory)
		if !ok {
			logx.L.Debugf("the curled content in memory is not guessed as an executable: %v", err)
		} else {
			logx.L.Debug("the curled content in memory is guessed as an executable")
		}
		return "", errors.New("the curled content in memory is not guessed as a gzipped file")
	}

	logx.L.Debug("the curled content in memory is guessed as a gzipped file")

	// ls contents
	err = util.ListTgzContentInMemory(fileInMemory)
	if err != nil {
		return "", err
	}

	// // save file to file
	// _, err = util.SaveToFile(logx.L, "/tmp/toto", fileInMemory)
	// if err != nil {
	// 	return "", err
	// }

	// handle applogic SUCCESS
	return "", nil
}

// location := "/usr/local/bin/luc"
// version := "0.0.1"
// doc := "https://github.com/abtransitionit/luc"
// git := "https://github.com/abtransitionit/luc"
// fmt.Printf("🔹 CLI is available at %s (version: %s)\n", location, version)
// fmt.Printf("🔹 Visit the official docs: %s\n", doc)
// fmt.Printf("🔹 Visit the official git: %s\n", git)
// Replace placeholders in URL
// url := cliConf.Url
// if strings.Contains(url, "$") {
// 	url = strings.ReplaceAll(url, "$TAG", cliConf.Tag)
// 	url = strings.ReplaceAll(url, "$NAME", cliConf.Name)
// 	url = strings.ReplaceAll(url, "$OS", osType)
// 	url = strings.ReplaceAll(url, "$ARCH", osArch)
// }
