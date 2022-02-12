package steprunner

import (
	"fmt"
	"github.com/lukasl-dev/ben/sheet/step"
	cp "github.com/otiai10/copy"
)

// Copy runs a copy step.
func Copy(base step.Base, copy step.Copy) error {
	if err := cp.Copy(copy.Copy.From, copy.Copy.To); err != nil {
		return fmt.Errorf("step: %s: %w", base.Name, err)
	}
	return nil
}
