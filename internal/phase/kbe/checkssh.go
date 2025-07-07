/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package kbe

import (
	"github.com/abtransitionit/luc/internal/config"
	"github.com/abtransitionit/luc/pkg/util"
)

func checkSsh(arg ...string) (string, error) {
	// launch this function
	_, err := util.CheckSsh(config.KbeListNode)
	if err != nil {
		return "", err
	}
	// success
	return "", nil
}
