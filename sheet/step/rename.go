package step

// Rename represents a step that renames files or directories.
type Rename struct {
	Rename RenameData `json:"rename,omitempty"`
}

// RenameData represents the type of Rename's Rename field.
type RenameData struct {
	// Old is the path to the file or folder to rename.
	Old string `json:"old,omitempty"`

	// New is the new path to the file or folder.
	New string `json:"new,omitempty"`
}
