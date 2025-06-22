/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/

package config

const (
	// Kind Apparmor service conf
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
	ApparmorFilePath = "/etc/apparmor.d/usr.local.bin.rootlesskit.rootlesskit"

	// Kind Apparmor service conf
	ContainerdUserServiceConf = `
		[Unit]
		Description=Containerd (Rootless)
		After=network.target

		[Service]
		ExecStart=%h/bin/containerd
		Restart=always
		Environment=PATH=%h/bin:/usr/local/bin:/usr/bin
		Delegate=yes
		NoNewPrivileges=true

		[Install]
		WantedBy=default.target
	`
)
