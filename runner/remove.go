package runner

import (
	"errors"
	"fmt"
	"github.com/lukasl-dev/ben/sheet/step"
	"os"
)

func remove(stp step.Step, _ Options) error {
	if stp.Remove == nil || *stp.Remove == "" {
		return errors.New("remove: no path specified")
	}

	if err := os.RemoveAll(*stp.Remove); err != nil {
		return fmt.Errorf("remove: %w", err)
	}
	return nil
}
