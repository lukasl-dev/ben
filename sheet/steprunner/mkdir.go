package steprunner

import (
	"fmt"
	"github.com/lukasl-dev/ben/sheet/step"
	"os"
)

// Mkdir runs a mkdir step.
func Mkdir(base step.Base, mkdir step.Mkdir) error {
	if err := os.MkdirAll(mkdir.Mkdir, os.ModeDir); err != nil {
		return fmt.Errorf("step: %s: %w", base.Name, err)
	}
	return nil
}
