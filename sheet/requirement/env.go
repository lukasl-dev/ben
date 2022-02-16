package requirement

// Env configure a environment variable requirement.
type Env struct {
	// Name is the name of the environment variable to require.
	Name string `json:"name,omitempty" hcl:"name,label"`

	// Like is a regular expression that is matched against the value of the
	// environment variable. If it is empty, the environment variable's value is
	// arbitrary.
	Like string `json:"like,omitempty"`
}
