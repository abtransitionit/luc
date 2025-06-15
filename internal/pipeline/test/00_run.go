/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package test

import (
	"github.com/abtransitionit/luc/pkg/logx"
)

const RunPipelineDescription = "Test run a pipeline."

var testList = []string{"kind", "containerd", "nerdctl", "rootlesskit", "slirp4netns"}

func RunPipeline() (string, error) {
	logx.L.Debug(RunPipelineDescription)

	chOutSource := make(chan string)

	Source(chOutSource, testList) // goroutine
	out := Stage1(chOutSource)    // goroutine
	out = Stage2(out)             // goroutine
	End(out)

	// on SUCCESS
	return "", nil
}
