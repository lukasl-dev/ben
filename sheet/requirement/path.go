package requirement

// Path configures a path requirement.
type Path struct {
	// Path is the path to the file or directory to require.
	Path string `json:"path,omitempty" hcl:"path,label"`

	// Exists reports whether the path must exist or not. If it is nil, the
	// path must exist.
	Exists *bool `json:"exists,omitempty" hcl:"exists,optional"`

	// Dir reports whether the path must be a directory or not. If it is nil,
	// the path must be a directory.
	Dir *bool `json:"dir,omitempty" hcl:"dir,optional"`

	// File reports whether the path must be a file or not. If it is nil, the
	// path must be a file.
	File *bool `json:"file,omitempty" hcl:"file,optional"`
}
