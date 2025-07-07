/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package path

import (
	"os"
	"sort"
	"strings"

	"github.com/jedib0t/go-pretty/table"
)

type PipelineData struct {
	HostName string
	Path     string
	FilePath string
	Err      error
}

// # Pupose
//
// Pretty print the Pipelined Data (usually for debugging)
func (obj PipelineData) String() string {
	t := table.NewWriter()
	t.SetStyle(table.StyleLight)
	t.AppendHeader(table.Row{"Field", "Value"})

	paths := strings.Split(obj.Path, string(os.PathListSeparator))
	sort.Strings(paths) // Sort alphabetically

	t.AppendRows([]table.Row{
		{"Host name", obj.HostName},
		// {"PATH", obj.Path},
		{"PATH", strings.Join(paths, "\n")},
		{"File PATH", obj.FilePath},
		{"Error", func() string {
			if obj.Err != nil {
				return obj.Err.Error()
			}
			return "-"
		}()},
	})

	return t.Render()
}
