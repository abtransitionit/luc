/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package phase

import (
	"errors"
	"fmt"
	"os"

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
//	    Run: CmdRun(phases, "Process initialization"),
//	}
func CmdRun(phases []Phase, initSDesc string) func(*cobra.Command, []string) {
	return func(cmd *cobra.Command, args []string) {

		// handle flag --show before checking arguments
		if handled, err := handleShowFlag(cmd, phases); handled {
			if err != nil {
				logx.L.Debugf("❌ system failure : %s", err)
			}
			return
		}

		// handle --runall flag before checking arguments
		if handled, err := handleRunAllFlag(cmd, phases); handled {
			if err != nil {
				logx.L.Debugf("❌ system failure : %s", err)
			}
			return
		}

		// manage CMD argument
		if len(args) == 0 {
			cmd.Help()
			fmt.Println(PhaseList(phases))
			return
		}
		if len(args) != 1 {
			logx.L.Debugf("❌ several or no argument provided. only 1 argument is required: the phase name.")
			return
		}

		// handle single phase run
		// logx.L.Info(initSDesc)
		if handled, err := handleSinglePhase(cmd, args[0], phases); handled {
			if err != nil {
				return
			}
			return
		}
	}
}

// lists all phases if the --list flag is set
//
// Returns:
//   - (true,  nil) if the flag is set and handled successfully.
//   - (false, nil) if the flag is not set.
func handleShowFlag(cmd *cobra.Command, phases []Phase) (bool, error) {
	if cmd.Flags().Changed("show") {
		logx.L.Debugf("cmd '%s' description is : %s", cmd.Name(), cmd.Short)
		logx.L.Infof("👉 list all phase names")
		// cast phases to PhaseList
		fmt.Println(PhaseList(phases))
		return true, nil
	}
	return false, nil
}

// # Purpose
//
// runs all phases in batch if the --runall flag is set.
// Asks for user confirmation if forceFlag is not set.
//
// # Returns
//   - (true, nil)   if the flag is set and all phases run successfully.
//   - (false, nil)  if the flag is not set.
//   - (true, error) if the user cancels the operation or any phase fails
func handleRunAllFlag(cmd *cobra.Command, phases []Phase) (bool, error) {
	if !cmd.Flags().Changed("runall") {
		return false, nil
	}

	logx.L.Infof("👉 Running all phase in batch")

	// action must be confirm
	if !forceFlag {
		confirmed := ui.ConfirmAction(logx.L)
		if !confirmed {
			logx.L.Debugf("❌ action canceled or system failure")
			return true, errors.New("user canceled operation")
		}
		// Here, the user confirmed the action
		logx.L.Infof("User confirmed action")
	}

	// log if force flag is used
	if forceFlag {
		logx.L.Infof("Force flag used")
	}
	// Here, user confirmed or use force flag - Play all the phases
	for _, phase := range phases {
		logx.L.Infof("👉 Running phase '%s'", phase.Name)
		// handle system FAILURE
		_, err := phase.Func()
		if err != nil {
			logx.L.Debugf("❌ Phase '%s' failed: %v", phase.Name, err)
			os.Exit(1)
			// return true, fmt.Errorf("phase %s failed: %w", phase.Name, err)
			return true, fmt.Errorf("phase %s failed: %w", phase.Name, err)
		}
		// handle applogic SUCCESS
		logx.L.Infof("✅ Phase %s succeeded.", phase.Name)
	}
	return true, nil
}

// Runs a single phase, if provided as argument to the cmd.
//
// Returns:
//   - (true, nil)   if a matching phase was found and run successfully
//   - (false, nil)  if no valid phase name argument was provided
//   - (true, error) if the matching phase failed during execution
func handleSinglePhase(cmd *cobra.Command, phaseName string, phases []Phase) (bool, error) {
	cmdName := cmd.Name()
	logx.L.Debugf("cmd '%s' description is : %s", cmdName, cmd.Short)
	logx.L.Debugf("cmd '%s' argument is    : %s", cmdName, phaseName)

	// play the code of the single phase provided as argument
	for _, phase := range phases {
		if phase.Name == phaseName {
			// logx.L.Infof("👉 Running phase: '%s', that '%s'", phase.Name, phase.Description)
			logx.L.Infof("👉 Running phase: '%s'", phase.Name)
			if _, err := phase.Func(cmdName); err != nil {
				// handle system FAILURE
				logx.L.Debugf("❌ Phase '%s': %v", phase.Name, err)
				return true, err
			}
			// handle applogic SUCCESS
			logx.L.Infof("✅ Phase %s succeeded.", phase.Name)
			return true, nil
		}
	}
	// Here argument provided is not a phase name
	// logx.L.Debugf("argument '%s' is not a phase name. Maybe it as a managed argument", phaseName)
	logx.L.Debugf("❌ argument (%s) is not a phase name.", phaseName)
	return false, nil
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
//	    phase.CmdInit(myCommand)
//	}
func CmdInit(cmd *cobra.Command) {
	cmd.Flags().BoolVarP(&forceFlag, "force", "f", false, "Bypass confirmation")
	cmd.Flags().BoolP("show", "s", false, "List all available phase names")
	cmd.Flags().BoolP("runall", "r", false, "Run all phases in sequence in batch mode")
	// Make them mutually exclusive
	cmd.MarkFlagsMutuallyExclusive("show", "runall")
}

// // handle --list flag before checking arguments
// if cmd.Flags().Changed("list") {
// 	logx.L.Infof("👉 list all phase name")
// ListPhases(phases)
// return
// }

// // handle --runall flag before checking arguments
// if cmd.Flags().Changed("runall") {
// 	logx.L.Infof("👉 Running all phase in batch")

// 	// action must be confirm
// 	if !forceFlag {
// 		confirmed := ui.ConfirmAction(logx.L)
// 		if !confirmed {
// 			logx.L.Infof("action canceled or system failure")
// 			return
// 		} else {
// 			// log user choice
// 			logx.L.Infof("User confirmed action")
// 		}
// 	}

// 	// log if force flag is used
// 	if forceFlag {
// 		logx.L.Infof("Force flag used")
// 	}

// 	// confirmation OK or flag --force
// 	for _, phase := range phases {
// 		logx.L.Infof("👉 Running  phase %s", phase.Name)
// 		// handle system FAILURE - when playing the phase
// 		if _, err := phase.Func(); err != nil {
// 			logx.L.Errorf("❌ Phase %s failed: %v", phase.Name, err)
// 			return
// 		}
// 		// handle applogic SUCCESS
// 		logx.L.Infof("✅ Phase %s succeded.", phase.Name)
// 	}
// 	return
// }

// 	cmdName := cmd.Name()
// 	phaseName := args[0]
// 	logx.L.Debugf("cmd '%s' provide phase name '%s' as argument", cmdName, phaseName)

// 	// play the code of the phase provided as argument
// 	for _, phase := range phases {
// 		if phase.Name == phaseName {
// 			logx.L.Infof("👉 Running phase: %s", phase.Name)
// 			// handle system FAILURE - when playing the phase code
// 			if _, err := phase.Func(cmdName); err != nil {
// 				logx.L.Errorf("❌ Phase %s failed: %v", phase.Name, err)
// 				return
// 			}
// 			// handle applogic SUCCESS
// 			logx.L.Infof("✅ Phase %s succeded.", phase.Name)
// 			return
// 		}
// 	}
// 	// Here the phaseName provided as argument is not a phase name
// 	// logx.L.Debugf("argument '%s' is not a phase name. Maybe it as a managed argument", phaseName)
// 	logx.L.Debugf("❌ argument (%s) is not a phase name.", phaseName)
// }
