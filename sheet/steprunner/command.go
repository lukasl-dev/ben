package steprunner

import (
	"fmt"
	"github.com/lukasl-dev/ben/sheet/step"
	"os/exec"
	"strings"
)

// CommandOptions holds optional parameter for Command().
type CommandOptions struct {
	// WorkDir is the working directory for the command.
	WorkDir string `json:"workDir,omitempty"`
}

// Command runs a command step.
func Command(base step.Base, cmd step.Command, opts CommandOptions) error {
	if cmd.Command == "" {
		return fmt.Errorf("step: %s: command must not be empty", base.Name)
	}

	split := strings.Split(cmd.Command, " ")
	if len(split) == 0 {
		return fmt.Errorf("step: %s: command must not be empty", base.Name)
	}

	path, err := exec.LookPath(split[0])
	if err != nil {
		return fmt.Errorf("step: %s: executable not found", base.Name)
	}

	c := &exec.Cmd{
		Path: path,
		Args: split,
		Dir:  opts.WorkDir,
	}

	if err := c.Start(); err != nil {
		return fmt.Errorf("step: %s: %w", base.Name, err)
	}

	return nil
}
