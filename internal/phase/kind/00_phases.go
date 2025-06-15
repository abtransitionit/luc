/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
// internal/kbe/phases.go
package kind

import "github.com/abtransitionit/luc/pkg/phase"

var Phases = []phase.Phase{
	phase.SetPhase("update", update, UpdateDescription),
	phase.SetPhase("cli", cli, CliDescription),
	phase.SetPhase("service", service, ServiceDescription),
	phase.SetPhase("env", env, EnvDescription),
}
