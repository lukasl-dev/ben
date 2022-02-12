package steprunner

import (
	"fmt"
	"github.com/lukasl-dev/ben/sheet/step"
	"os"
)

// Rename runs a rename step.
func Rename(base step.Base, rename step.Rename) error {
	if err := os.Rename(rename.Rename.Old, rename.Rename.New); err != nil {
		return fmt.Errorf("step: %s: %w", base.Name, err)
	}
	return nil
}
