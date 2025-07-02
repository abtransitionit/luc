package gocli

import (
	"github.com/jedib0t/go-pretty/table"
)

type PipelineData struct {
	HostName    string //
	ServiceName string //
	Err         error  // If any stage fails
}

// # Pupose
//
// Pretty print the Pipelined Data (usually for debugging)
func (obj PipelineData) String() string {
	t := table.NewWriter()
	t.SetStyle(table.StyleLight)
	t.AppendHeader(table.Row{"Field", "Value"})

	t.AppendRows([]table.Row{
		{"Host name", obj.HostName},
		{"service name", obj.ServiceName},
		{"Error", func() string {
			if obj.Err != nil {
				return obj.Err.Error()
			}
			return "-"
		}()},
	})

	return t.Render()
}
