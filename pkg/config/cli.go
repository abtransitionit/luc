/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/

package config

import (
	"errors"
	"fmt"
	"runtime"
	"strings"

	"github.com/abtransitionit/luc/pkg/errorx"
	"github.com/abtransitionit/luc/pkg/logx"
	"go.uber.org/zap"
)

const (
	NerdctlCliName = "nerdctl" // contaiNERD ConTroL
)

type CLIConfig struct {
	Name    string
	Tag     string
	Url     string
	DocUrl  string
	GitUrl  string
	UrlType string
}

var cliConfigMap = map[string]CLIConfig{
	"cobra": {
		Name:    "cobra",
		Tag:     "latest",
		Url:     "github.com/spf13/$NAME-cli@$TAG",
		DocUrl:  "https://cobra.dev",
		GitUrl:  "https://github.com/spf13/cobra-cli",
		UrlType: "go",
	},
	"containerd": {
		Name:    "containerd",
		Tag:     "2.1.1",
		Url:     "https://github.com/$NAME/$NAME/releases/download/v$TAG/$NAME-$TAG-$OS-$ARCH.tar.gz",
		DocUrl:  "https://github.com/containerd/nerdctl/blob/main/docs/command-reference.md",
		GitUrl:  "https://github.com/containerd/nerdctl",
		UrlType: "tgz",
	},
	"helm": {
		Name:    "helm",
		Tag:     "v3.17.3",
		Url:     "https://github.com/helm/helm.git",
		DocUrl:  "https://helm.sh/",
		GitUrl:  "https://github.com/helm/helm",
		UrlType: "git",
	},
	"kind": {
		Name:    "kind",
		Tag:     "latest",
		Url:     "https://$NAME.sigs.k8s.io/dl/$TAG/$NAME-$OS-$ARCH",
		DocUrl:  "https://kind.sigs.k8s.io/",
		GitUrl:  "https://github.com/kubernetes-sigs/kind",
		UrlType: "file",
	},
	"kubebuilder": {
		Name:    "kubebuilder",
		Tag:     "v4.5.2",
		Url:     "https://github.com/kubernetes-sigs/kubebuilder.git",
		DocUrl:  "https://kubebuilder.io",
		GitUrl:  "https://github.com/kubernetes-sigs/kubebuilder",
		UrlType: "git",
	},
	"kubectl": {
		Name:    "kubectl",
		Tag:     "v1.32.2",
		Url:     "https://dl.k8s.io/release/$TAG/bin/$OS/$ARCH/$NAME",
		DocUrl:  "https://kubernetes.io/docs/reference/kubectl/",
		GitUrl:  "https://github.com/kubernetes/kubectl",
		UrlType: "file",
	},
	"nerdctl": {
		Name:    "nerdctl",
		Tag:     "2.1.2",
		Url:     "https://github.com/containerd/$NAME/releases/download/v$TAG/nerdctl-$TAG-$OS-$ARCH.tar.gz",
		DocUrl:  "https://github.com/containerd/nerdctl/blob/main/docs/command-reference.md",
		GitUrl:  "https://github.com/containerd/nerdctl",
		UrlType: "tgz",
	},
	"rootlesskit": {
		Name:    "rootlesskit",
		Tag:     "2.1.1",
		Url:     "https://github.com/$NAME/$NAME/releases/download/v$TAG/$NAME-$TAG-$OS-$ARCH.tar.gz",
		DocUrl:  "https://github.com/containerd/nerdctl/blob/main/docs/command-reference.md",
		GitUrl:  "https://github.com/rootless-containers/rootlesskit",
		UrlType: "tgz",
	},
	"runc": {
		Name:    "runc",
		Tag:     "1.3.0",
		Url:     "https://github.com/opencontainers/$NAME/releases/download/v$TAG/$NAME.$ARCH",
		DocUrl:  "https://github.com/opencontainers/runc/tree/main",
		GitUrl:  "https://github.com/opencontainers/runc",
		UrlType: "xxx",
	},
	"sonobuoy": {
		Name:    "sonobuoy",
		Tag:     "v0.57.3",
		Url:     "https://github.com/vmware-tanzu/sonobuoy.git",
		DocUrl:  "https://sonobuoy.io/docs/main/",
		GitUrl:  "https://github.com/vmware-tanzu/sonobuoy",
		UrlType: "git",
	},
}

// retrieves a specific property of a CLI from the configuration map.

