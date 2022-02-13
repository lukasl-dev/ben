package step

// Command represents a step that executes a command.
type Command struct {
	Command CommandData
}

// CommandData represents the type of Command's Command field.
type CommandData struct {
	// Command is the command to execute.
	Run string `json:"run,omitempty"`

	// WorkDir is the working directory for the command. It overwrites the
	// working directory of the sheet.
	WorkDir string `json:"workDir,omitempty"`

	// ExitCodes is a slice of exit codes that are considered successful.
	ExitCodes []int `json:"exitCodes,omitempty"`
}
