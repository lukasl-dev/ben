package step

type Clean struct {
	// Path is the path to the directory to clean.
	Path string `json:"path,omitempty"`

	// Exclude contains a list of files and directories to
	// exclude from the clean.
	Exclude []string `json:"exclude,omitempty"`
}
