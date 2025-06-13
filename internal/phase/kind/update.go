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
	// launch the pipeline that manage an OS update
	upd.RunPipeline()
	// on SUCCESS
	return "", nil
}
