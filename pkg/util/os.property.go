/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/

package util

import (
	"fmt"
	"os"
	"os/user"
	"runtime"
	"sort"
	"strings"

	"github.com/jedib0t/go-pretty/table"
	"github.com/opencontainers/selinux/go-selinux"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
)

type PropertyHandler func() (string, error)

// map a string to a function
var OsPropertyMap = map[string]PropertyHandler{
	"cpu":          getCpu,
	"cgroup":       getCgroupVersion,
	"init":         getInitSystem,
	"host":         getHost,
	"netip":        getNetIp,
	"netgateway":   getNetGateway,
	"osuser":       getOsUser,
	"ostype":       getOsType,
	"osarch":       getOsArch,
	"osversion":    getOsVersion,
	"osdistro":     getOsDistro,
	"oskversion":   getOsKernelVersion,
	"osfamily":     getOsFamily,
	"path":         getPath,
	"ram":          getRam,
	"selstatus":    getSelinuxStatus,
	"selmode":      getSelinuxMode,
	"uuid":         getUuid,
	"selinfos":     getSelinuxInfos,
	"osinfos":      getOsInfos,
	"rebootstatus": getReboot,
}

// return OsPropertyMap

func getReboot() (string, error) {
	// Ensure we're on Linux
	osType, err := getOsType()
	if err != nil {
		return "", fmt.Errorf("could not detect OS type: %w", err)
	}
	if osType != "linux" {
		return "", fmt.Errorf("unsupported OS type: %s (only linux is supported)", osType)
	}

	// Detect the OS family
	osFamily, err := getOsFamily()
	if err != nil {
		return "", fmt.Errorf("could not detect OS family: %w", err)
	}

	// Select appropriate command
	var cli string
	switch strings.TrimSpace(osFamily) {
	case "debian":
		cli = "test -f /var/run/reboot-required && echo true || echo false"
	case "rhel":
		cli = "	command -v needs-restarting >/dev/null && needs-restarting -r | grep -q 'Reboot is required' && echo true || echo false"
	default:
		return "", fmt.Errorf("unsupported OS family: %s", osFamily)
	}

	// Run the command
	output, err := RunCLILocal(cli)
	if err != nil {
		return "", fmt.Errorf("failed to check reboot requirement: %w", err)
	}

	return strings.TrimSpace(output), nil
}

func getHost() (string, error) {
	osType, err := getOsType()
	if err != nil {
		return "", err
	}
	if osType != "linux" {
		return "Unsupported OS", nil
	}
	cmd := "systemd-detect-virt"
	out, err := RunCLILocal(cmd)
	if err != nil {
		return "", err
	}
	return out, nil
}

func getCpu() (string, error) {
	output, err := cpu.Info()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", output[0].Cores), nil
}

func getInitSystem() (string, error) {
	output, err := RunCLILocal("ps -p 1 -o comm=")
	if err != nil {
		return "", fmt.Errorf("getting init > %v", err)
	}
	if strings.Contains(output, "systemd") {
		return "systemd (cgroup v2)", nil
	}
	return "initd (likely cgroup v1)", nil
}

func getRam() (string, error) {
	output, err := mem.VirtualMemory()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", output.Total/(1024*1024*1024)), nil
}

func getOsUser() (string, error) {
	output, err := user.Current()
	if err != nil {
		return "", err
	}
	return output.Username, nil
}

func getOsType() (string, error) {
	return runtime.GOOS, nil
}

func getOsVersion() (string, error) {
	info, err := host.Info()
	if err != nil {
		return "", err
	}
	return info.PlatformVersion, nil
}

func getOsDistro() (string, error) {
	info, err := host.Info()
	if err != nil {
		return "", err
	}
	return info.Platform, nil
}

func getOsKernelVersion() (string, error) {
	info, err := host.Info()
	if err != nil {
		return "", err
	}
	return info.KernelVersion, nil
}

func getOsFamily() (string, error) {
	info, err := host.Info()
	if err != nil {
		return "", err
	}
	return info.PlatformFamily, nil
}

func getOsArch() (string, error) {
	return runtime.GOARCH, nil
}

