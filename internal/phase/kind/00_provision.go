/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
// internal/kbe/phases.go
package kind

import (
	"github.com/abtransitionit/luc/pkg/phase"
	"github.com/abtransitionit/luc/pkg/util"
)

var ProvisionPhases = []phase.Phase{
	phase.SetPhase("show", show, DisplayDescription),
	phase.SetPhase("checkssh", checkSsh, util.CheckSshDescription),
	phase.SetPhase("cpluc", cpLuc, CpLucDescription),
	phase.SetPhase("upgrade", upgrade, UpgradeDescription),
	phase.SetPhase("dapack", daPack, DaPackDescription),
	phase.SetPhase("gocli", goCli, GoCliDescription),
	phase.SetPhase("service", service, ServiceDescription),
	phase.SetPhase("linger", linger, LingerDescription),
	phase.SetPhase("path", path, PathDescription),
	// phase.SetPhase("rc", rc, RcDescription),
	// phase.SetPhase("create", create, CreateDescription),
	// phase.SetPhase("check", check, CheckDescription),
}
