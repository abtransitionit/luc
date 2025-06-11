/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package gocli

import "github.com/abtransitionit/luc/pkg/logx"

const DstDescription = "create final destination folder for the binary."

func dst(arg ...string) (string, error) {
	logx.L.Info(DstDescription)
	return "", nil
}
