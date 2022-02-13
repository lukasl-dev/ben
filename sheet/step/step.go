package step

// Step represents a general step in a configuration file. Only one of the
// embedded fields can be set.
type Step struct {
	// Name is the name of the step. It is used for logging and representative
	// purposes.
	Name string `json:"name,omitempty"`

	// Clean is the data for cleaning a directory. If it is set, the step is a
	// Clean step.
	Clean *Clean `json:"clean,omitempty"`

	// Command is the data for running a command. If it is set, the step is a
	// Command step.
	Command *Command `json:"command,omitempty"`

	// Copy is the data for copying a file or directory into another localtion.
	// If it is set, the step is a Copy step.
	Copy *Copy `json:"copy,omitempty"`

	// Mkdir is the path to the directory to create. If it is set, the step is a
	// Mkdir step.
	Mkdir *string `json:"mkdir,omitempty"`

	// Remove is th path to the file or directory to remove. If it is set, the
	// step is a Remove step.
	Remove *string `json:"remove,omitempty"`

	// Rename is the data for renaming a file or directory. If it is set, the
	// step is a Rename step.
	Rename *Rename `json:"rename,omitempty"`
}

// Validate validates s.
func (s Step) Validate() error {
	// TODO
	return nil
}
