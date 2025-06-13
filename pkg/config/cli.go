/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/

package config

import (
	"errors"
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/abtransitionit/luc/pkg/errorx"
	"github.com/jedib0t/go-pretty/table"
	"go.uber.org/zap"
)

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
		UrlType: UrlXxx,
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

// # Purpose
//
// - get a CLI's metadata for a given CLI in a CliConfigMap.
//
// # Parameters
//   - cliName: The name of the CLI.
//
// # Returns
// - CLIConfig: A struct containing metadata about the CLI.
// - bool: A boolean indicating whether the CLI's metadata was found or not.
//
// # Example
//
//	config, ok := GetCLIConfigMap("cobra")
//	if !ok {
//	    log.Fatalf("CLI config not found")
//	}
//	fmt.Println("Download URL:", config.Url)
//
// # Notes
//   - The function performs a lookup in a pre-defined internal map of CLI configurations (cliConfigMap).
//   - If the CLI name does not exist in the map, the returned boolean will be `false` and the CLIConfig will be zero-valued.
func GetCLIConfigMap(cliName string) (CLIConfig, bool) {
	c, ok := cliConfigMap[cliName]
	return c, ok
}

// get a specific property of a CLI in a CliConfigMap
//
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
		log.Debugf("❌ use GetCliSpecificUrl instead")
		return errorx.StringError("", "", errors.New("use GetCliSpecificUrl instead"))
	case "urltype":
		value = string(cliConf.UrlType)
	default:
		msg := fmt.Sprintf("property '%s' not found", property)
		log.Debugf("❌ %s", msg)
		return errorx.StringError("find property (%s) in cli ", cliConf.Name, errors.New(""))
	}
	return value, nil
}

// returns a CLI:Url whith all placeholders replaced
//
// Supported placeholders:
//   - $NAME: replaced by the CLI name (e.g., "kubectl")
//   - $TAG:  replaced by the version tag (e.g., "v1.32.2")
//   - $OS:   replaced by the provided or detected operating system
//   - $ARCH: replaced by the provided or detected architecture
//
// Parameters:
//   - log:  a logger
//   - name: the CLI name (must exist in cliConfigMap)
//   - osArch (optional): optional ordered comma separated values of the following
//   	-- ARCH : (i.e., "linux", "darwin", "windows")
//   	-- OS   : (i.e., "amd64", "arm64")
//
//
// Example usage:
//
//	url,_ := config.GetCliSpecificUrl("kubectl", "linux", "amd64")
//	url,_ := config.GetCliSpecificUrl("helm") // OS and Arch auto-detected
//	url,_ := config.GetCliSpecificUrl(logx.L, "helm") // OS and Arch auto-detected
//
// Notes:
//
// - If ARCH or OS are not provided, they will be get/inferred at runtime.

func GetCliSpecificUrl(log *zap.SugaredLogger, cliName string, osArch ...string) (string, error) {
	cliConf, exists := cliConfigMap[cliName]
	if !exists {
		msg := fmt.Sprintf("CLI (%s) not found in map", cliName)
		log.Debugf("❌ %s", msg)
		return errorx.StringError("found element (%s) in map", cliName, errors.New(""))
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

	// Replace placeholders in URL
	url := cliConf.Url
	url = strings.ReplaceAll(url, "$NAME", cliConf.Name)
	url = strings.ReplaceAll(url, "$TAG", cliConf.Tag)
	url = strings.ReplaceAll(url, "$OS", osType)
	url = strings.ReplaceAll(url, "$ARCH", archType)

	return url, nil
}

// # Purpose
//
// - checks if a CLI:Url is CURLable (based on UrlType).
//
// # Returns
//   - bool: true if the UrlType is curlable, false otherwise
//   - error: always nil in the current implementation
//
// # Possible returns
//   - (true, nil)  → when the CLI:UrlType is considered OK
//   - (false, nil) → when the CLI:UrlType is considered NOTOK
//
// # Usage examples
//
//   // Example 1 (variable-based usage)
//   var u UrlType = UrlExe
//   if ok, _ := u.IsCurlable(); ok {
//       fmt.Println("The URL type is curlable.")
//   } else {
//       fmt.Println("The URL type is not curlable.")
//   }func (u UrlType) IsCurlable() (bool, error) {
//
//   // Example 2
//   ok, _ := UrlExe.IsCurlable()
//   fmt.Println("UrlExe curlable?", ok) // → true
//
//   // Example 3
//   ok, _ = UrlTgz.IsCurlable()
//   fmt.Println("UrlTgz curlable?", ok) // → true
//
//   // Example 4
//   ok, _ = UrlGit.IsCurlable()
//   fmt.Println("UrlGit curlable?", ok) // → false

func (u UrlType) IsCurlable() (bool, error) {
	switch u {
	case UrlExe, UrlTgz:
		return true, nil
	default:
		return false, nil
	}
}

// # Purpose
//
// - Checks if a CLI:Url is GITable (based on UrlType).
//
// # Returns
//   - bool: true if the UrlType is gitable, false otherwise
//   - error: always nil in the current implementation
//
// # Possible returns
//   - (true, nil)  → when the CLI:UrlType is considered OK
//   - (false, nil) → when the CLI:UrlType is considered NOTOK
//
// # Usage examples
//
//	// Example 1 (variable-based usage)
//	var u UrlType = UrlGit
//	if ok, _ := u.IsGitable(); ok {
//	    fmt.Println("The URL type is gitable.")
//	} else {
//	    fmt.Println("The URL type is not gitable.")
//	}
//
//	// Example 2
//	ok, _ := UrlGit.IsGitable()
//	fmt.Println("UrlGit gitable?", ok) // → true
//
//	// Example 3
//	ok, _ = UrlExe.IsGitable()
//	fmt.Println("UrlExe gitable?", ok) // → false
//
//	// Example 4
//	ok, _ = UrlTgz.IsGitable()
//	fmt.Println("UrlTgz gitable?", ok) // → false
func (u UrlType) IsGitable() (bool, error) {
	switch u {
	case UrlGit:
		return true, nil
	default:
		return false, nil
	}
}

func ShowCliConfigMap() {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)

	// Simple header
	t.AppendHeader(table.Row{"CLI Name", "Version", "Type", "Doc", "Git"})

	// Add rows
	for name, cfg := range cliConfigMap {
		t.AppendRow(table.Row{
			name,
			cfg.Tag,
			cfg.UrlType,
			cfg.DocUrl,
			cfg.GitUrl,
		})
	}

	// Render with default style
	t.Render()
}
