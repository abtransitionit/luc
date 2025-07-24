package repo

import (
	"github.com/abtransitionit/luc/pkg/config"
	"github.com/jedib0t/go-pretty/table"
)

type PipelineData struct {
	HostName              string                  // The VM name
	Config                config.DnfaptRepoConfig // Full config (e.g., luc, helm, etc.)
	CName                 string
	RepositoryList        []string // The list of packages to install.
	HostType              string   // Vm or container
	OsDistro              string
	OsFamily              string // Rhel, Debian, fedora
	OsVersion             string
	OskernelVersionBefore string
	OskernelVersionAfter  string
	RebootStatus          string
	GenericUrlRepo        string
	GenericUrlGpg         string
	UrlRepo               string
	UrlGpg                string
	RepoFilePath          string
	GpgFilePath           string
	Version               string
	Err                   error // error if any
}

// # Pupose
//
// Pretty print
func (obj PipelineData) String() string {
	t := table.NewWriter()
	t.SetStyle(table.StyleLight)
	t.SetTitle("VM dnfapt repository provisioning Status")
	t.AppendHeader(table.Row{"Field", "Value"})

	t.AppendRows([]table.Row{
		{"Host name", obj.HostName},
		{"Host Type", obj.HostType},
		{"OS Family", obj.OsFamily},
		{"OS Distro", obj.OsDistro},
		{"Repo NAME", obj.Config.Name},
		{"Repo CNAME", obj.CName},
		{"Repo version", obj.Version},
		{"Repo Url", obj.GenericUrlRepo},
		{"Repo Url", obj.UrlRepo},
		{"Repo file path", obj.RepoFilePath},
		{"Gpg Url", obj.GenericUrlGpg},
		{"Gpg Url", obj.UrlGpg},
		{"Gpg file path", obj.GpgFilePath},
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
