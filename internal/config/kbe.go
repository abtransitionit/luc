/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/

package config

const (
	K8sVersion = "1.32.0"
)
const (
	// KBE version
	KbeVersion = K8sVersion // follow up kubernetes version

	// KBE nodes
	KbeListNodeList = "o1u o2a o3r o5d o6f"
	KbeNodeCplane   = "o1u"
	KbeListWorker   = "o1u"

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
