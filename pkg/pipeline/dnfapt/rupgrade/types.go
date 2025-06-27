/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package rupgrade

import "github.com/jedib0t/go-pretty/table"

type PipelineData struct {
	HostName              string //
	HostType              string // Vm or container
	OsFamily              string // Rhel, Debian, fedora
	OsDistro              string // ubuntu, centos, alma, ...
	OsVersion             string //
	OskernelVersionBefore string //
	OskernelVersionAfter  string //
	Err                   error  // error if any
}

// # Pupose
//
// Pretty print the Pipelined Data (usually for debugging)
func (obj PipelineData) String() string {
	t := table.NewWriter()
	t.SetStyle(table.StyleLight)
	t.SetTitle("VM OS dnfapt upgrade Status")
	t.AppendHeader(table.Row{"Field", "Value"})

	t.AppendRows([]table.Row{
		{"Host Name", obj.HostName},
		{"Host Type", obj.HostType},
		{"OS Family", obj.OsFamily},
		{"OS Distro", obj.OsDistro},
		{"Os Version", obj.OsVersion},
		{"Kernel Version (Before)", obj.OskernelVersionBefore},
		{"Kernel Version (After)", obj.OskernelVersionAfter},
		{"Error", func() string {
			if obj.Err != nil {
				return obj.Err.Error()
			}
			return "-"
		}()},
	})

	return t.Render()
}
