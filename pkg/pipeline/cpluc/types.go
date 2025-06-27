package cpluc

import "github.com/jedib0t/go-pretty/table"

type PipelineData struct {
	VmName        string // The Vm concerned
	localOutput   string // Output of the local build for the current platform
	localExePath  string // Output of the local build for the current platform
	localOutXptf  string // Output of the local build for the cross platform
	remoteTmpPath string // remote temporary exe path
	remoteExePath string // remote final exe path
	Err           error  // If any step fails
}

func (obj PipelineData) String() string {
	t := table.NewWriter()
	t.SetStyle(table.StyleLight)
	t.SetTitle("LUC remote copy status")
	t.AppendHeader(table.Row{"Field", "Value"})

	t.AppendRows([]table.Row{
		{"Vm Name", obj.VmName},
		{"local Output", obj.localOutput},
		{"local deploy path", obj.localExePath},
		{"local Output for Xptf", obj.localOutXptf},
		{"remote tmp path", obj.remoteTmpPath},
		{"remote deploy path", obj.remoteExePath},
		{"Error", func() string {
			if obj.Err != nil {
				return obj.Err.Error()
			}
			return "-"
		}()},
	})

	return t.Render()
}
