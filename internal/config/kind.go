/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/

package config

import "github.com/abtransitionit/luc/pkg/config"

const (
	// where to install kind
	// KindVm = "o1u"
	KindVm        = "o1u o2a"
	KindDnfaptCli = "uidmap"

	// Apparmor service configuration file content
	ApparmorServiceConf = `
		# Allow rootlesskit to create user namespaces (userns)
		# Ref: https://ubuntu.com/blog/ubuntu-23-10-restricted-unprivileged-user-namespaces
		abi <abi/4.0>,
		include <tunables/global>

		/usr/local/bin/rootlesskit/rootlesskit flags=(unconfined) {
			userns,

			# Site-specific additions and overrides. See local/README for details.
			include if exists <local/usr.local.bin.rootlesskit.rootlesskit>
		}
	`
	// Apparmor service configuration file path
	ApparmorFilePath = "/etc/apparmor.d/usr.local.bin.rootlesskit.rootlesskit"
)

// # Purpose
//
// - List of go CLI to install
var KindDaCliConfigMap = config.CustomCLIConfigMap{
	"uidmap": {
		Name:    "uidmap",
		Version: "",
	},
}

// # Purpose
//
// - List of go CLI to install
var KindGoCliConfigMap = config.CustomCLIConfigMap{
	"cni": {
		Name:      "cni",
		Version:   "1.7.1",
		DstFolder: "/usr/local/bin", // default: /opt/cni/bin
	},
	"containerd": {
		Name:      "containerd",
		Version:   "2.1.1",
		DstFolder: "/usr/local/bin",
	},
	"kind": {
		Name:      "kind",
		Version:   "latest",
		DstFolder: "/usr/local/bin",
	},
	"nerdctl": {
		Name:      "nerdctl",
		Version:   "2.1.2",
		DstFolder: "/usr/local/bin",
	},
	"rootlesskit": {
		Name:      "rootlesskit",
		Version:   "2.3.5",
		DstFolder: "/usr/local/bin",
	},
	"runc": {
		Name:      "runc",
		Version:   "1.3.0",
		DstFolder: "/usr/local/bin",
	},
	"slirp4netns": {
		Name:      "slirp4netns",
		Version:   "1.3.3",
		DstFolder: "/usr/local/bin",
	},
}
