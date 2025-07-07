/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package kbe

import (
	"github.com/abtransitionit/luc/pkg/logx"
	rebootpip "github.com/abtransitionit/luc/pkg/pipeline/reboot"
)

const RebootDescription = "Reboot an OS."

func reboot(arg ...string) (string, error) {
	logx.L.Info(RebootDescription)
	rebootpip.RunPipeline()
	// on SUCCESS
	return "", nil
}
