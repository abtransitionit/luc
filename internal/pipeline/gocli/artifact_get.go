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

func ArtifactGet(in <-chan PipelineData, out chan<- PipelineData) {
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
		switch data.Config.UrlType {
		case config.UrlExe, config.UrlTgz:
			processedData = helperCurl(data)

		case config.UrlGit:
			processedData = helperGitClone(data)

		default:
			data.Err = fmt.Errorf("❌ Unsupported or not yet managed UrlType :  '%s'", data.Config.UrlType)
			logx.L.Debugf("❌ Error detected 1")
			out <- data
			continue
		}

		// send
		out <- processedData
	}
}

// curl the File into memory
func helperCurl(data PipelineData) PipelineData {
	if data.SpecificUrl == "" {
		data.Err = fmt.Errorf("❌ SpecificUrl is empty for curl")
		logx.L.Debugf("❌ Error detected 2")
		return data
	}

	memoryFile, err := util.GetPublicFile(data.SpecificUrl)
	if err != nil {
		data.Err = err
		logx.L.Debugf("❌ Error detected 3")
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
func helperGitClone(data PipelineData) PipelineData {
	if data.SpecificUrl == "" {
		data.Err = fmt.Errorf("❌ git URL is empty")
		logx.L.Debugf("❌Error detected 4")
		return data
	}
	// log before action
	logx.L.Debugf("[%s] git cloning artifact'%s'", data.Config.Name, data.SpecificUrl)

	// Create a temporary directory for the git clone
	tmpDir, err := os.MkdirTemp("/tmp", data.Config.Name+"-git-*")
	if err != nil {
		data.Err = fmt.Errorf("❌ failed to create temp directory: %w", err)
		logx.L.Debugf("❌ Error detected 5")
		return data
	}

	// set this instance property
	data.FofTmpPath = tmpDir
	// fmt.Println(data.String())

	// play CLI
	cli := fmt.Sprintf("git clone --branch %s --depth 1 %s %s", data.Config.Tag, data.SpecificUrl, tmpDir)
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
