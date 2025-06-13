/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package phase

// denotes/represents a "STEP" in a deployment, provisioning process.
type Phase struct {
	Name        string
	Func        func(arg ...string) (string, error) // go function, code, that perfomrs actions and accepts 0..N arguments
	Description string
	// ExecuteFunc func(ctx context.Context) error // Reserved for future contextual execution.
}
