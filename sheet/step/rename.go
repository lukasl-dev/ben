package step

type Rename struct {
	// Old is the path to the file or folder to rename.
	Old string `json:"old,omitempty"`

	// New is the new path to the file or folder.
	New string `json:"new,omitempty"`
}
