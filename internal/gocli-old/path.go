/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package gocli

import (
	"github.com/abtransitionit/luc/pkg/logx"
)

const PathDescription = "Update the CLI path."

func pathx(arg ...string) (string, error) {
	logx.L.Info(PathDescription)
	return "", nil
}
