/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/

package config

// private variable called the CliConfigMap
var cliConfigMap = map[string]CLIConfig{
	"cobra": {
		Name:    "cobra",
		Tag:     "latest",
		Url:     "github.com/spf13/$NAME-cli@$TAG",
		DocUrl:  "https://cobra.dev",
		GitUrl:  "https://github.com/spf13/cobra-cli",
		UrlType: UrlGo,
	},
	"luc": {
		Name:    "luc",
		Tag:     "0.0.1",
		Url:     "https://github.com/abtransitionit/$NAME/releases/download/v$TAG-main/$NAME-$OS-$ARCH",
		DocUrl:  "https://github.com/abtransitionit/luc",
		GitUrl:  "https://github.com/abtransitionit/luc",
		UrlType: UrlExe,
	},
	"containerd": {
		Name:    "containerd",
		Tag:     "2.1.1",
		Url:     "https://github.com/$NAME/$NAME/releases/download/v$TAG/$NAME-$TAG-$OS-$ARCH.tar.gz",
		DocUrl:  "https://github.com/containerd/nerdctl/blob/main/docs/command-reference.md",
		GitUrl:  "https://github.com/containerd/nerdctl",
		UrlType: UrlTgz,
	},
	"helm": {
		Name:    "helm",
		Tag:     "v3.17.3",
		Url:     "https://github.com/helm/helm.git",
		DocUrl:  "https://helm.sh/",
		GitUrl:  "https://github.com/helm/helm",
		UrlType: UrlGit,
	},
	"kind": {
		Name:    "kind",
		Tag:     "latest",
		Url:     "https://$NAME.sigs.k8s.io/dl/$TAG/$NAME-$OS-$ARCH",
		DocUrl:  "https://kind.sigs.k8s.io/",
		GitUrl:  "https://github.com/kubernetes-sigs/kind",
		UrlType: UrlExe,
	},
	"kubebuilder": {
		Name:    "kubebuilder",
		Tag:     "v4.5.2",
		Url:     "https://github.com/kubernetes-sigs/kubebuilder.git",
		DocUrl:  "https://kubebuilder.io",
		GitUrl:  "https://github.com/kubernetes-sigs/kubebuilder",
		UrlType: UrlGit,
	},
	"kubectl": {
		Name:    "kubectl",
		Tag:     "v1.32.2",
		Url:     "https://dl.k8s.io/release/$TAG/bin/$OS/$ARCH/$NAME",
		DocUrl:  "https://kubernetes.io/docs/reference/kubectl/",
		GitUrl:  "https://github.com/kubernetes/kubectl",
		UrlType: UrlExe,
	},
	"nerdctl": {
		Name:    "nerdctl",
		Tag:     "2.1.2",
		Url:     "https://github.com/containerd/$NAME/releases/download/v$TAG/nerdctl-$TAG-$OS-$ARCH.tar.gz",
		DocUrl:  "https://github.com/containerd/nerdctl/blob/main/docs/command-reference.md",
		GitUrl:  "https://github.com/containerd/nerdctl",
		UrlType: UrlTgz,
	},
	"nerdctlf": {
		Name:    "nerdctlf",
		Tag:     "2.1.2",
		Url:     "https://github.com/containerd/$NAME/releases/download/v$TAG/nerdctl-$TAG-$OS-$ARCH.tar.gz",
		DocUrl:  "https://github.com/containerd/nerdctl/blob/main/docs/command-reference.md",
		GitUrl:  "https://github.com/containerd/nerdctl",
		UrlType: UrlTgz,
	},
	"rootlesskit": {
		Name:    "rootlesskit",
		Tag:     "2.3.5",
		Url:     "https://github.com/rootless-containers/$NAME/releases/download/v$TAG/$NAME-$UNAME.tar.gz",
		DocUrl:  "https://github.com/rootless-containers/rootlesskit",
		GitUrl:  "https://github.com/rootless-containers/rootlesskit",
		UrlType: UrlTgz,
	},
	"runc": {
		Name:    "runc",
		Tag:     "1.3.0",
		Url:     "https://github.com/opencontainers/$NAME/releases/download/v$TAG/$NAME.$ARCH",
		DocUrl:  "https://github.com/opencontainers/runc/tree/main",
		GitUrl:  "https://github.com/opencontainers/runc",
		UrlType: UrlOth,
	},
	"slirp4netns": {
		Name:    "slirp4netns",
		Tag:     "1.3.3",
		Url:     "https://github.com/rootless-containers/$NAME/releases/download/v$TAG/$NAME-$UNAME",
		DocUrl:  "https://man.archlinux.org/man/extra/slirp4netns/slirp4netns.1.en",
		GitUrl:  "https://github.com/rootless-containers/slirp4netns",
		UrlType: UrlExe,
	},
	"sonobuoy": {
		Name:    "sonobuoy",
		Tag:     "v0.57.3",
		Url:     "https://github.com/vmware-tanzu/sonobuoy.git",
		DocUrl:  "https://sonobuoy.io/docs/main/",
		GitUrl:  "https://github.com/vmware-tanzu/sonobuoy",
		UrlType: UrlGit,
	},
}
