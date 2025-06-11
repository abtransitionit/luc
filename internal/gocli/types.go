package gocli

import "github.com/abtransitionit/luc/pkg/config"

type PipelineData struct {
	Config       config.CLIConfig // Full config (e.g., luc, helm, etc.)
	GenericUrl   string           // The Generic Url as in the ConfigMap
	SpecificUrl  string           // After placeholders are replaced
	ArtifactName string           // Name of the Artifact
	ArtifactPath string           // Path to saved file
	ArtifactType string           // file or Tgz
	MemoryFile   []byte           // In-memory file content	after curl is successful
	Err          error            // If any step fails
}
