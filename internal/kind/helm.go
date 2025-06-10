/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package kind

import (
	"github.com/abtransitionit/luc/pkg/logx"
)

const HelmDescription = "provision the Helm client CLI"

func helm(arg ...string) error {
	logx.L.Info(HelmDescription)
	// Actual implementation would go here
	return nil
}
