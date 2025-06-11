/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package gocli

import (
	"github.com/abtransitionit/luc/pkg/logx"
)

const MvDescription = "move the binary to final destination."

func mv(arg ...string) (string, error) {
	logx.L.Info(MvDescription)
	return "", nil
}
