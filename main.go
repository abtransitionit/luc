/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package main

import (
	"github.com/abtransitionit/luc/cmd"
	"github.com/abtransitionit/luc/pkg/logx"
)

func main() {
	// init the logger
	logx.Init(true)
	cmd.Execute()
}
