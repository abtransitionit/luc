/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package kind

import (
	rebootpip "github.com/abtransitionit/luc/internal/pipeline/reboot"
	"github.com/abtransitionit/luc/pkg/logx"
)

const RebootDescription = "Reboot an OS."

func reboot(arg ...string) (string, error) {
	logx.L.Info(RebootDescription)
	rebootpip.RunPipeline()
	// on SUCCESS
	return "", nil
}
