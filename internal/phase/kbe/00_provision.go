/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
// internal/kbe/phases.go
package kbe

import "github.com/abtransitionit/luc/pkg/phase"

var ProvisionPhases = []phase.Phase{
	phase.SetPhase("show", show, DisplayDescription),
	phase.SetPhase("checkssh", checkSsh, CheckSshDescription),
	phase.SetPhase("cpluc", cpLuc, CpLucDescription),
	phase.SetPhase("upgradeos", upgradeOs, UpgradeOsDescription),

	// phase.SetPhase("update", update, UpdateDescription),
	// phase.SetPhase("reboot", reboot, RebootDescription),
}
