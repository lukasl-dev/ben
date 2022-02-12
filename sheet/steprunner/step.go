package steprunner

import (
	"fmt"
	"github.com/lukasl-dev/ben/sheet/step"
)

// Step runs a step according to its type.
func Step(step step.Step) error {
	switch {
	case step.Command != nil:
		return Command(step.Base, *step.Command)
	case step.Copy != nil:
		return Copy(step.Base, *step.Copy)
	case step.Remove != nil:
		return Remove(step.Base, *step.Remove)
	case step.Rename != nil:
		return Rename(step.Base, *step.Rename)
	default:
		return fmt.Errorf("step: %s: unsupported or invalid task specified", step.Base.Name)
	}
}
