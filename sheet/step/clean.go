package step

// Clean represents a step that cleans a directory.
type Clean struct {
	Clean CleanData `json:"clean,omitempty"`
}

// CleanData represents the type of Clean's Clean field.
type CleanData struct {
	// Path is the path to the directory to clean.
	Path string `json:"path,omitempty"`

	// Exclude contains a list of files and directories to
	// exclude from the clean.
	Exclude []string `json:"exclude,omitempty"`
}
