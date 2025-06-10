/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package kind

import (
	"github.com/abtransitionit/luc/pkg/logx"
)

const KubectlDescription = "provision the Kubectl CLI"

func kubectl(arg ...string) error {
	logx.L.Info(KubectlDescription)
	// Actual implementation would go here
	return nil
}
