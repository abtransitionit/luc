/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package kind

import (
	upd "github.com/abtransitionit/luc/internal/pipeline/dnfapt/update"
	"github.com/abtransitionit/luc/pkg/logx"
)

const UpdateDescription = "upgrade The Kind VM OS packages and packages repositories to version latest."

func update(arg ...string) (string, error) {
	logx.L.Info(UpdateDescription)
	logx.L.Debug("Launching pipeline that manage OS update")
	// Launch the pipeline
	upd.RunPipeline()
	// on SUCCESS
	return "", nil
}
