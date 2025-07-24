/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package repo

import (
	"path/filepath"
	"strings"

	"github.com/abtransitionit/luc/pkg/logx"
)

func setPath(in <-chan PipelineData, out chan<- PipelineData) {
	defer close(out)

	for data := range in {

		vm := data.HostName
		repoCName := data.CName
		repoName := data.Config.Name

		if data.Err != nil {
			out <- data
			logx.L.Debugf("❌ Previous error detected")
			continue
		}

		// define placeholders based on OS:family
		var repoExt, repoFolder, gpgFolder, gpgExt string
		// var repoFolder, repoExtension string
		switch strings.TrimSpace(strings.ToLower(data.OsFamily)) {
		case "rhel", "fedora":
			repoFolder = "/etc/yum.repos.d"
			repoExt = ".repo"
		case "debian":
			repoFolder = "/etc/apt/sources.list.d"
			gpgFolder = "/etc/apt/keyrings"
			repoExt = ".list"
			gpgExt = "-apt-keyring.gpg"
		default:
			repoFolder = "generic"
			repoExt = "generic"
			logx.L.Warnf("[%s] [%s] unknown OS family, using generic placeholders", vm, repoName)
		}

		// set instance property
		data.RepoFilePath = filepath.Join(repoFolder, repoCName+repoExt)

		switch strings.TrimSpace(strings.ToLower(data.OsFamily)) {
		case "debian":
			data.GpgFilePath = filepath.Join(gpgFolder, repoCName+gpgExt)
		}

		// data.RepoFilePath = filepath.Join(repoFolder, repoCName+repoExt)
		// data.GpgFilePath = filepath.Join(repoFolder, repoCName+"-apt-keyring.gpg")

		// success
		logx.L.Debugf("[%s] [%s] replaced placeholders for repo path, and gpg path", vm, repoName)
		out <- data
	}
}
