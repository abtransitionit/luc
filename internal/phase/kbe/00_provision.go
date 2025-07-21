/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
// internal/kbe/phases.go
package kbe

import (
	"github.com/abtransitionit/luc/pkg/phase"
	"github.com/abtransitionit/luc/pkg/util"
)

var ProvisionPhases = []phase.Phase{
	phase.SetPhase("show", show, DisplayDescription),
	phase.SetPhase("checkssh", checkSsh, util.CheckSshDescription),
	phase.SetPhase("cpluc", cpLuc, CpLucDescription),
	phase.SetPhase("upgrade", upgrade, UpgradeDescription),
	phase.SetPhase("dapack1", daPackStd, DaPackStdDescription),
	phase.SetPhase("darepo", daRepo, DaRepoDescription),
	// phase.SetPhase("dapack", daPack, DaPackDescription),
	// phase.SetPhase("gocli", goCli, GoCliDescription),

	// phase.SetPhase("update", update, UpdateDescription),
	// phase.SetPhase("reboot", reboot, RebootDescription),
}
