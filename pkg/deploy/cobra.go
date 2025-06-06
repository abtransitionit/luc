/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package deploy

import (
	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/abtransitionit/luc/pkg/ui"
	"github.com/spf13/cobra"
)

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
		phaseName := args[0]
		logx.L.Debugf("phase name provided as argument: '%s'", phaseName)

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

// SetupCommonFlags configures flags that are shared across commands
func SharedInit(cmd *cobra.Command) {
	cmd.Flags().BoolVarP(&forceFlag, "force", "f", false, "Bypass confirmation")
	cmd.Flags().BoolP("list", "l", false, "List all available phases")
	cmd.Flags().BoolP("runall", "r", false, "Run all phases in sequence in batch mode")
	// Make them mutually exclusive
	cmd.MarkFlagsMutuallyExclusive("list", "runall")
}
