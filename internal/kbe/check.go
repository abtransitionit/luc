/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package kbe

import (
	"github.com/abtransitionit/luc/pkg/logx"
)

const CheckDescription = "check basic metrics before starting deployment"

func check(arg ...string) error {
	logx.L.Info(CheckDescription)
	// Actual implementation would go here
	return nil
}
