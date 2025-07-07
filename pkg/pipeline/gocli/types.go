package gocli

import (
	"github.com/abtransitionit/luc/pkg/config"
	"github.com/jedib0t/go-pretty/table"
)

type PipelineData struct {
	HostName     string           // Full config (e.g., luc, helm, etc.)
	Config       config.CLIConfig // Full config (e.g., luc, helm, etc.)
	GenericUrl   string           // Url as it appears in the ConfigMap
	HostUrl      string           // Url after placeholders are replaced
	MemoryFile   []byte           // In-memory file content	after curl is successful
	ArtName      string           // artifact name as it appears in the specific URL
	ArtPath1     string           // Path to saved artifact on the host FS after it is curled
	ArtPath2     string           // Path to saved artifact on the host FS after it is curled
	ArtifactType string           // guessed filetype : Exe or Tgz
	Version      string           //
	DstFolder    string           //
	Err          error            // If any step fails
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
		{"Go CLI name", obj.Config.Name},
		{"CLI version", obj.Version},
		{"Generic Url", obj.GenericUrl},
		{"Host Url", obj.HostUrl},
		{"Artifact Name", obj.ArtName},
		{"Artifact Tmp Path 1", obj.ArtPath1},
		{"Artifact Tmp Path 2", obj.ArtPath2},
		{"Artifact Guessed Type", obj.ArtifactType},
		{"Dst Folder", obj.DstFolder},
		{"Error", func() string {
			if obj.Err != nil {
				return obj.Err.Error()
			}
			return "-"
		}()},
	})

	return t.Render()
}
