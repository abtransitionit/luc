/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package gocli

import (
	"fmt"
	"os"

	"github.com/abtransitionit/luc/pkg/config"
	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/abtransitionit/luc/pkg/util"
)

func getArtifact(in <-chan PipelineData, out chan<- PipelineData) {
	defer close(out)

	for data := range in {
		if data.Err != nil {
			out <- data
			logx.L.Debugf("❌ Previous error detected")
			continue
		}

		// declaradef variable
		var processedData PipelineData

		// handle artifact download method (based on UrlType)
		logx.L.Debugf("[%s] donwloading atrifact", data.Config.Name)

		switch data.Config.UrlType {
		case config.UrlExe, config.UrlTgz:
			processedData = handleCurl(data)

		case config.UrlGit:
			processedData = handleGit(data)

		default:
			data.Err = fmt.Errorf("[%s] ❌ Unsupported or not yet managed UrlType :  '%s'", data.CliName, data.Config.UrlType)
			logx.L.Debugf("[%s] ❌ Error detected 1", data.CliName)
			out <- data
			continue
		}

		// log information
		logx.L.Debugf("[%s] donwloaded atrifact", data.Config.Name)
		// send
		out <- processedData
	}
}

// curl the File into memory
func handleCurl(data PipelineData) PipelineData {
	if data.AppUrl == "" {
		data.Err = fmt.Errorf("[%s] ❌ AppUrl is empty for curl", data.Config.Name)
		logx.L.Debugf("[%s] ❌ Error detected 2", data.Config.Name)
		return data
	}

	memoryFile, err := util.GetPublicFile(data.AppUrl)
	if err != nil {
		data.Err = err
		logx.L.Debugf("[%s] ❌ Error detected 3", data.Config.Name)
		return data
	}

	// set this instance property
	data.MemoryFile = memoryFile
	data.FofTmpPath = data.ArtifactPath
	// log information
	logx.L.Infof("[%s] File Downloaded into Memory", data.Config.Name)

	return data
}

// git clone the repo
func handleGit(data PipelineData) PipelineData {
	if data.AppUrl == "" {
		data.Err = fmt.Errorf("[%s] ❌ git URL is empty", data.Config.Name)
		logx.L.Debugf("[%s] ❌ Error detected 4", data.Config.Name)
		return data
	}
	// log before action
	logx.L.Debugf("[%s] git cloning artifact'%s'", data.Config.Name, data.AppUrl)

	// Create a temporary directory for the git clone
	tmpDir, err := os.MkdirTemp("/tmp", data.Config.Name+"-git-*")
	if err != nil {
		data.Err = fmt.Errorf("[%s] ❌ failed to create temp directory: %w", data.Config.Name, err)
		logx.L.Debugf("[%s] ❌ Error detected 5", data.Config.Name)
		return data
	}

	// set this instance property
	data.FofTmpPath = tmpDir
	// fmt.Println(data.String())

	// play CLI
	cli := fmt.Sprintf("git clone --branch %s --depth 1 %s %s", data.Config.Tag, data.AppUrl, tmpDir)
	_, err = util.RunCLILocal(cli)
	if err != nil {
		data.Err = fmt.Errorf("❌ git clone failed: %w", err)
		logx.L.Debugf("❌ Error detected 6")
		return data
	}

	// log after action
	logx.L.Infof("✅ Git repository successfully cloned to: %s", data.FofTmpPath)

	return data
}

// cmd := exec.Command("git", "clone", data.Config.Url, tmpDir)
// cmd.Stdout = os.Stdout
// cmd.Stderr = os.Stderr
