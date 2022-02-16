package sheet

import (
	"github.com/lukasl-dev/ben/sheet/job"
	"github.com/lukasl-dev/ben/sheet/requirement"
)

// Sheet is a collection of jobs. It represents the content of a configuration
// file.
type Sheet struct {
	// Name is the display name of the sheet.
	Name string `json:"name,omitempty" hcl:"name"`

	// Requirements contains requirements that must be met to run this sheet.
	Requirements *requirement.Requirements `json:"requirements,omitempty" hcl:"requirements,block"`

	// Description is a short description about the usage of the sheet.
	Description string `json:"description,omitempty" hcl:"description,optional"`

	// Jobs is a slice of jobs that can be executed on the sheet. The key is the
	// unique name of a job.
	Jobs []job.Job `json:"jobs,omitempty" hcl:"job,block"`
}

// Validate validates s.
func (s Sheet) Validate() error {
	for _, j := range s.Jobs {
		for _, step := range j.Steps {
			if err := step.Validate(); err != nil {
				return err
			}
		}
	}
	return nil
}