// # Parameters:
//   - log: a *zap.SugaredLogger used for debug logging.
//   - name: the CLI name to look up in cliConfigMap.
//   - property: the name of the property to retrieve (case-insensitive).
//
// # Supported properties:
//   - "name"
//   - "tag"
//   - "url"
//   - "docurl"
//   - "giturl"
//   - "urltype"
//
// # Returns:
//   - string : the requested CLI property value if successful (empty string on failure)
//   - error  : failure details if any (nil on success)
//
// # Example usage:
//
//	 value, err := GetCliProperty(log, "cobra", "docurl")
//		if err != nil {
//		    log.Warnf("Error: %v", err)
//		} else {
//
//		    fmt.Println("Cobra doc URL:", value)
//		}
func GetCliProperty(log *zap.SugaredLogger, name string, property string) (string, error) {
	value := ""
	cliConf, ok := cliConfigMap[name]
	if !ok {
		log.Debugf("❌ CLI (%s) not found in map", name)
		return errorx.StringError("find CLI (%s) in map", name, errors.New(""))
	}
	switch strings.ToLower(property) {
	case "name":
		value = cliConf.Name
	case "tag":
		value = cliConf.Tag
	case "url":
		value = cliConf.Url
	case "docurl":
		value = cliConf.DocUrl
	case "giturl":
		log.Debugf("❌ use GetCliUrl instead")
		return errorx.StringError("", "", errors.New("use GetCliUrl instead"))
	case "urltype":
		value = cliConf.UrlType
	default:
		msg := fmt.Sprintf("property '%s' not found", property)
		log.Debugf("❌ %s", msg)
		return errorx.StringError("find property (%s) in cli ", cliConf.Name, errors.New(""))
	}
	return value, nil
}

// GetCliUrl returns the resolved download URL of a CLI tool
// by replacing placeholders in the configured URL.
//
// Supported placeholders:
//   - $NAME: replaced by the CLI name (e.g., "kubectl")
//   - $TAG:  replaced by the version tag (e.g., "v1.32.2")
//   - $OS:   replaced by the provided or detected operating system
//   - $ARCH: replaced by the provided or detected architecture
//
// Parameters:
//   - name: the CLI name (must exist in cliConfigMap)
//   - osArch (optional): provide up to two values:
//     osType (e.g., "linux", "darwin", "windows")
//     archType (e.g., "amd64", "arm64")
//
// If osType or archType are not provided, they will be inferred from runtime.
//
// Example usage:
//
//	url,_ := GetCliUrl("kubectl", "linux", "amd64")
//	url,_ := GetCliUrl("helm") // OS and Arch auto-detected
func GetCliUrl(log *zap.SugaredLogger, name string, osArch ...string) (string, error) {
	cliConf, exists := cliConfigMap[name]
	if !exists {
		msg := fmt.Sprintf("CLI (%s) not found in map", name)
		log.Debugf("❌ %s", msg)
		return errorx.StringError("found element (%s) in map", name, errors.New(""))
	}

	// Detect OS and ARCH if not provided
	osType := runtime.GOOS
	archType := runtime.GOARCH

	if len(osArch) > 0 {
		osType = osArch[0]
	}
	if len(osArch) > 1 {
		archType = osArch[1]
	}

	// Normalize values
	osType = strings.ToLower(osType)
	archType = strings.ToLower(archType)

	// replacements := map[string]string{
	// 	"$NAME": cliConf.Name,
	// 	"$TAG":  cliConf.Tag,
	// 	"$OS":   osType,
	// 	"$ARCH": osArch,
	// }

	// url := cliConf.Url
	// for placeholder, value := range replacements {
	// 	url = strings.ReplaceAll(url, placeholder, value)
	// }

	// Replace placeholders in URL
	url := cliConf.Url
	url = strings.ReplaceAll(url, "$NAME", cliConf.Name)
	url = strings.ReplaceAll(url, "$TAG", cliConf.Tag)
	url = strings.ReplaceAll(url, "$OS", osType)
	url = strings.ReplaceAll(url, "$ARCH", archType)

	return url, nil
}

// prints out information about the map.
//
// Example output:
//
//	cobra      go github.com/spf13/cobra-cli@latest
//	containerd tgz https://github.com/containerd/containerd/releases/download/v2.1.1/containerd-2.1.1-linux-amd64.tar.gz
//
//	List of url type: go tgz
//
// Usage:
//
//   - DisplayCliConfigInfo()
func DisplayCliCondfigInfo() {
	// list name and url for current platform
	for cliName, cliConf := range cliConfigMap {
		url, _ := GetCliUrl(logx.L, cliName)
		fmt.Printf("%12s %5s %s \n", cliName, cliConf.UrlType, url)
	}

	// list all Url types
	types := map[string]bool{}
	for _, conf := range cliConfigMap {
		types[conf.UrlType] = true
	}
	fmt.Printf("\n\nList of url type: ")
	for t := range types {
		fmt.Printf(" %s ", t)
	}
	fmt.Println("\n")
}
