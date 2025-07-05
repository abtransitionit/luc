/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package kind

import (
	"github.com/abtransitionit/luc/internal/config"
	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/abtransitionit/luc/pkg/pipeline/oservice"
)

const ServiceDescription = "configure OS services on Kind VMs."

func service(arg ...string) (string, error) {
	_, err := oservice.RunPipeline(config.KindVm, config.KindServiceMap)
	if err != nil {
		logx.L.Debugf("%s", err)
		return "", err
	}
	return "", nil
}

// // define the user service file path for containerd
// containerdServiceFilePath := "/etc/systemd/system/containerd.service"
// // Create the service file for containerd
// if err := util.CreateUserServiceFile(config.ContainerdUserServiceConf, containerdServiceFilePath); err != nil {
// 	logx.L.Debugf("❌ Errod detected 1")
// 	return "", fmt.Errorf("%s", err)
// }
