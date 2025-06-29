package packagex

import "github.com/jedib0t/go-pretty/table"

type PipelineData struct {
	HostName       string // The package name
	PackageName    string // The package name
	PackageVersion string // The package version
	HostType       string // Vm or container
	OsDistro       string
	OsFamily       string // Rhel, Debian, fedora
	OsVersion      string
	Err            error // error if any
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
		{"package Name", obj.PackageName},
		{"Host Type", obj.HostType},
		{"OS Family", obj.OsFamily},
		{"Error", func() string {
			if obj.Err != nil {
				return obj.Err.Error()
			}
			return "-"
		}()},
	})

	return t.Render()
}
