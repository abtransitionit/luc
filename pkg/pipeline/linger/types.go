/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package linger

import (
	"github.com/jedib0t/go-pretty/table"
)

type PipelineData struct {
	HostName     string
	OsFamily     string // Rhel, Debian, fedora
	osUser       string
	LingerStatus string
	Err          error
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
		{"OS   Family", obj.OsFamily},
		{"user name", obj.osUser},
		{"linger status", obj.LingerStatus},
		{"Error", func() string {
			if obj.Err != nil {
				return obj.Err.Error()
			}
			return "-"
		}()},
	})

	return t.Render()
}
