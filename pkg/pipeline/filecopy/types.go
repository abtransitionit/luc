/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package filecopy

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
func (p PipelineData) String() string {
	t := table.NewWriter()
	t.SetStyle(table.StyleLight)
	t.AppendHeader(table.Row{"Field", "Value"})

	t.AppendRows([]table.Row{
		{"Vm name", p.Node},
		{"Src file", p.SrcFile},
		{"Dst file", p.DstFile},
		{"Nb nodes", p.NbNode},
		{"Error", func() string {
			if p.Err != nil {
				return p.Err.Error()
			}
			return "-"
		}()},
	})

	return t.Render()
}
