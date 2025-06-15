/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package kind

import (
	"github.com/abtransitionit/luc/internal/pipeline/test"
	"github.com/abtransitionit/luc/pkg/logx"
)

const CliDescription = "provision needed CLI"

func cli(arg ...string) (string, error) {
	logx.L.Info(CliDescription)
	// Launch the pipeline attach to this phase
	test.RunPipeline()
	// on SUCCESS
	return "", nil
}
