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
	"github.com/jedib0t/go-pretty/table"
	"go.uber.org/zap"
)

type UrlType string

const (
	UrlExe UrlType = "exe"
	UrlTgz UrlType = "tgz"
	UrlGit UrlType = "git"
	UrlGo  UrlType = "go"
	UrlOth UrlType = "oth"
	// etc.
)

type CLIConfig struct {
	Name    string
	Tag     string
	Url     string
	DocUrl  string
	GitUrl  string
	UrlType UrlType
}

type CLIConfigMap map[string]CLIConfig

// # Purpose
//
// - get a CLI's metadata for a given CLI in a SharedCliConfigMap.
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
//	config, ok := GetCLIConfig("cobra")
//	if !ok {
//	    log.Fatalf("CLI config not found")
//	}
//	fmt.Println("Download URL:", config.Url)
//
// # Notes
//   - The function performs a lookup in a pre-defined internal map of CLI configurations (cliConfigMap).
//   - If the CLI name does not exist in the map, the returned boolean will be `false` and the CLIConfig will be zero-valued.
func GetCLIConfig(cliName string) (CLIConfig, bool) {
	c, ok := SharedCliConfigMap[cliName]
	return c, ok
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
//	// Example 1 (variable-based usage)
//	var u UrlType = UrlExe
//	if ok, _ := u.IsCurlable(); ok {
//	    fmt.Println("The URL type is curlable.")
//	} else {
//	    fmt.Println("The URL type is not curlable.")
//	}func (u UrlType) IsCurlable() (bool, error) {
//
//	// Example 2
//	ok, _ := UrlExe.IsCurlable()
//	fmt.Println("UrlExe curlable?", ok) // → true
//
//	// Example 3
//	ok, _ = UrlTgz.IsCurlable()
//	fmt.Println("UrlTgz curlable?", ok) // → true
//
//	// Example 4
//	ok, _ = UrlGit.IsCurlable()
//	fmt.Println("UrlGit curlable?", ok) // → false
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

func (obj CLIConfigMap) String() string {
	t := table.NewWriter()
	t.SetStyle(table.StyleLight)
	t.SetTitle("CLI Config Map")
	t.AppendHeader(table.Row{"CLI Name", "Version", "Type", "Doc", "Git"})

	for name, item := range obj {
		t.AppendRow(table.Row{
			name,
			item.Tag,
			item.UrlType,
			item.DocUrl,
			item.GitUrl,
		})
	}

	return t.Render()
}

// get a specific property of a CLI in a SharedCliConfigMap
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
	cliConf, ok := SharedCliConfigMap[name]
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
	cliConf, exists := SharedCliConfigMap[cliName]
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
