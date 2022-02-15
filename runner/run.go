package runner

import (
	"github.com/lukasl-dev/ben/sheet/step"
	"io"
	"reflect"
)

// runner represents a runner function whereby 'j' is the index of the job
// in sh's jobs slice and 'st' is the index of the step in the job's steps.
type runner = func(stp step.Step, opts Options) error

// runners is a map from a field name of step.Step to a runner.
var runners = map[string]runner{
	"Clean":   clean,
	"Command": command,
	"Copy":    ccopy,
	"Mkdir":   mkdir,
	"Remove":  remove,
	"Rename":  rename,
}

type Options struct {
	Stdout io.Writer `json:"stdout,omitempty"`
	Stderr io.Writer `json:"stderr,omitempty"`
}

// NoOptions are the default Options.
var NoOptions = Options{}

// Run runs stp with the given opts.
func Run(stp step.Step, opts Options) error {
	run := findRunner(stp)
	if run == nil {
		return newError(stp, "no task specified")
	}

	if err := run(stp, opts); err != nil {
		return newError(stp, err)
	}
	return nil
}

func findRunner(stp step.Step) runner {
	val := reflect.Indirect(reflect.ValueOf(stp))
	for name, run := range runners {
		f := val.FieldByName(name)
		if !f.IsZero() {
			return run
		}
	}
	return nil
}
