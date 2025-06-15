/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package test

import (
	"github.com/abtransitionit/luc/pkg/logx"
)

// Move File or folder to final destination
func End(in <-chan string) {
	logx.L.Infof("Hello from Output")
	for item := range in {
		logx.L.Infof("Output: %s", item)
	}
}
