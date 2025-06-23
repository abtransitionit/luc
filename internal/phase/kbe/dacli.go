/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package kbe

import (
	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/abtransitionit/luc/pkg/pipeline/dnfapt/provision"
)

const DaCliDescription = "provision needed Dnfapt CLI"

func daCli(arg ...string) (string, error) {
	logx.L.Info(DaCliDescription)
	// Launch the pipeline attach to this phase
	provision.RunPipeline("uidmap")
	// on SUCCESS
	return "", nil
}
