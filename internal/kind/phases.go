/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
// internal/kbe/phases.go
package kind

import "github.com/abtransitionit/luc/pkg/deploy"

var Phases = []deploy.Phase{
	deploy.SetPhase("ssh", ssh, SshDescription),
	deploy.SetPhase("cpluc", cpluc, CpLucDescription),
	deploy.SetPhase("check", check, CheckDescription),
	deploy.SetPhase("update", update, UpdateDescription),
	deploy.SetPhase("cli", cli, CliDescription),
	deploy.SetPhase("service", service, ServiceDescription),
}
