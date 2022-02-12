package step

// Command represents a step that executes a command.
type Command struct {
	Base

	// Command is the command to execute.
	Command string `json:"command,omitempty"`
}
