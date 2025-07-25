/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package reboot

import "github.com/jedib0t/go-pretty/table"

type PipelineData struct {
	HostType              string // Vm or container
	OsFamily              string // Rhel, Debian, fedora
	OsDistro              string // ubuntu, centos, alma, ...
	OsVersion             string //
	OskernelVersionBefore string //
	OskernelVersionAfter  string //
	NeedReboot            string // true or false
	Err                   error  // error if any
}

// # Pupose
//
// Pretty print the Pipelined Data (usually for debugging)
func (obj PipelineData) String() string {
	t := table.NewWriter()
	t.SetStyle(table.StyleLight)
	t.AppendHeader(table.Row{"Field", "Value"})

	t.AppendRows([]table.Row{
		{"HostType", obj.HostType},
		{"OS Family", obj.OsFamily},
		{"OS Distro", obj.OsDistro},
		{"Os Version", obj.OsVersion},
		{"Kernel Version (Before)", obj.OskernelVersionBefore},
		{"Kernel Version (After)", obj.OskernelVersionAfter},
		{"Need Reboot", obj.NeedReboot},
		{"Error", func() string {
			if obj.Err != nil {
				return obj.Err.Error()
			}
			return "-"
		}()},
	})

	return t.Render()
}
