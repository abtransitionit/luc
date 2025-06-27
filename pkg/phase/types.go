/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package phase

import (
	"github.com/jedib0t/go-pretty/table"
)

// denotes/represents a "STEP" in a deployment, provisioning process.
type Phase struct {
	Name        string
	Func        func(arg ...string) (string, error) // go function, code, that perfomrs actions and accepts 0..N arguments
	Description string
	// ExecuteFunc func(ctx context.Context) error // Reserved for future contextual execution.
}

// mandatory to implement pretty print
type PhaseList []Phase

// # Purpose
//
// pretty print
//
// # Usage
//
// fmt.Println(PhaseList(ASliceOfPhases))
func (obj PhaseList) String() string {
	t := table.NewWriter()
	t.SetStyle(table.StyleLight)
	t.SetTitle("List of Phases")
	t.AppendHeader(table.Row{"ID", "Name", "Description"})

	for i, phase := range obj {
		t.AppendRow(table.Row{i + 1, phase.Name, phase.Description})
	}

	return t.Render()
}

// func ShowPhase(phases []Phase) {
// 	t := table.NewWriter()
// 	t.SetOutputMirror(os.Stdout)
// 	t.SetStyle(table.StyleLight)
// 	t.Style().Title.Align = text.AlignCenter

// 	t.SetTitle("List of phases")
// 	t.AppendHeader(table.Row{"ID", "Name", "Description"})

// 	for i, phase := range phases {
// 		t.AppendRow(table.Row{
// 			i + 1,
// 			phase.Name,
// 			phase.Description,
// 		})
// 	}

// 	t.Render()
// }

// creates and returns a new Phase with the given name, function, and description.
// This is a convenience constructor for building Phase instances.
//
// Parameters:
//   - name: A phase ID.
//   - fn:   The function attached to the phase. It must return an error.
//   - desc: The purpose of the phase.
//
// Returns:
//
//	A Phase struct with all fields initialized:
//	  - Name is set to the provided name.
//	  - Func is set to the provided function.
//	  - Description is set to the provided description.
//
// Example usage:
//
//	import "github.com/abtransitionit/luc/pkg/deploy"
//	p := phase.SetPhase("check", checkFunc, "Check system health before deployment")
func SetPhase(name string, fn func(cmd ...string) (string, error), desc string) Phase {
	return Phase{Name: name, Func: fn, Description: desc}
}
