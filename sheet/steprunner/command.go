package steprunner

import (
	"fmt"
	"github.com/lukasl-dev/ben/sheet"
	"github.com/lukasl-dev/ben/sheet/step"
	"os/exec"
	"strings"
	"unicode"
)

// Command runs a command step.
func Command(sheet sheet.Sheet, base step.Base, cmd step.Command) error {
	if cmd.Command == "" {
		return fmt.Errorf("step: %s: command must not be empty", base.Name)
	}

	split := args(cmd.Command)
	if len(split) == 0 {
		return fmt.Errorf("step: %s: command must not be empty", base.Name)
	}

	path, err := exec.LookPath(split[0])
	if err != nil {
		return fmt.Errorf("step: %s: executable not found", base.Name)
	}

	workDir := sheet.WorkDir
	if cmd.WorkDir != "" {
		workDir = cmd.WorkDir
	}

	c := &exec.Cmd{
		Path: path,
		Args: split,
		Dir:  workDir,
	}

	if err := c.Run(); err != nil {
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
