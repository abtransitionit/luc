/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/

package config

// # Purpose
//
// - private variable called the SharedCliConfigMap
//
// # Notes
//
// - map a string to a struct
// - map a key    to a value that have a type
var SharedCliConfigMap = CLIConfigMap{
	// "toto": {
	// 	Name:    "toto",
	// 	Tag:     "latest",
	// 	Url:     "github.com/spf13/$NAME-cli@$TAG",
	// 	DocUrl:  "https://cobra.dev",
	// 	GitUrl:  "https://github.com/spf13/cobra-cli",
	// 	UrlType: UrlGit,
	// },
	"cni": {
		Name: "cni",
		// Tag:     "1.7.1",
		Url:     "https://github.com/containernetworking/plugins/releases/download/v$TAG/$NAME-plugins-$OS-$ARCH-v$TAG.tgz",
		DocUrl:  "https://www.cni.dev/",
		GitUrl:  "https://github.com/containernetworking/plugins",
		UrlType: UrlTgz,
	},
	"cobra": {
		Name: "cobra",
		// Tag:     "latest",
		Url:     "github.com/spf13/$NAME-cli@$TAG",
		DocUrl:  "https://cobra.dev",
		GitUrl:  "https://github.com/spf13/cobra-cli",
		UrlType: UrlGo,
	},
	"containerd": {
		Name: "containerd",
		// Tag:     "2.1.1",
		Url:     "https://github.com/$NAME/$NAME/releases/download/v$TAG/$NAME-$TAG-$OS-$ARCH.tar.gz",
		DocUrl:  "https://github.com/containerd/nerdctl/blob/main/docs/command-reference.md",
		GitUrl:  "https://github.com/containerd/nerdctl",
		UrlType: UrlTgz,
	},
	"go": {
		Name: "go",
		// Tag:     "1.24.4",
		Url:     "https://go.dev/dl/$TAG.$TAG.$OS-$ARCH.tar.gz",
		DocUrl:  "https://go.dev/dl/",
		GitUrl:  "https://github.com/golang/go",
		UrlType: UrlTgz,
	},
	"helm": {
		Name: "helm",
		// Tag:     "v3.17.3",
		Url:     "https://github.com/helm/helm.git",
		DocUrl:  "https://helm.sh/",
		GitUrl:  "https://github.com/helm/helm",
		UrlType: UrlGit,
	},
	"kind": {
		Name: "kind",
		// Tag:     "latest",
		Url:     "https://$NAME.sigs.k8s.io/dl/$TAG/$NAME-$OS-$ARCH",
		DocUrl:  "https://kind.sigs.k8s.io/",
		GitUrl:  "https://github.com/kubernetes-sigs/kind",
		UrlType: UrlExe,
	},
	"kubebuilder": {
		Name: "kubebuilder",
		// Tag:     "v4.5.2",
		Url:     "https://github.com/kubernetes-sigs/kubebuilder.git",
		DocUrl:  "https://kubebuilder.io",
		GitUrl:  "https://github.com/kubernetes-sigs/kubebuilder",
		UrlType: UrlGit,
	},
	"kubeadm": {
		Name: "kubeadm",
	},
	"kubectl": {
		Name: "kubectl",
		// Tag:     "v1.32.2",
		Url:     "https://dl.k8s.io/release/$TAG/bin/$OS/$ARCH/$NAME",
		DocUrl:  "https://kubernetes.io/docs/reference/kubectl/",
		GitUrl:  "https://github.com/kubernetes/kubectl",
		UrlType: UrlExe,
	},
	"luc": {
		Name: "luc",
		// Tag:     "0.0.1",
		Url:     "https://github.com/abtransitionit/$NAME/releases/download/v$TAG-main/$NAME-$OS-$ARCH",
		DocUrl:  "https://github.com/abtransitionit/luc",
		GitUrl:  "https://github.com/abtransitionit/luc",
		UrlType: UrlExe,
	},
	"nerdctl": {
		Name: "nerdctl",
		// Tag:     "2.1.2",
		Url:     "https://github.com/containerd/$NAME/releases/download/v$TAG/nerdctl-$TAG-$OS-$ARCH.tar.gz",
		DocUrl:  "https://github.com/containerd/nerdctl/blob/main/docs/command-reference.md",
		GitUrl:  "https://github.com/containerd/nerdctl",
		UrlType: UrlTgz,
	},
	"nerdctlf": {
		Name: "nerdctlf",
		// Tag:     "2.1.2",
		Url:     "https://github.com/containerd/$NAME/releases/download/v$TAG/nerdctl-$TAG-$OS-$ARCH.tar.gz",
		DocUrl:  "https://github.com/containerd/nerdctl/blob/main/docs/command-reference.md",
		GitUrl:  "https://github.com/containerd/nerdctl",
		UrlType: UrlTgz,
	},
	"rootlesskit": {
		Name: "rootlesskit",
		// Tag:     "2.3.5",
		Url:     "https://github.com/rootless-containers/$NAME/releases/download/v$TAG/$NAME-$UNAME.tar.gz",
		DocUrl:  "https://github.com/rootless-containers/rootlesskit",
		GitUrl:  "https://github.com/rootless-containers/rootlesskit",
		UrlType: UrlTgz,
	},
	"runc": {
		Name: "runc",
		// Tag:     "1.3.0",
		Url:     "https://github.com/opencontainers/$NAME/releases/download/v$TAG/$NAME.$ARCH",
		DocUrl:  "https://github.com/opencontainers/runc/tree/main",
		GitUrl:  "https://github.com/opencontainers/runc",
		UrlType: UrlExe,
	},
	"slirp4netns": {
		Name: "slirp4netns",
		// Tag:     "1.3.3",
		Url:     "https://github.com/rootless-containers/$NAME/releases/download/v$TAG/$NAME-$UNAME",
		DocUrl:  "https://man.archlinux.org/man/extra/slirp4netns/slirp4netns.1.en",
		GitUrl:  "https://github.com/rootless-containers/slirp4netns",
		UrlType: UrlExe,
	},
	"sonobuoy": {
		Name: "sonobuoy",
		// Tag:     "v0.57.3",
		Url:     "https://github.com/vmware-tanzu/$NAME.git",
		DocUrl:  "https://sonobuoy.io/docs/main/",
		GitUrl:  "https://github.com/vmware-tanzu/sonobuoy",
		UrlType: UrlGit,
	},
}

var SharedDnfaptRepoConfigMap = DnfaptRepoConfigMap{
	"crio": {
		Name:    "crio",
		UrlRepo: "https://download.opensuse.org/repositories/isv:/cri-o:/stable:/$TAG/$PACK/",
		UrlGpg:  "https://download.opensuse.org/repositories/isv:/cri-o:/stable:/$TAG/$PACK/$GPG",
	},
	"k8s": {
		Name:    "k8s",
		UrlRepo: "https://pkgs.k8s.io/core:/stable:/$TAG/$PACK/",
		UrlGpg:  "https://pkgs.k8s.io/core:/stable:/$TAG/$PACK/$GPG",
	},
}
