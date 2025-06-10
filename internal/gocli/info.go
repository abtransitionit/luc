/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package gocli

import (
	"fmt"

	"github.com/abtransitionit/luc/pkg/logx"
)

const InfoDescription = "infos about the CLI."

func info(arg ...string) error {
	logx.L.Info(InfoDescription)
	location := "/usr/local/bin/luc"
	version := "0.0.1"
	doc := "https://github.com/abtransitionit/luc"
	git := "https://github.com/abtransitionit/luc"
	fmt.Printf("🔹 CLI is available at %s (version: %s)\n", location, version)
	fmt.Printf("🔹 Visit the official docs: %s\n", doc)
	fmt.Printf("🔹 Visit the official git: %s\n", git)
	return nil
}
