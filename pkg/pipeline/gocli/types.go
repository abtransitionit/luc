package gocli

import (
	"github.com/abtransitionit/luc/pkg/config"
	"github.com/jedib0t/go-pretty/table"
)

type PipelineData struct {
	Config       config.CLIConfig // Full config (e.g., luc, helm, etc.)
	GenericUrl   string           // Url as it appears in the ConfigMap
	SpecificUrl  string           // Url after placeholders are replaced
	MemoryFile   []byte           // In-memory file content	after curl is successful
	ArtifactName string           // artifact name as it appears in the specific URL
	ArtifactPath string           // Path to saved artifact on the host FS after it is curled
	FofTmpPath   string           // file Or folder Path after CLI artefact is unzip (Tgz, Exe) or git clone (Git)
	ArtifactType string           // guessed filetype : Exe or Tgz
	CliName      string           //
	Version      string           //
	DstFolder    string           //
	Err          error            // If any step fails
}

// # Pupose
//
// Pretty print the Pipelined Data (usually for debugging)
func (p PipelineData) String() string {
	t := table.NewWriter()
	t.SetStyle(table.StyleLight)
	t.AppendHeader(table.Row{"Field", "Value"})

	t.AppendRows([]table.Row{
		{"gocli name", p.Config.Name},
		{"CLI version", p.Version},
		{"Generic Url", p.GenericUrl},
		{"Specific Url", p.SpecificUrl},
		{"Artifact Name", p.ArtifactName},
		{"Artifact Guessed Type", p.ArtifactType},
		{"Artifact Path", p.ArtifactPath},
		{"Artifact FofTmpPath", p.FofTmpPath},
		{"Dst Folder", p.DstFolder},
		{"Error", func() string {
			if p.Err != nil {
				return p.Err.Error()
			}
			return "-"
		}()},
	})

	return t.Render()
}
