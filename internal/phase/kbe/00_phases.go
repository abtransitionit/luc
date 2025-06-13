/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
// internal/kbe/phases.go
package kbe

import "github.com/abtransitionit/luc/pkg/phase"

var Phases = []phase.Phase{
	phase.SetPhase("ssh", ssh, SshDescription),
	phase.SetPhase("cpluc", cpluc, CpLucDescription),
	phase.SetPhase("check", check, CheckDescription),
	phase.SetPhase("update", update, UpdateDescription),
	phase.SetPhase("sysctl", sysctl, SysctlDescription),
	phase.SetPhase("selinux", selinux, SelinuxDescription),
	phase.SetPhase("dnfapt", dnfapt, DnfaptDescription),
	phase.SetPhase("service", service, ServiceDescription),
	phase.SetPhase("cplane", cplane, CplaneDescription),
	phase.SetPhase("worker", worker, WorkerDescription),
	phase.SetPhase("kubectl", kubectl, KubectlDescription),
	phase.SetPhase("helm", helm, HelmDescription),
	phase.SetPhase("cni", cni, CniDescription),
	phase.SetPhase("health", health, HealthDescription),
	phase.SetPhase("custom", custom, CustomDescription),
	phase.SetPhase("secure", secure, SecureDescription),
}
