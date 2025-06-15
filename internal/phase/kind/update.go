/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package kind

import (
	upd "github.com/abtransitionit/luc/internal/pipeline/dnfapt/update"
	"github.com/abtransitionit/luc/pkg/logx"
)

const UpdateDescription = "upgrade Kind VM OS."

func update(arg ...string) (string, error) {
	logx.L.Info(UpdateDescription)
	// Launch the pipeline attach to this phase
	upd.RunPipeline()
	// on SUCCESS
	return "", nil
}
