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
		logx.L.Debugf("❌ CLI '%s' not found in the ConfigMap", cliName)
		return "", errors.New("")
	}
	// is the URL Curlable ?
	curlable, err := cli.UrlType.IsCurlable()
	if err != nil || !curlable {
		logx.L.Debugf("❌ URL type '%s' is not curlable for CLI '%s'", cli.UrlType, cliName)
		return "", errors.New("URL type not curlable")
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

	// ls contents
	err = util.ListTgzInMemory(fileInMemory)
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
