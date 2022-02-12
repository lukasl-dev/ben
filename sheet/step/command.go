package step

// Command represents a step that executes a command.
type Command struct {
	// Command is the command to execute.
	Command string `json:"command,omitempty"`

	// WorkDir is the working directory for the command. It overwrites the
	// working directory of the sheet.
	WorkDir string `json:"workdir,omitempty"`
}
