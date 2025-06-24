/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
// internal/kbe/phases.go
package kbe

import "github.com/abtransitionit/luc/pkg/phase"

var Phases = []phase.Phase{
	phase.SetPhase("display", display, DisplayDescription),
	phase.SetPhase("checkssh", checkSsh, CheckSshDescription),
	phase.SetPhase("cpluc", cpLuc, CpLucDescription),

	// phase.SetPhase("update", update, UpdateDescription),
	// phase.SetPhase("reboot", reboot, RebootDescription),
}
