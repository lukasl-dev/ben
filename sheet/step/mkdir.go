package step

// Mkdir represents a Step that creates a directory.
type Mkdir struct {
	// Mkdir is the path to the directory to create.
	Mkdir string `json:"mkdir,omitempty"`
}
