package runner

import (
	"errors"
	"fmt"
	"github.com/lukasl-dev/ben/sheet/step"
	"os/exec"
	"strings"
	"unicode"
)

func command(stp step.Step, opts Options) error {
	cmd, err := createCmd(stp, opts)
	if err != nil {
		return err
	}
	return runCmd(stp, cmd)
}

// createCmd creates a command to run base on stp.
func createCmd(stp step.Step, opts Options) (*exec.Cmd, error) {
	path, args, err := prepareCmd(stp.Command.Run)
	if err != nil {
		return nil, err
	}
	return &exec.Cmd{
		Path:   path,
		Args:   args,
		Dir:    stp.Command.WorkDir,
		Stdout: opts.Stdout,
		Stderr: opts.Stderr,
	}, nil
}

// prepareCmd prepares the command components from the raw command string.
func prepareCmd(run string) (path string, args []string, err error) {
	if run == "" {
		return "", nil, errors.New("command: no command specified")
	}

	args = splitArgs(run)
	exe := args[0]

	path, err = exec.LookPath(exe)
	if err != nil {
		return "", nil, fmt.Errorf("command: executable %q not found", exe)
	}
	return path, args, nil
}

// splitArgs splits the command string into arguments.
func splitArgs(run string) []string {
	var openingQuote rune
	return strings.FieldsFunc(run, func(r rune) bool {
		end := r == openingQuote
		if end {
			openingQuote = 0
		}
		if r == '"' || r == '\'' || r == '`' {
			openingQuote = r
		}
		return unicode.IsSpace(r) && !end
	})
}

// runCmd runs cmd.
func runCmd(st step.Step, cmd *exec.Cmd) error {
	err := cmd.Run()
	if err == nil {
		return nil
	}

	exitErr, ok := err.(*exec.ExitError)
	if ok && isExitCodeIgnored(exitErr.ExitCode(), st.Command.ExitCodes) {
		return nil
	}
	return err
}

// isExitCodeIgnored checks if the exit code should be considered as successful.
func isExitCodeIgnored(code int, ignored []int) bool {
	for _, i := range ignored {
		if i == code {
			return true
		}
	}
	return false
}
