package steprunner

import (
	"os"

	"github.com/lukasl-dev/ben/internal"
	"github.com/lukasl-dev/ben/sheet/step"
)

// rename runs a rename step.
func rename(st step.Step) error {
	if err := os.Rename(st.Rename.Old, st.Rename.New); err != nil {
		return &internal.Error{
			Prefix: "step",
			Origin: st.Name,
			Err:    err,
			Suggestions: []string{
				"Check if the provided paths are valid",
			},
		}
	}
	return nil
}
