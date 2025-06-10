/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package kind

import (
	"github.com/abtransitionit/luc/pkg/logx"
)

const CliDescription = "provision needed CLI"

func cli(arg ...string) error {
	logx.L.Info(CliDescription)
	// Actual implementation would go here
	return nil
}
