package update

import "github.com/abtransitionit/luc/pkg/config"

type PipelineData struct {
	Config       config.CLIConfig // Full config (e.g., luc, helm, etc.)
	GenericUrl   string           // Url as it appears in the ConfigMap
	SpecificUrl  string           // Url after placeholders are replaced
	MemoryFile   []byte           // In-memory file content	after curl is successful
	ArtifactName string           // artifact name as it appears in the specific URL
	ArtifactPath string           // Path to saved artifact on the host FS after it is curled
	FofTmpPath   string           // file Or folder Path after CLI artefact is unzip (Tgz, Exe) or git clone (Git)
	ArtifactType string           // guessed filetype : Exe or Tgz
	Err          error            // If any step fails
}
