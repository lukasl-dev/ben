package internal

import "fmt"

// Error is a custom error type that can be used to return a more descriptive
// error to the user.
type Error struct {
	// Err is the actual error that occurred.
	Err error `json:"error,omitempty"`

	// Prefix is the prefix to be prepended to the error message.
	Prefix string `json:"prefix,omitempty"`

	// Origin is the origin of the error. This is usually the name of the
	// process that caused the error.
	Origin string `json:"origin,omitempty"`

	// Message is a detailed description of the error. It gives the user a
	// deeper insight into the problem.
	Message string `json:"message,omitempty"`

	// Suggestions is a slice of suggestions that can be used to fix the
	// error.
	Suggestions []string `json:"suggestions,omitempty"`
}

// Error returns err.Err's message prefixed with err.Prefix and err.Origin.
func (err *Error) Error() string {
	return fmt.Sprintf("%s: %s: %s", err.Prefix, err.Origin, err.Err)
}
