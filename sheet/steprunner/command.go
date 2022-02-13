package steprunner

import (
	"fmt"
	"os/exec"
	"strings"
	"unicode"

	"github.com/lukasl-dev/ben/sheet"
	"github.com/lukasl-dev/ben/sheet/step"
)

// Command runs a command step.
func Command(sheet sheet.Sheet, base step.Base, cmd step.Command) error {
	if cmd.Command.Run == "" {
		return fmt.Errorf("step: %s: command must not be empty", base.Name)
	}

	split := args(cmd.Command.Run)
	if len(split) == 0 {
		return fmt.Errorf("step: %s: command must not be empty", base.Name)
	}

	path, err := exec.LookPath(split[0])
	if err != nil {
		return fmt.Errorf("step: %s: executable not found", base.Name)
	}

	workDir := sheet.WorkDir
	if cmd.Command.WorkDir != "" {
		workDir = cmd.Command.WorkDir
	}

	c := &exec.Cmd{
		Path: path,
		Args: split,
		Dir:  workDir,
	}

	err = c.Run()
	if err != nil {
		exit, ok := err.(*exec.ExitError)

		// ignore failed exit if it is considered as successful in the
		// given step's configuration
		if ok && isInExitCodes(exit.ExitCode(), cmd.Command.ExitCodes) {
			return nil
		}

		return fmt.Errorf("step: %s: %w", base.Name, err)
	}

	return nil
}

// args splits s into command arguments.
func args(s string) []string {
	quoted := false
	return strings.FieldsFunc(s, func(r rune) bool {
		if r == '"' || r == '\'' || r == '`' {
			quoted = !quoted
		}
		return unicode.IsSpace(r) && !quoted
	})
}

// isInExitCodes returns true if exitCode is found in exitCodes.
func isInExitCodes(exitCode int, exitCodes []int) bool {
	for _, ec := range exitCodes {
		if exitCode == ec {
			return true
		}
	}
	return false
}
