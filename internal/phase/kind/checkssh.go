/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package kind

import (
	"github.com/abtransitionit/luc/internal/config"
	"github.com/abtransitionit/luc/pkg/util"
)

func checkSsh(arg ...string) (string, error) {
	util.CheckSsh(config.KindVm)
	return "", nil
}
