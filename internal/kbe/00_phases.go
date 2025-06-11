/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
// internal/kbe/phases.go
package kbe

import "github.com/abtransitionit/luc/pkg/deploy"

var Phases = []deploy.Phase{
	deploy.SetPhase("ssh", ssh, SshDescription),
	deploy.SetPhase("cpluc", cpluc, CpLucDescription),
	deploy.SetPhase("check", check, CheckDescription),
	deploy.SetPhase("update", update, UpdateDescription),
	deploy.SetPhase("sysctl", sysctl, SysctlDescription),
	deploy.SetPhase("selinux", selinux, SelinuxDescription),
	deploy.SetPhase("dnfapt", dnfapt, DnfaptDescription),
	deploy.SetPhase("service", service, ServiceDescription),
	deploy.SetPhase("cplane", cplane, CplaneDescription),
	deploy.SetPhase("worker", worker, WorkerDescription),
	deploy.SetPhase("kubectl", kubectl, KubectlDescription),
	deploy.SetPhase("helm", helm, HelmDescription),
	deploy.SetPhase("cni", cni, CniDescription),
	deploy.SetPhase("health", health, HealthDescription),
	deploy.SetPhase("custom", custom, CustomDescription),
	deploy.SetPhase("secure", secure, SecureDescription),
}
