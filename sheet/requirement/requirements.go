package requirement

// Requirements contains required specifications that must be met to run a
// sheet.
type Requirements struct {
	// Paths is a slice of file or directory paths that are required to run a
	// sheet.
	Paths []Path `json:"paths,omitempty" hcl:"path,block"`
}
