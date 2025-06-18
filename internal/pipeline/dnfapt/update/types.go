package update

import "github.com/jedib0t/go-pretty/table"

type PipelineData struct {
	HostType              string // Vm or container
	OsFamily              string // Rhel, Debian, dedora
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
func (p PipelineData) String() string {
	t := table.NewWriter()
	t.SetStyle(table.StyleLight)
	t.AppendHeader(table.Row{"Field", "Value"})

	t.AppendRows([]table.Row{
		{"HostType", p.HostType},
		{"OS Family", p.OsFamily},
		{"OS Distro", p.OsDistro},
		{"Os Version", p.OsVersion},
		{"Kernel Version (Before)", p.OskernelVersionBefore},
		{"Kernel Version (After)", p.OskernelVersionAfter},
		{"Need Reboot", p.NeedReboot},
		{"Error", func() string {
			if p.Err != nil {
				return p.Err.Error()
			}
			return "-"
		}()},
	})

	return t.Render()
}
