package runner

import (
	"errors"
	"github.com/lukasl-dev/ben/sheet/step"
	"os"
)

func rename(stp step.Step, _ Options) error {
	switch {
	case stp.Rename.Old == "":
		return errors.New("rename: old name is not set")
	case stp.Rename.New == "":
		return errors.New("rename: new name is not set")
	}
	return os.Rename(stp.Rename.Old, stp.Rename.New)
}
