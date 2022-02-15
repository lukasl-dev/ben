package runner

import (
	"errors"
	"fmt"
	"github.com/lukasl-dev/ben/sheet/step"
	"os"
)

func mkdir(stp step.Step, _ Options) error {
	if stp.Mkdir == nil || *stp.Mkdir == "" {
		return errors.New("mkdir: no path specified")
	}

	if err := os.MkdirAll(*stp.Mkdir, os.ModeDir); err != nil {
		return fmt.Errorf("mkdir: %w", err)
	}
	return nil
}
