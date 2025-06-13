/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package phase

import (
	"os"

	"github.com/jedib0t/go-pretty/table"
)

// denotes/represents a "STEP" in a deployment, provisioning process.
type Phase struct {
	Name        string
	Func        func(arg ...string) (string, error) // go function, code, that perfomrs actions and accepts 0..N arguments
	Description string
	// ExecuteFunc func(ctx context.Context) error // Reserved for future contextual execution.
}

func ShowPhase(setPhase []Phase) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)

	// Simple header
	t.AppendHeader(table.Row{"A", "B", "C"})

	// Simple header

	// Add rows
	for name, cfg := range setPhase {
		t.AppendRow(table.Row{
			name,
			cfg.Description,
			cfg.Name,
		})
	}

	// Render with default style
	t.Render()
}
