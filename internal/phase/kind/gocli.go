/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package kind

import (
	"github.com/abtransitionit/luc/internal/pipeline/gocli"
	"github.com/abtransitionit/luc/pkg/logx"
)

const GoCliDescription = "provision needed go CLI"

func goCli(arg ...string) (string, error) {
	logx.L.Info(GoCliDescription)
	// Launch the pipeline attach to this phase
	gocli.RunPipeline("kind", "nerdctl", "containerd", "rootlesskit", "slirp4netns")
	// goclipip.RunPipeline("toto", "kind")
	// goclipip.RunPipeline("kind")
	// goclipip.RunPipeline("toto", "kind", "nerdctl", "containerd", "rootlesskit", "slirp4netns")
	// on SUCCESS
	return "", nil
}
