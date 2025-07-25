/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/

package config

import "github.com/abtransitionit/luc/pkg/config"

const (
	K8sVersion              = "1.32.0"
	CrCrioDnfaptRepoVersion = "v1.32"
	K8sDnfaptRepoVersion    = "v1.32"
)
const (
	// KBE version
	KbeVersion = K8sVersion // follow up kubernetes version

	// where to install KBE
	// KbeListNode = "o1u o2a o3r o5d o6f" // o6f generate error
	KbeListNode = "o1u o2a o3r o4f o5d"
	// KbeListNode = "o1u o2a"
	// KbeListNode = "o1u o5d"
	// KbeListNode = "o2a"
	// KbeListNode       = "o1u"
	KbeListNodeCplane = "o1u"
	KbeListNodeWorker = "o1u"

	// KBE CLI
	KbeKubeadmCliVersion = K8sVersion
	KbeKubectlCliVersion = K8sVersion
	KbeHelmVersion       = "3.17.3" // compatible with K8sVersion

	// KBE component
	KbeKubeletServiceVersion = K8sVersion

	// KBE custom CIDR
	KbePodCidr     = "192.168.0.0/16"
	KbeServiceCidr = "172.16.0.0/16"

	// dnfapt CR/CRIO
	KbeCrCrioDnfaptRepoName       = "kbe-crio"
	KbeCrCrioDnfaptPackageVersion = "v1.32"

	// dnfapt K8s
	KbeK8sDnfaptRepotName      = "kbe-k8s"
	KbeK8sDnfaptPackageVersion = "v1.32"

	// OS Kernel
	KbeKernelFileName = "99-kbe.conf"
)

var KbeDnfaptCliConfigMap = config.CustomCLIConfigMap{
	"crio": {
		Name:      "crio",
		Version:   "1.7.1",
		DstFolder: "/usr/local/bin", // default: /opt/cni/bin
	},
}
var KbeDnfaptRepoConfigMap = config.CustomDnfaptRepoConfigMap{
	"crio": {
		Name:    "crio",
		CName:   "kbe-crio",
		Version: CrCrioDnfaptRepoVersion,
	},
	"k8s": {
		Name:    "k8s",
		CName:   "kbe-k8s",
		Version: K8sDnfaptRepoVersion,
	},
}

var KbeGoCliConfigMap = config.CustomCLIConfigMap{
	"kubeadm": {
		Name:      "kubeadm",
		Version:   KbeVersion,
		DstFolder: "/usr/local/bin",
	},
	"kubectl": {
		Name:      "kubectl",
		Version:   KbeVersion,
		DstFolder: "/usr/local/bin",
	},
	"helm": {
		Name:      "helm",
		Version:   "3.17.3",
		DstFolder: "/usr/local/bin",
	},
}
