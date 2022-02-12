package step

// Copy represents a step that copies a file or directory into another location.
type Copy struct {
	// From the path to the file or directory to copy. It can be a URL for remote
	// files.
	From string `json:"from,omitempty"`

	// To the path to the destination.
	To string `json:"to,omitempty"`
}
