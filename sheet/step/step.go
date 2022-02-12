package step

import (
	"errors"
)

// Step represents a general step in a configuration file. Only one of the
// embedded fields can be set.
type Step struct {
	Base
	*Command
	*Copy
}

// Validate validates s.
func (s Step) Validate() error {
	if s.Command != nil && s.Copy != nil {
		return errors.New("step: a step cannot inherit multiple tasks")
	}
	if s.Command == nil && s.Copy == nil {
		return errors.New("step: a step must inherit one task")
	}
	return nil
}
