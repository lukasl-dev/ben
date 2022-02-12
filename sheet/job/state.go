package job

// State is the initial state of a Job.
type State string

const (
	// Required indicates that a job is required to be run.
	Required State = "required"

	// Enabled indicates that a job enabled by default.
	Enabled State = "enabled"

	// Disabled indicates that a job is disabled by default.
	Disabled State = "disabled"
)
