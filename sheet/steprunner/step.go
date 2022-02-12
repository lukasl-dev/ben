package steprunner

import (
	"fmt"
	"github.com/lukasl-dev/ben/sheet/step"
)

// StepOptions holds optional parameter for Step().
type StepOptions struct {
	// Command are the CommandOptions to use.
	Command CommandOptions
}

// Step runs a step according to its type.
func Step(step step.Step, options StepOptions) error {
	switch {
	case step.Command != nil:
		return Command(step.Base, *step.Command, options.Command)
	case step.Copy != nil:
		return Copy(step.Base, *step.Copy)
	case step.Rename != nil:
		return Rename(step.Base, *step.Rename)
	default:
		return fmt.Errorf("step: %s: unsupported or invalid task specified", step.Base.Name)
	}
}
