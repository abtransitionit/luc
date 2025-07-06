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

type PropertyHandler func(...string) (string, error)

// map a string to a function
var OsPropertyMap = map[string]PropertyHandler{
	"cpu":            getCpu,
	"cgroup":         getCgroupVersion,
	"init":           getInitSystem,
	"host":           getHost,
	"userlinger":     getLinger,
	"netip":          getNetIp,
	"netgateway":     getNetGateway,
	"osuser":         getOsUser,
	"ostype":         getOsType, // e.g. linux, windows, darwin
	"osarch":         getOsArch,
	"osversion":      getOsVersion,
	"osdistro":       getOsDistro,
	"oskversion":     getOsKernelVersion,
	"osfamily":       getOsFamily,
	"osinfos":        getOsInfos,
	"path":           getPath,
	"pathext":        getPathExtend,
	"pathtree":       getPathTree,
	"ram":            getRam,
	"selstatus":      getSelinuxStatus,
	"selmode":        getSelinuxMode,
	"selinfos":       getSelinuxInfos,
	"serviceStatus":  getServiceStatus,
	"serviceEnabled": getServiceEnabled,
	"serviceinfos":   getServiceInfos,
	"rebootstatus":   getReboot,
	"uuid":           getUuid,
	"uname":          getUnameM,
}

func getLinger(params ...string) (string, error) {
	if len(params) < 1 {
		return "", fmt.Errorf("user name required")
	}

	// get input
	OsUserName := params[0]

	// play test cli - same as testing if cli : loginctl exists
	cli := fmt.Sprintf(`loginctl show-user %s`, OsUserName)
	if _, err := RunCLILocal(cli); err != nil {
		return "", err
	}
	// now grep is safe
	cli = fmt.Sprintf(`loginctl show-user %s | grep -i linger | cut -d= -f2`, OsUserName)
	output, err := RunCLILocal(cli)
	if err != nil {
		return "", err
	}
	// success
	return output, nil
}
func getServiceInfos(params ...string) (string, error) {
	if len(params) < 1 {
		return "", fmt.Errorf("service name required")
	}
	// get service
	serviceName := params[0]

	// get
	isActive, err := getServiceStatus(serviceName)
	if err != nil {
		return "", fmt.Errorf("serviceStatus: %v", err)
	}

	// get
	isEnabled, err := getServiceEnabled(serviceName)
	if err != nil {
		return "", fmt.Errorf("serviceEnabled: %v", err)
	}

	// return
	return fmt.Sprintf("%-6s / %-6s", isActive, isEnabled), nil

}

func getServiceEnabled(params ...string) (string, error) {
	if len(params) < 1 {
		return "", fmt.Errorf("service name required")
	}
	service := params[0]
	cli := fmt.Sprintf("systemctl is-enabled %s", service)

	return RunCLILocal(cli)
}

func getServiceStatus(params ...string) (string, error) {

	// manage argument
	if len(params) < 1 {
		return "", fmt.Errorf("service name required")
	}
	// get service name
	service := params[0]

	// play cli
	cli := fmt.Sprintf("systemctl is-active %s", service)
	output, err := RunCLILocal(cli)

	// in theses case the cli returns err = nil
	if output == "active" || output == "inactive" || output == "failed" {
		return output, nil
	}

	// manage other real errors
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(output)), nil
}

func getUnameM(_ ...string) (string, error) {
	cli := "uname -m"
	output, err := RunCLILocal(cli)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(output), nil
}

