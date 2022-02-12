package step

// Remove represents a step that removes a file or directory.
type Remove struct {
	// Remove is the path to the file to remove.
	Remove string `json:"remove,omitempty"`
}
