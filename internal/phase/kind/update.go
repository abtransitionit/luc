/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package kind

import (
	updatepip "github.com/abtransitionit/luc/internal/pipeline/dnfapt/update"
	"github.com/abtransitionit/luc/pkg/logx"
)

const UpdateDescription = "upgrade Kind VM OS using pipeleine."

func update(arg ...string) (string, error) {
	logx.L.Info(UpdateDescription)
	updatepip.RunPipeline()
	// on SUCCESS
	return "", nil
}
