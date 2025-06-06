/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/

package config

const (
	// Kernel
	OsKernelModuleFolder = "/etc/modules-load.d"
	OsKernelParamFolder  = "/etc/sysctl.d"
	// dnf
	OsRhelRepoFolder = "/etc/yum.repos.d"
	// apt
	OsDebianRepoFolder    = "/etc/apt/sources.list.d"
	OsDebianGpgRepoFolder = "/etc/apt/keyrings"
	// Selinux
	OsSelinuxFilePath = "/etc/selinux/config"
	// Binary folder
	OsBinFolder = "/usr/local/bin"
)
