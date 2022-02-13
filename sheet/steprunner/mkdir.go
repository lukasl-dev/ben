package steprunner

import (
	"os"

	"github.com/lukasl-dev/ben/internal"
	"github.com/lukasl-dev/ben/sheet/step"
)

// Mkdir runs a mkdir step.
func Mkdir(st step.Step) error {
	if err := os.MkdirAll(*st.Mkdir, os.ModeDir); err != nil {
		return &internal.Error{
			Prefix: "step",
			Origin: st.Name,
			Err:    err,
			Suggestions: []string{
				"Check if the path is valid",
			},
		}
	}
	return nil
}
