/*
Copyright © 2025 Amar BELGACEM abtransitionit@hotmail.com
*/

// provides user interaction
//
// Features:
//   - read user input from standard input
//   - ask user to choose between continue or stop some function code
//

package ui

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/abtransitionit/luc/pkg/errorx"
	"go.uber.org/zap"
)

// read a line from standard input after displaying a custom or default prompt.
//
// Parameters:
//   - prompt: Message to display before waiting for user input
//
// Returns:
//   - (string, nil): Successful read with trimmed input (may be empty string)
//   - ("", error): Input failure:
//     -- io.ErrUnexpectedEOF if input ends prematurely
//     -- io.EOF if stdin is closed
//     -- Other OS-level reading errors
func ReadUserInput(prompt string) (string, error) {
	if prompt != "" {
		fmt.Print(prompt)
	}

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return errorx.StringError("read user input", prompt, err)
	}

	return strings.TrimSpace(input), nil
}

// Local helper function for ConfirmAction()
func askConfirmAction(prompt string) (bool, error) {
	userInput, err := ReadUserInput(prompt)
	// handle system FAILURE
	if err != nil {
		return errorx.BoolError("read user input", "from stdin", err)
	}

	// handle applogic SUCCESS
	if strings.ToLower(userInput) != "y" && strings.ToLower(userInput) != "yes" {
		return false, nil
	}
	return true, nil
}

//
// Usage:
//   - ui.ConfirmAction(logx.L) // use default prompt
//   - ui.ConfirmAction(logx.L, "do action") // use custom prompt
//
// Note:
//	userPrompt ...string is a workaround to act like optional arguments
//	use helper function : ConfirmAction()

func ConfirmAction(log *zap.SugaredLogger, userPrompt ...string) bool {
	var (
		prompt string = "confirmation required to continue (y/n): " // default prompt
	)

	// manage arguments
	if len(userPrompt) > 0 {
		prompt = fmt.Sprintf("%s : %s", userPrompt[0], prompt)
	}

	log.Debugf("prompt user for confirmation")
	confirmed, err := askConfirmAction(prompt)
	// handle system FAILURE
	if err != nil {
		return false
	}
	// handle applogic SUCCESS
	if !confirmed {
		return false
	} else {
		return true
	}

}
