package packagex

import "github.com/jedib0t/go-pretty/table"

type PipelineData struct {
	HostName              string   // The VM name
	PackageList           []string // The list of packages to install.
	HostType              string   // Vm or container
	OsDistro              string
	OsFamily              string // Rhel, Debian, fedora
	OsVersion             string
	OskernelVersionBefore string //
	OskernelVersionAfter  string //
	RebootStatus          string //
	Err                   error  // error if any
}

// # Pupose
//
// Pretty print
func (obj PipelineData) String() string {
	t := table.NewWriter()
	t.SetStyle(table.StyleLight)
	t.SetTitle("VM dnfapt insta Status")
	t.AppendHeader(table.Row{"Field", "Value"})

	t.AppendRows([]table.Row{
		{"Host name", obj.HostName},
		{"Host Type", obj.HostType},
		{"OS Family", obj.OsFamily},
		{"OS Distro", obj.OsDistro},
		{"Package List", obj.PackageList},
		{"Kerner version (before)", obj.OskernelVersionBefore},
		{"Kerner version (after)", obj.OskernelVersionAfter},
		{"Reboot status", obj.RebootStatus},
		{"Error", func() string {
			if obj.Err != nil {
				return obj.Err.Error()
			}
			return "-"
		}()},
	})

	return t.Render()
}
