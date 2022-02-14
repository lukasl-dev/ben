package steprunner

import (
	"errors"
	"os/exec"
	"strings"
	"unicode"

	"github.com/lukasl-dev/ben/internal"
	"github.com/lukasl-dev/ben/sheet"
	"github.com/lukasl-dev/ben/sheet/step"
)

// command runs a command step.
func command(sh sheet.Sheet, st step.Step) error {
	cmd, err := cmd(sh, st)
	if err != nil {
		return err
	}
	return runCmd(st, cmd)
}

// cmd creates an executable command from st.
func cmd(sh sheet.Sheet, st step.Step) (*exec.Cmd, error) {
	workDir := sh.WorkDir
	if st.Command.WorkDir != "" {
		workDir = st.Command.WorkDir
	}

	path, args, err := lookUpCmdPath(st)
	if err != nil {
		return nil, err
	}

	return &exec.Cmd{Path: path, Args: args, Dir: workDir}, nil
}

// lookupCmdPath looks up the command path and arguments from st.
func lookUpCmdPath(st step.Step) (string, []string, error) {
	if st.Command.Run == "" {
		return "", nil, &internal.Error{
			Prefix: "step",
			Origin: st.Name,
			Err:    errors.New("a command must have a run value"),
			Suggestions: []string{
				"Define a run value in the command step's configuration",
			},
		}
	}

	args := cmdArgs(st.Command.Run)
	if len(args) == 0 {
		return "", nil, &internal.Error{
			Prefix: "step",
			Origin: st.Name,
			Err:    errors.New("a command must have a run value"),
			Suggestions: []string{
				"Define a run value in the command step's configuration",
			},
		}
	}

	path, err := exec.LookPath(args[0])
	if err != nil {
		return "", nil, &internal.Error{
			Prefix: "step",
			Origin: st.Name,
			Err:    err,
			Suggestions: []string{
				"Check if the command is available on your system",
			},
		}
	}

	return path, args, nil
}

// cmdArgs splits s into command arguments.
func cmdArgs(s string) []string {
	quoted := false
	return strings.FieldsFunc(s, func(r rune) bool {
		if r == '"' || r == '\'' || r == '`' {
			quoted = !quoted
		}
		return unicode.IsSpace(r) && !quoted
	})
}

// runCmd runs cmd and returns an error if the exit code is not considered
// as successfully.
func runCmd(st step.Step, cmd *exec.Cmd) error {
	err := cmd.Run()
	if err == nil {
		return nil
	}

	exit, ok := err.(*exec.ExitError)
	if ok && isInExitCodes(exit.ExitCode(), st.Command.ExitCodes) {
		return nil
	}

	return &internal.Error{
		Prefix: "step",
		Origin: st.Name,
		Err:    err,
		Suggestions: []string{
			"Check if the arguments given to the command are correct",
			"Check if the exit code is 0. If not, you can specify the exit codes in the command step's configuration",
		},
	}
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
