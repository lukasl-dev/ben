package sheet

import "github.com/lukasl-dev/ben/sheet/job"

// Sheet is a collection of jobs. It represents the content of a configuration
// file.
type Sheet struct {
	// Name is the display name of the sheet.
	Name string `json:"name,omitempty"`

	// Description is a short description about the usage of the sheet.
	Description string `json:"description,omitempty"`

	// Jobs is a map of jobs that can be executed on the sheet. The key is the
	// unique name of a job.
	Jobs map[string]job.Job `json:"jobs,omitempty"`
}
