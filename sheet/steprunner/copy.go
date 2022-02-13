package steprunner

import (
	"github.com/lukasl-dev/ben/internal"
	"github.com/lukasl-dev/ben/sheet/step"
	cp "github.com/otiai10/copy"
)

// copy runs a copy step.
func copy(st step.Step) error {
	err := cp.Copy(st.Copy.From, st.Copy.To)
	if err != nil {
		return &internal.Error{
			Prefix: "step",
			Origin: st.Name,
			Err:    err,
			Suggestions: []string{
				"Check if the source and destination paths are valid",
			},
		}
	}
	return nil
}
