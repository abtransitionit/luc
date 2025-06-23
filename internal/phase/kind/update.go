/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package kind

import (
	"github.com/abtransitionit/luc/pkg/logx"
	updatepip "github.com/abtransitionit/luc/pkg/pipeline/dnfapt/update"
)

const UpdateDescription = "upgrade Kind VM OS using pipeleine."

func update(arg ...string) (string, error) {
	logx.L.Info(UpdateDescription)
	updatepip.RunPipeline()
	// on SUCCESS
	return "", nil
}
