/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package kind

import (
	"github.com/abtransitionit/luc/internal/config"
	"github.com/abtransitionit/luc/pkg/logx"
	lingerP "github.com/abtransitionit/luc/pkg/pipeline/linger"
)

const LingerDescription = "Allow non root user to run OS services."

func linger(arg ...string) (string, error) {
	_, err := lingerP.RunPipeline(config.KindVm)
	if err != nil {
		logx.L.Debugf("%s", err)
		return "", err
	}
	return "", nil
}

// // Enable lingering for the user
// logx.L.Debugf("enabling lingering for user %s", osUser)
// if err := util.EnableUserService(osUser); err != nil {
// 	logx.L.Debugf("❌ Error detected 6")
// 	return "", err
// }
// // logx.L.Infof("lingering enabled for user %-5s", osUser)