func getOsInfos() (string, error) {
	family, err := getOsFamily()
	if err != nil {
		return "", fmt.Errorf("osfamily: %v", err)
	}

	distro, err := getOsDistro()
	if err != nil {
		return "", fmt.Errorf("osdistro: %v", err)
	}

	version, err := getOsVersion()
	if err != nil {
		return "", fmt.Errorf("osversion: %v", err)
	}

	kernel, err := getOsKernelVersion()
	if err != nil {
		return "", fmt.Errorf("oskversion: %v", err)
	}

	return fmt.Sprintf("family: %-6s :: distro: %-10s :: OsVersion: %-6s :: OsKernelVersion: %s", family, distro, version, kernel), nil
}

func getPath() (string, error) {
	path := os.Getenv("PATH")
	if path == "" {
		return "", fmt.Errorf("PATH environment variable is not set")
	}
	return path, nil
}

func getSelinuxStatus() (string, error) {
	if selinux.GetEnabled() {
		return "enabled", nil
	}
	return "disabled", nil
}

func getSelinuxMode() (string, error) {
	switch selinux.EnforceMode() {
	case selinux.Enforcing:
		return "enforcing", nil
	case selinux.Permissive:
		return "permissive", nil
	case selinux.Disabled:
		return "disabled", nil
	default:
		return "unknown", nil
	}
}

func getSelinuxInfos() (string, error) {
	status, err := getSelinuxStatus()
	if err != nil {
		return "", fmt.Errorf("selstatus: %v", err)
	}

	mode, err := getSelinuxMode()
	if err != nil {
		return "", fmt.Errorf("selmode: %v", err)
	}

	return fmt.Sprintf("status: %-10s :: mode: %s", status, mode), nil
}

func getUuid() (string, error) {
	cmd := "sudo cat /sys/class/dmi/id/product_uuid"
	output, err := RunCLILocal(cmd)
	if err != nil {
		return "", fmt.Errorf("getUuid failed: %w", err)
	}
	return output, nil
}

func getCgroupVersion() (string, error) {
	content, err := os.ReadFile("/proc/self/cgroup")
	if err != nil {
		return "", fmt.Errorf("getting cgroup > %w", err)
	}
	if strings.Contains(string(content), "0::/") {
		return "v2", nil
	}
	return "v1", nil
}

// return "", fmt.Errorf("getting net-ip > %s", cErr)

func getNetIp() (string, error) {
	cmd := "curl -s ifconfig.me -4"
	output, err := RunCLILocal(cmd)
	if err != nil {
		return "", fmt.Errorf("getting net-ip > %v", err)
	}
	return output, nil
}

func getNetGateway() (string, error) {
	cmd := "ip route get 2.2.2.2"
	output, err := RunCLILocal(cmd)
	if err != nil {
		return "", fmt.Errorf("getting net-ip > %v", err)
	}
	// First line only
	line := strings.Split(output, "\n")[0]
	return strings.TrimSpace(line), nil
}

func GetOsPropertyMap() map[string]PropertyHandler {
	return OsPropertyMap
}

// Example Usage:
//
//	props := []string{"cpu", "ram", "osarch", "uuid", "cgroup"}
//
//	for _, prop := range props {
//		value, err := util.GetLocalProperty(prop)
//		if err != nil {
//			// logx.L.Debugf("%s", err)
//			continue
//		}
//		fmt.Printf("prop: %s value: %s\n", prop, value)
//	}
func GetLocalProperty(property string) (string, error) {
	handler, ok := OsPropertyMap[property]
	if !ok {
		return "", fmt.Errorf("❌ unknown property requested: %s", property)
	}

	output, err := handler()
	if err != nil {
		return "", fmt.Errorf("❌ error getting %s: %w", property, err)
	}

	return output, nil
}

// Idea - execute this function remotly
// output, err := RunCLILocal(cli)
// switch between local and remote transparently.
//
//	type CommandRunner interface {
//	  Run(cmd string) (string, error)
//	}
func GetRemoteProperty(property string, vm string) (string, error) {
	cmd := fmt.Sprintf(`luc util getprop %s`, property)
	return RunCLIRemote2(cmd, vm)
}

func ShowMapProperty() {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)

	// Simple header
	t.AppendHeader(table.Row{"Property Name"})

	// sort keys
	var listPropertyName []string
	for name := range OsPropertyMap {
		listPropertyName = append(listPropertyName, name)
	}
	sort.Strings(listPropertyName)

	// Add rows
	for _, name := range listPropertyName {
		t.AppendRow(table.Row{
			name,
		})
	}

	// Render with default style
	t.Render()
}
