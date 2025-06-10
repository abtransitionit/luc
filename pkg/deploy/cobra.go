/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package deploy

import (
	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/abtransitionit/luc/pkg/ui"
	"github.com/spf13/cobra"
)

// # Purpose
//
// Provides shared functionality for executing different "phases" of a process.
// Support for the following features:
//   - Listing available phases (--list flag)
//   - Batch running all phases (--runall flag)
//   - Running a single phase
//   - User confirmation for batch operations
//   - Consistent logging and error handling
//
// # Parameters
//
//   - phases:    Slice defining available phases and their functions
//   - initSDesc: Short description to log at execution start
//
// # Error handling
//
//   - Validates argument count (exactly 1 required for single phase)
//   - Validates phase existence
//   - Terminates execution on phase failure with error logging
//
// # Example usage
//
//	cmd := &cobra.Command{
//	    Run: SharedRun(phases, "Process initialization"),
//	}
func SharedRun(phases []Phase, initSDesc string) func(*cobra.Command, []string) {
	return func(cmd *cobra.Command, args []string) {
		logx.L.Info(initSDesc)

		// handle --list flag before checking arguments
		if cmd.Flags().Changed("list") {
			logx.L.Infof("👉 list all phase name")
			ListPhases(phases)
			return
		}

		// handle --runall flag before checking arguments
		if cmd.Flags().Changed("runall") {
			logx.L.Infof("👉 Running all phase in batch")

			// action must be confirm
			if !forceFlag {
				confirmed := ui.ConfirmAction(logx.L)
				if !confirmed {
					logx.L.Infof("action canceled or system failure")
					return
				} else {
					// log user choice
					logx.L.Infof("User confirmed action")
				}
			}

			// log if force flag is used
			if forceFlag {
				logx.L.Infof("Force flag used")
			}

			// confirmation OK or flag --force
			for _, phase := range phases {
				logx.L.Infof("👉 Running  phase %s", phase.Name)
				// handle system FAILURE - when playing the phase
				if err := phase.Func(); err != nil {
					logx.L.Errorf("❌ Phase %s failed: %v", phase.Name, err)
					return
				}
				// handle applogic SUCCESS
				logx.L.Infof("✅ Phase %s succeded.", phase.Name)
			}
			return
		}

		// manage argument
		if len(args) == 0 {
			cmd.Help()
			return
		}
		if len(args) != 1 {
			logx.L.Debugf("❌ several or no argument provided. only 1 argument is required: the phase name.")
			return
		}
		cmdName := cmd.Name()
		phaseName := args[0]
		logx.L.Debugf("cmd '%s' provide phase name '%s' as argument", cmdName, phaseName)

		// play the code of the phase provided as argument
		for _, phase := range phases {
			if phase.Name == phaseName {
				logx.L.Infof("👉 Running phase: %s", phase.Name)
				// handle system FAILURE - when playing the phase code
				if err := phase.Func(); err != nil {
					logx.L.Errorf("❌ Phase %s failed: %v", phase.Name, err)
					return
				}
				// handle applogic SUCCESS
				logx.L.Infof("✅ Phase %s succeded.", phase.Name)
				return
			}
		}
		// Here the phaseName provided as argument is not a phase name
		// logx.L.Debugf("argument '%s' is not a phase name. Maybe it as a managed argument", phaseName)
		logx.L.Debugf("❌ argument (%s) is not a phase name.", phaseName)
	}
}

var forceFlag bool

// # Purpose
//
// Shared function that is intended to be reused across different commands.
// Sets up common command-line flags for commands using Cobra.
//
// It adds the following flags to the given Cobra command:
//   - --force, -f: A boolean flag to bypass confirmation prompts (default: false).
//   - --list,  -l: A boolean flag to list all available phases (default: false).
//   - --runall, -r: A boolean flag to run all phases sequentially in batch mode (default: false).
//
// # Parameters
//
//	cmd *cobra.Command: The Cobra command to which these shared flags should be added.
//
// # Example usage
//
//	func init() {
//	    deploy.SharedInit(myCommand)
//	}
func SharedInit(cmd *cobra.Command) {
	cmd.Flags().BoolVarP(&forceFlag, "force", "f", false, "Bypass confirmation")
	cmd.Flags().BoolP("list", "l", false, "List all available phases")
	cmd.Flags().BoolP("runall", "r", false, "Run all phases in sequence in batch mode")
	// Make them mutually exclusive
	cmd.MarkFlagsMutuallyExclusive("list", "runall")
}
