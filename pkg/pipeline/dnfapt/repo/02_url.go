/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package repo

import (
	"strings"

	"github.com/abtransitionit/luc/pkg/logx"
)

func setUrl(in <-chan PipelineData, out chan<- PipelineData) {
	defer close(out)

	for data := range in {

		vm := data.HostName
		repoName := data.Config.Name

		if data.Err != nil {
			out <- data
			logx.L.Debugf("❌ Previous error detected")
			continue
		}

		// define placeholders based on OS:family
		var repoType, gpg string
		switch strings.TrimSpace(strings.ToLower(data.OsFamily)) {
		case "rhel", "fedora":
			repoType = "rpm"
			gpg = "repodata/repomd.xml.key"
		case "debian":
			repoType = "deb"
			gpg = "Release.key"
		default:
			repoType = "generic"
			gpg = "generic"
			logx.L.Warnf("[%s] [%s] unknown OS family, using generic placeholders", vm, repoName)
		}

		// define specific repo url
		urlRepo := data.GenericUrlRepo // e.g. https://pkgs.k8s.io/core:/stable:/$TAG/$PACK/
		urlRepo = strings.ReplaceAll(urlRepo, "$TAG", data.Version)
		urlRepo = strings.ReplaceAll(urlRepo, "$PACK", repoType)

		// define specific gpg url
		urlGpg := data.GenericUrlGpg // e.g. https://pkgs.k8s.io/core:/stable:/$TAG/$PACK/$GPG
		urlGpg = strings.ReplaceAll(urlGpg, "$TAG", data.Version)
		urlGpg = strings.ReplaceAll(urlGpg, "$PACK", repoType)
		urlGpg = strings.ReplaceAll(urlGpg, "$GPG", gpg)

		// set instance property
		data.UrlRepo = urlRepo
		data.UrlGpg = urlGpg

		// log
		logx.L.Debugf("[%s] [%s] replaced placeholders for URLs repo, and Url gpg", vm, repoName)
		out <- data
	}
}

// GenericRepoFilePath = $RepoFolder + RepoName + $RepoExt

// "rhel", "fedora"
// - RepoFolder = /etc/yum.repos.d
// - RepoExt = .repo

// "debian"
// - RepoFolder = /etc/apt/sources.list.d
// - RepoExt = .list
