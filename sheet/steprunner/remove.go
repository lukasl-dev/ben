package steprunner

import (
	"os"

	"github.com/lukasl-dev/ben/internal"
	"github.com/lukasl-dev/ben/sheet/step"
)

// remove runs a remove step.
func remove(st step.Step) error {
	if err := os.Remove(*st.Remove); err != nil {
		return &internal.Error{
			Prefix: "step",
			Origin: st.Name,
			Err:    err,
			Suggestions: []string{
				"Check if the file or directory exists",
			},
		}
	}
	return nil
}
