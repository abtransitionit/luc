/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package kbe

import (
	"fmt"
	"os"

	"github.com/abtransitionit/luc/internal/config"
	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/abtransitionit/luc/pkg/util"
)

const ServiceDescription = "configure OS services."

func service(arg ...string) (string, error) {
	// filePath := "/tmp/usr.local.bin.rootlesskit.rootlesskit"

	// Create the service file for apparmor
	if err := util.CreateServiceFile(config.ApparmorServiceConf, config.ApparmorFilePath); err != nil {
		logx.L.Debugf("❌ Errod detected 1")
		return "", fmt.Errorf("%s", err)
	}
	// // define the user service file path for containerd
	// containerdServiceFilePath := "/etc/systemd/system/containerd.service"
	// // Create the service file for containerd
	// if err := util.CreateUserServiceFile(config.ContainerdUserServiceConf, containerdServiceFilePath); err != nil {
	// 	logx.L.Debugf("❌ Errod detected 1")
	// 	return "", fmt.Errorf("%s", err)
	// }

	// get property
	osFamily, err := util.OsPropertyGet("osfamily")
	if err != nil {
		logx.L.Debugf("❌ Error detected 2")
		return "", err
	}

	// // get property
	// envPath, err := util.OsPropertyGet("path")
	// if err != nil {
	// 	logx.L.Debugf("❌ Error detected 4")
	// 	return "", err
	// }
	if osFamily == "debian" {
		// Restart the services only for ubuntu
		logx.L.Debugf("restarting service apparmor")
		if err := util.RestartService("apparmor.service"); err != nil {
			logx.L.Debugf("❌ Errod detected 5")
			return "", err
		}
	} else {
		// Do nothing for other OS families.
	}

	// get property
	osUser, err := util.OsPropertyGet("osuser")
	if err != nil {
		logx.L.Debugf("❌ Error detected 3")
		return "", err
	}

	// Enable lingering for the user
	logx.L.Debugf("enabling lingering for user %s", osUser)
	if err := util.EnableUserService(osUser); err != nil {
		logx.L.Debugf("❌ Error detected 6")
		return "", err
	}
	// logx.L.Infof("lingering enabled for user %-5s", osUser)

	// build a tree PATH
	basePath := "/usr/local/bin"
	logx.L.Debugf("building tree path from: '%s'", basePath)
	treePath, err := util.BuildPath(basePath)
	if err != nil {
		logx.L.Debugf("❌ Error detected 7")
		return "", err
	}

	// update PATH with this tree path
	// logx.L.Debugf("parevious $PATH is : '%s'", envPath)
	// logx.L.Debugf("updating envar $PATH with : '%s'", treePath)
	updatedPath, err := util.UpdPath(treePath)
	if err != nil {
		logx.L.Debugf("❌ Error detected 8")
		return "", err
	}

	// set env var
	err = os.Setenv("PATH", updatedPath)
	if err != nil {
		logx.L.Debugf("❌ Error detected 9")
		return "", err
	}

	// Play CLI
	cli := "containerd-rootless-setuptool.sh install"
	_, err = util.RunCLILocal(cli)
	if err != nil {
		logx.L.Debugf("❌ Error detected 10")
		return "", err
	}

	// SUCCESS
	return "Service file created and AppArmor restarted successfully", nil
}
