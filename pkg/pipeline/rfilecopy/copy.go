/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package rfilecopy

import (
	"fmt"

	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/abtransitionit/luc/pkg/util"
)

// # Purpose
//
// - This stage (like all other stage) get each instance in the piepeline, one by one
// - processes the instance, and send the result to the next pipeline stage
// - copies the file to the remote VM
// - Uses `scp` assuming SSH config is already set up
func copyFile(in <-chan PipelineData, out chan<- PipelineData) {
	defer close(out)

	for data := range in {
		if data.Err != nil {
			out <- data
			logx.L.Debugf("❌ Previous error detected")
			continue
		}

		cmd := fmt.Sprintf("scp %s %s:%s", data.SrcFile, data.Node, data.DstFile)
		_, err := util.RunCLILocal(cmd)
		if err != nil {
			data.Err = err
			logx.L.Debugf("❌ error detected")
		} else {
			logx.L.Infof("[%s] ✅ Copied successfully", data.Node)
		}

		out <- data
	}
}
