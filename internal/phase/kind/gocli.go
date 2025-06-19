/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package kind

import (
	goclipip "github.com/abtransitionit/luc/internal/pipeline/gocli"
	"github.com/abtransitionit/luc/pkg/logx"
)

const GocliDescription = "provision needed go CLI"

func gocli(arg ...string) (string, error) {
	logx.L.Info(GocliDescription)
	// Launch the pipeline attach to this phase
	// goclipip.RunPipeline("toto", "kind", "nerdctl", "containerd", "rootlesskit", "slirp4netns")
	// goclipip.RunPipeline("toto", "kind")
	goclipip.RunPipeline("kind")
	// goclipip.RunPipeline("toto", "kind", "nerdctl", "containerd", "rootlesskit", "slirp4netns")
	// on SUCCESS
	return "", nil
}
