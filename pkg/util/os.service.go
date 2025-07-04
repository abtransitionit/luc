/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/

package util

import (
	"fmt"
)

func ReloadAndApplyService(action string, listServiceName ...string) error {

	// get property
	osType, err := GetLocalProperty("ostype")
	if err != nil {
		return err
	}

	// manage linux only
	if osType != "linux" {
		return fmt.Errorf("unsupported OS type: %s (only linux is supported)", osType)
	}

	if _, err := RunCLILocal("sudo systemctl daemon-reload"); err != nil {
		// return fmt.Errorf("daemon-reload failed: %w", err)
		return err
	}

	for _, service := range listServiceName {
		command := fmt.Sprintf("sudo systemctl %s %s", action, service)
		if _, err := RunCLILocal(command); err != nil {
			// return fmt.Errorf("%s failed for %s: %w", action, service, err)
			return err
		}
	}

	return nil
}

func RestartService(listServiceName ...string) error {
	return ReloadAndApplyService("restart", listServiceName...)
}

func EnableService(listServiceName ...string) error {
	return ReloadAndApplyService("enable", listServiceName...)
}

func StartService(listServiceName ...string) error {
	return ReloadAndApplyService("start", listServiceName...)
}

func StopService(listServiceName ...string) error {
	return ReloadAndApplyService("stop", listServiceName...)
}

func DisableService(listServiceName ...string) error {
	return ReloadAndApplyService("disable", listServiceName...)
}

func StatusService(listServiceName ...string) (map[string]string, error) {
	results := make(map[string]string)

	// get property
	osType, err := GetLocalProperty("ostype")
	if err != nil {
		return nil, err
	}

	// manage linux only
	if osType != "linux" {
		return nil, fmt.Errorf("unsupported OS type: %s (only linux is supported)", osType)
	}

	for _, service := range listServiceName {
		// Play CLI
		command := fmt.Sprintf("systemctl --no-pager --full status %s", service)
		out, err := RunCLILocal(command)
		if err != nil {
			// return results, fmt.Errorf("status check failed for %s: %w", service, err)
			return nil, err
		}

		// Add to map
		results[service] = out
	}
	return results, nil
}

// func RestartService(listServiceName ...string) error {
// 	for _, serviceName := range listServiceName {
// 		command := fmt.Sprintf("sudo systemctl restart %s", serviceName)
// 		if _, err := RunCLILocal(command); err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }

// func CreateServiceFileRemote(stringContent string, filePath string, vm string) error {
// 	cli := fmt.Sprintf(`luc util oservice cfile %s %s`, stringContent, filePath)
// 	_, err := RunCLIRemote(cli, vm)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

func CreateServiceFile(stringContent string, filePath string) error {

	// get property
	osType, err := GetLocalProperty("ostype")
	if err != nil {
		return err
	}

	// manage only linux
	if osType != "linux" {
		return fmt.Errorf("unsupported OS type: %s (only linux is supported)", osType)
	}

	_, err = SaveStringToFile(stringContent, filePath, true)
	if err != nil {
		return err
	}
	return nil
}
func CreateUserServiceFile(stringContent string, filePath string) error {

	// get property
	osType, err := GetLocalProperty("ostype")
	if err != nil {
		return err
	}

	// manage only linux
	if osType != "linux" {
		return fmt.Errorf("unsupported OS type: %s (only linux is supported)", osType)
	}

	_, err = SaveStringToFile(stringContent, filePath, false)
	if err != nil {
		return err
	}
	return nil
}

// Enable user services to start after a reboot
func EnableUserService(username string) error {
	command := fmt.Sprintf("sudo loginctl enable-linger %s", username)
	_, err := RunCLILocal(command)
	if err != nil {
		return err
		// return fmt.Errorf("failed to enable lingering for user %s: %w", username, err)
	}
	return nil
}
