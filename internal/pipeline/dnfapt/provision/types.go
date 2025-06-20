package provision

import "github.com/jedib0t/go-pretty/table"

type PipelineData struct {
	Name     string // The package name
	HostType string // Vm or container
	OsFamily string // Rhel, Debian, fedora
	Err      error  // error if any
}

// # Pupose
//
// Pretty print the Pipelined Data (usually for debugging)
func (p PipelineData) String() string {
	t := table.NewWriter()
	t.SetStyle(table.StyleLight)
	t.AppendHeader(table.Row{"Field", "Value"})

	t.AppendRows([]table.Row{
		{"Package", p.Name},
		{"HostType", p.HostType},
		{"OS Family", p.OsFamily},
		{"Error", func() string {
			if p.Err != nil {
				return p.Err.Error()
			}
			return "-"
		}()},
	})

	return t.Render()
}
