/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package gocli

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/abtransitionit/luc/pkg/config"
	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/abtransitionit/luc/pkg/util"
)

func GetArtifact(in <-chan PipelineData, out chan<- PipelineData) {
	go func() {
		// close channel
		defer close(out)

		for data := range in {
			// Step 1: propagate error if any
			if data.Err != nil {
				out <- data
				continue
			}
			// declaradef variable
			var processedData PipelineData

			// Step 2: handle artifact download method (based on UrlType)
			switch data.Config.UrlType {
			case config.UrlExe, config.UrlTgz:
				processedData = helperCurl(data)

			case config.UrlGit:
				// log information
				logx.L.Debugf("UrlType '%s' not need to be curled but git cloned", data.Config.UrlType)
				processedData = helperGitClone(data)

			default:
				// log information
				logx.L.Debug("UrlType '%s' not need to be curled nor git cloned", data.Config.UrlType)
				// set Error to propagate
				data.Err = fmt.Errorf("UrlType '%s' not need to be curled nor git cloned", data.Config.UrlType)
				// send data to next step
				out <- data
				// Keep reading data from channel
				continue
			}

			// Step 3: send result to next pipeline step
			out <- processedData
		}
	}()
}

// curl the File into memory
func helperCurl(data PipelineData) PipelineData {
	if data.SpecificUrl == "" {
		data.Err = fmt.Errorf("SpecificUrl is empty for curl")
		return data
	}

	memoryFile, err := util.GetPublicFile(logx.L, data.SpecificUrl)
	if err != nil {
		// data.Err = fmt.Errorf("failed to download file: %w", err)
		return data
	}

	// define this property
	data.MemoryFile = memoryFile

	// log information
	logx.L.Infof("File Downloaded into Memory")

	return data
}

func helperGitClone(data PipelineData) PipelineData {
	if data.Config.Url == "" {
		data.Err = fmt.Errorf("git URL is empty")
		return data
	}
	// log before action
	logx.L.Debugf("git cloning artifact from '%s'", data.Config.Url)

	// Create a temporary directory for the git clone
	tmpDir, err := os.MkdirTemp("", data.Config.Name+"-git-*")
	if err != nil {
		data.Err = fmt.Errorf("failed to create temp directory: %w", err)
		return data
	}

	// define property
	data.FofTmpPath = tmpDir

	// build arguments for the CLI
	args := []string{"clone", "--branch", data.Config.Tag, "--depth", "1", data.Config.Url, tmpDir}

	// play the CLI
	cmd := exec.Command("git", args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		data.Err = fmt.Errorf("git clone failed: %w\nOutput:\n%s", err, string(output))
		logx.L.Debugf("❌ Git clone failed: %v", err)
		return data
	}

	// log after action
	logx.L.Infof("✅ Git repository successfully cloned to: %s", data.FofTmpPath)

	return data
}

// cmd := exec.Command("git", "clone", data.Config.Url, tmpDir)
// cmd.Stdout = os.Stdout
// cmd.Stderr = os.Stderr
