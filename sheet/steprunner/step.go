package steprunner

import (
	"errors"
	"fmt"
	"strings"

	"github.com/lukasl-dev/ben/internal"
	"github.com/lukasl-dev/ben/sheet"
	"github.com/lukasl-dev/ben/sheet/step"
)

// Step runs a step according to its type.
func Step(sh sheet.Sheet, st step.Step) error {
	switch {
	case st.Clean != nil:
		return clean(st)
	case st.Command != nil:
		return command(sh, st)
	case st.Copy != nil:
		return copy(st)
	case st.Mkdir != nil:
		return mkdir(st)
	case st.Remove != nil:
		return remove(st)
	case st.Rename != nil:
		return rename(st)
	default:
		return &internal.Error{
			Prefix: "step",
			Origin: st.Name,
			Err:    errors.New("unsupported or invalid task specified"),
			Suggestions: []string{
				fmt.Sprintf("Check if the step has one of the following fields: %s", strings.Join(step.StepFields(), ", ")),
			},
		}
	}
}
