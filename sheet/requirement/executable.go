package requirement

// Executable configures a executable requirement.
type Executable struct {
	// Name is the name of the executable.
	Name string `json:"name,omitempty" hcl:"name,label"`
}
