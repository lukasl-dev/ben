package job

import "github.com/lukasl-dev/ben/sheet/step"

// Job is a collection of steps that can be executed by a runner.
type Job struct {
	// Name is the display name of a job.
	Name string `json:"name,omitempty" hcl:"name,label"`

	// State is the initial state of a job.
	State State `json:"state,omitempty" hcl:"state,optional"`

	// Steps is a slice of steps that can be executed.
	Steps []step.Step `json:"steps,omitempty" hcl:"step,block"`
}
