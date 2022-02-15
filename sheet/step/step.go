package step

import "reflect"

// Step represents a general step in a configuration file. Only one of the
// embedded fields can be set.
type Step struct {
	// Name is the name of the step. It is used for logging and representative
	// purposes.
	Name string `json:"name,omitempty" hcl:"name,label"`

	// Clean is the data for cleaning a directory. If it is set, the step is a
	// Clean step.
	Clean *Clean `json:"clean,omitempty" hcl:"clean,block"`

	// Command is the data for running a command. If it is set, the step is a
	// Command step.
	Command *Command `json:"command,omitempty" hcl:"command,block"`

	// Copy is the data for copying a file or directory into another location.
	// If it is set, the step is a Copy step.
	Copy *Copy `json:"copy,omitempty" hcl:"copy,block"`

	// Mkdir is the path to the directory to create. If it is set, the step is a
	// Mkdir step.
	Mkdir *string `json:"mkdir,omitempty" hcl:"mkdir,optional"`

	// Remove is th path to the file or directory to remove. If it is set, the
	// step is a Remove step.
	Remove *string `json:"remove,omitempty" hcl:"remove,optional"`

	// Rename is the data for renaming a file or directory. If it is set, the
	// step is a Rename step.
	Rename *Rename `json:"rename,omitempty" hcl:"rename,block"`
}

// Validate validates s.
func (s Step) Validate() error {
	// TODO
	return nil
}

// StepFields returns the json names of all step configuration fields.
func StepFields() []string {
	typ := reflect.TypeOf(Step{})
	names := make([]string, typ.NumField()+1)

	for i := 1; i < typ.NumField(); i++ {
		names[i] = typ.Field(i).Tag.Get("json")
	}
	return names
}
