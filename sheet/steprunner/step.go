package steprunner

import (
	"fmt"
	"github.com/lukasl-dev/ben/sheet/step"
)

// StepOptions holds optional parameter for Step().
type StepOptions struct {
	// Command are the CommandOptions to use.
	Command CommandOptions

	// Copy are the CommandOptions to use.
	Copy CopyOptions
}

// Step runs a step according to its type.
func Step(step step.Step, options StepOptions) error {
	switch {
	case step.Command != nil:
		return Command(step.Base, *step.Command, options.Command)
	case step.Copy != nil:
		return Copy(step.Base, *step.Copy, options.Copy)
	default:
		return fmt.Errorf("step: %s: unsupported or invalid task specified", step.Base.Name)
	}
}
