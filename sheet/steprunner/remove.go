package steprunner

import (
	"fmt"
	"github.com/lukasl-dev/ben/sheet/step"
	"os"
)

// Remove runs a remove step.
func Remove(base step.Base, remove step.Remove) error {
	if err := os.Remove(remove.Remove); err != nil {
		return fmt.Errorf("step: %s: %w", base.Name, err)
	}
	return nil
}
