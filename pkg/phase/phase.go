/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package phase

import "fmt"

// represents a set of deployment phases to be run in sequence.
type Deployment struct {
	Phases []Phase // Phases is the ordered list of steps in the deployment process.
}

// creates and returns a new Phase with the given name, function, and description.
// This is a convenience constructor for building Phase instances in a clear and concise way.
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

// Assuming phase.Phase has a Name field or method to get the phase name.
// If the structure is different, adjust accordingly.
// func ListPhases(phases []Phase) {
// 	for i, phase := range phases {
// 		fmt.Printf("%d - %9s - %s\n", i+1, phase.Name, phase.Description)
// 	}
// }

func ListPhases(phases []Phase) {
	// Header
	fmt.Println("ID   Name      Description")
	fmt.Println("---  --------  -----------")

	// Rows
	for i, phase := range phases {
		fmt.Printf("%-3d  %-10s  %s\n",
			i+1,
			phase.Name,
			phase.Description)
	}
}
