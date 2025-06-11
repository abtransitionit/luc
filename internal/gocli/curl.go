/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package gocli

import (
	"github.com/abtransitionit/luc/pkg/logx"
)

const CurlDescription = "curl the artifact."

func curl(arg ...string) (string, error) {
	if len(arg) == 0 {
		logx.L.Debugf("❌ No argument provided", arg)
		// return errorx.StringError("No argument provided", arg, errors.New(""))
	}
	logx.L.Infof("%s : %s", CurlDescription, arg[0])

	logx.L.Infof(CurlDescription)
	logx.L.Infof("nb arg passed: '%d'", len(arg))
	// logx.L.Info("want to download the artifact of '%s'", arg[0])
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