func getReboot(_ ...string) (string, error) {
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

func getHost(_ ...string) (string, error) {
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

func getCpu(_ ...string) (string, error) {
	output, err := cpu.Info()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", output[0].Cores), nil
}

func getInitSystem(_ ...string) (string, error) {
	output, err := RunCLILocal("ps -p 1 -o comm=")
	if err != nil {
		return "", fmt.Errorf("getting init > %v", err)
	}
	if strings.Contains(output, "systemd") {
		return "systemd (cgroup v2)", nil
	}
	return "initd (likely cgroup v1)", nil
}

func getRam(_ ...string) (string, error) {
	output, err := mem.VirtualMemory()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", output.Total/(1024*1024*1024)), nil
}

func getOsUser(_ ...string) (string, error) {
	output, err := user.Current()
	if err != nil {
		return "", err
	}
	return output.Username, nil
}

func getOsType(_ ...string) (string, error) {
	return runtime.GOOS, nil
}

func getOsVersion(_ ...string) (string, error) {
	info, err := host.Info()
	if err != nil {
		return "", err
	}
	return info.PlatformVersion, nil
}

func getOsDistro(_ ...string) (string, error) {
	info, err := host.Info()
	if err != nil {
		return "", err
	}
	return info.Platform, nil
}

func getOsKernelVersion(_ ...string) (string, error) {
	info, err := host.Info()
	if err != nil {
		return "", err
	}
	return info.KernelVersion, nil
}

func getOsFamily(_ ...string) (string, error) {
	info, err := host.Info()
	if err != nil {
		return "", err
	}
	return info.PlatformFamily, nil
}

func getOsArch(_ ...string) (string, error) {
	return runtime.GOARCH, nil
}

func getOsInfos(_ ...string) (string, error) {
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

func getPathTree(params ...string) (string, error) {
	if len(params) < 1 {
		return "", fmt.Errorf("base path name required")
	}

	// get input
	basePath := params[0]

	// play code
	cli := fmt.Sprintf(`find %s -type d | sort | paste -sd\;`, basePath)
	path, err := RunCLILocal(cli)
	if err != nil {
		return "", err
	}
	return path, nil
}

func getPath(_ ...string) (string, error) {
	path := os.Getenv("PATH")
	if path == "" {
		return "", fmt.Errorf("PATH environment variable is not set")
	}
	return path, nil
}

func getPathExtend(params ...string) (string, error) {
	if len(params) < 1 {
		return "", fmt.Errorf("semi-colon separated paths required")
	}

	path := os.Getenv("PATH")
	if path == "" {
		return "", fmt.Errorf("PATH environment variable is not set")
	}
	pathExtend, err := UpdatePath(params[0])
	if err != nil {
		return "", err
	}
	return pathExtend, nil
}

func getSelinuxStatus(_ ...string) (string, error) {
	if selinux.GetEnabled() {
		return "enabled", nil
	}
	return "disabled", nil
}

func getSelinuxMode(_ ...string) (string, error) {
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

func getSelinuxInfos(_ ...string) (string, error) {
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

func getUuid(_ ...string) (string, error) {
	cmd := "sudo cat /sys/class/dmi/id/product_uuid"
	output, err := RunCLILocal(cmd)
	if err != nil {
		return "", fmt.Errorf("getUuid failed: %w", err)
	}
	return output, nil
}

func getCgroupVersion(_ ...string) (string, error) {
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

func getNetIp(_ ...string) (string, error) {
	cmd := "curl -s ifconfig.me -4"
	output, err := RunCLILocal(cmd)
	if err != nil {
		return "", fmt.Errorf("getting net-ip > %v", err)
	}
	return output, nil
}

func getNetGateway(_ ...string) (string, error) {
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
//		value, err := util.GetPropertyLocal(prop)
//		if err != nil {
//			// logx.L.Debugf("%s", err)
//			continue
//		}
//		fmt.Printf("prop: %s value: %s\n", prop, value)
//	}
func GetPropertyLocal(property string, params ...string) (string, error) {
	fn, ok := OsPropertyMap[property]
	if !ok {
		return "", fmt.Errorf("❌ unknown property requested: %s", property)
	}

	output, err := fn(params...)
	if err != nil {
		return "", fmt.Errorf("❌ error getting %s: %w", property, err)
	}

	return output, nil
}

func GetPropertyRemote(vm string, property string, params ...string) (string, error) {
	cli := fmt.Sprintf(`luc util getprop %s`, property)

	// Append optional params if any
	if len(params) > 0 {
		cli = fmt.Sprintf(`luc util getprop %s %s`, property, strings.Join(params, " "))
	}

	out, err := RunCLIRemote(vm, cli)
	if err != nil {
		return "", err
	}
	return out, nil
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
