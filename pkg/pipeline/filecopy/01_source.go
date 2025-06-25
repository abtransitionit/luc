/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package filecopy

import (
	"strings"

	"github.com/abtransitionit/luc/pkg/logx"
)

// # Purpose
//
// - This stage creates a PipelineData instance for each VM
// - e.g 6 VM in the vmMap => 6 instances of the structure PipelineData
// - It sends (out chan<-) each instance into the output channel
func source(out chan<- PipelineData, vmList string, srcFile, dstFile string) {
	// close channel when this code ended
	// closing it make it available for next stage
	// because it is defined outside
	defer close(out)

	// log information
	logx.L.Debugf("defining data to be pipelined")
	logx.L.Debugf("Preparing file copy pipeline for %d VM(s)", len(vmList))

	// loop over all items in the list
	vms := strings.Fields(vmList) // convert ListAsString to slice
	for _, vm := range vms {
		// create a new instance per item (VM)
		data := PipelineData{
			Node:    vm,
			SrcFile: srcFile,
			DstFile: dstFile,
			NbNode:  len(vmList),
		}

		// log information
		logx.L.Infof("[%s] data to be pipelined defined", vm)
		// send the instance to the channel (for next stage or final step)
		out <- data
	}
}
