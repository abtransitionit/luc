/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package rfilecopy

import (
	"github.com/jedib0t/go-pretty/table"
)

type PipelineData struct {
	Node    string // hostname of the VM
	SrcFile string // local absolute path to the file
	DstFile string // remote folder where to copy
	NbNode  int    // remote folder where to copy
	Err     error  // error in the pipeline step
}

// # Pupose
//
// Pretty print the Pipelined Data (usually for debugging)
func (obj PipelineData) String() string {
	t := table.NewWriter()
	t.SetStyle(table.StyleLight)
	t.SetTitle("LUC remote copy status")
	t.AppendHeader(table.Row{"Field", "Value"})

	t.AppendRows([]table.Row{
		{"Vm name", obj.Node},
		{"Src file", obj.SrcFile},
		{"Dst file", obj.DstFile},
		{"Nb nodes", obj.NbNode},
		{"Error", func() string {
			if obj.Err != nil {
				return obj.Err.Error()
			}
			return "-"
		}()},
	})

	return t.Render()
}
