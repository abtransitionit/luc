/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package repo

import (
	"fmt"
	"strings"

	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/abtransitionit/luc/pkg/util"
)

func saveRepoFile(in <-chan PipelineData, out chan<- PipelineData) {
	defer close(out)

	for data := range in {

		vm := data.HostName
		repoName := data.Config.Name

		if data.Err != nil {
			out <- data
			logx.L.Debugf("❌ Previous error detected")
			continue
		}

		// define file content
		logx.L.Debugf("[%s] [%s] defining repo file content", vm, repoName)
		var stringContent string
		switch strings.TrimSpace(strings.ToLower(data.OsFamily)) {
		case "rhel", "fedora":
			stringContent = helperDefineRepoFileContentRhel(data)
		case "debian":
			stringContent = helperDefineRepoFileContentDebian(data)
		default:
			logx.L.Warnf("[%s] [%s] unknown OS family, using generic placeholders", vm, repoName)
		}
		logx.L.Debugf("[%s] [%s] defined repo file content", vm, repoName)

		// remote create file:repo
		logx.L.Debugf("[%s] [%s] creating repo file", vm, repoName)
		if err := helperSaveRepofile(vm, stringContent, data.RepoFilePath); err != nil {
			data.Err = err
			logx.L.Debugf("❌ [%s][%s] Error saving repo file: %v", vm, repoName, err)
			out <- data
			continue
		}
		logx.L.Debugf("[%s] [%s] created repo file", vm, repoName)
		// send
		out <- data
	}
}

func helperDefineRepoFileContentRhel(data PipelineData) string {
	content := fmt.Sprintf(`
		[%s]
		name=%s
		enabled=1
		gpgcheck=1
		baseurl=%s
		gpgkey=%s
	`, data.CName, data.CName, data.UrlRepo, data.UrlGpg)

	return content
}

func helperDefineRepoFileContentDebian(data PipelineData) string {
	content := fmt.Sprintf(`
		deb [signed-by=%s] %s /
	`, data.GpgFilePath, data.UrlRepo)

	return content
}

func helperSaveRepofile(vm string, stringContent string, FilePath string) error {
	cli := fmt.Sprintf(`luc do SaveStringToFile '%s' %s true`, stringContent, FilePath)
	outp, err := util.RunCLIRemote(vm, cli)
	if err != nil {
		return fmt.Errorf("%v, %s", err, outp)
	}
	return nil
}
