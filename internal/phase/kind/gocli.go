/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package kind

import (
	"github.com/abtransitionit/luc/internal/config"
	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/abtransitionit/luc/pkg/pipeline/gocli"
)

const GoCliDescription = "provision needed go CLI"

func goCli(arg ...string) (string, error) {
	logx.L.Info(GoCliDescription)
	// Launch the pipeline attach to this phase
	// gocli.RunPipeline("kind", "nerdctl", "containerd", "rootlesskit", "slirp4netns")
	gocli.RunPipeline(config.KindCliConfigMap)
	// on SUCCESS
	return "", nil
}
