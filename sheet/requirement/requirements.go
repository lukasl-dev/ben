package requirement

// Requirements contains required specifications that must be met to run a
// sheet.
type Requirements struct {
	// Envs is a slice of environmental variables that are required.
	Envs []Env `json:"envs,omitempty" hcl:"env,block"`

	// Paths is a slice of file or directory paths that are required.
	Paths []Path `json:"paths,omitempty" hcl:"path,block"`

	// Executable is a slice of executables that need to be installed.
	Executables []Executable `json:"executables,omitempty" hcl:"executable,block"`
}
