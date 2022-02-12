package handler

// Error represents an error that suggests a way to fix the error.
type Error struct {
	// Err is the actual error.
	Err error `json:"error,omitempty"`

	// Title is the short title of the error.
	Title string `json:"title,omitempty"`

	// Suggestions is a slice of suggestions to fix the error.
	Suggestions []string `json:"suggestions,omitempty"`
}

// Error returns s's error message.
func (s *Error) Error() string {
	return s.Err.Error()
}
