package runner

import (
	"fmt"
	"github.com/lukasl-dev/ben/sheet/step"
)

// Error is a custom error that occurs during the execution of a step.
type Error struct {
	// Step is the step that caused the error.
	Step step.Step `json:"step,omitempty"`

	// Message is the error message.
	Message interface{} `json:"message,omitempty"`
}

// newError returns a new Error.
func newError(stp step.Step, msg interface{}) *Error {
	return &Error{Step: stp, Message: msg}
}

// Error returns err's message prepended with a prefix.
func (err *Error) Error() string {
	return fmt.Sprintf("runner: %s: %s", err.Step.Name, err.Message)
}
