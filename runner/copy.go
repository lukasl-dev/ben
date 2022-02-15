package runner

import (
	"errors"
	"fmt"
	"github.com/lukasl-dev/ben/sheet/step"
	cp "github.com/otiai10/copy"
)

//goland:noinspection ALL
func ccopy(stp step.Step, _ Options) error {
	switch {
	case stp.Copy.From == "":
		return errors.New("copy: no source path specified")
	case stp.Copy.To == "":
		return errors.New("copy: no destination path specified")
	}

	if err := cp.Copy(stp.Copy.From, stp.Copy.To); err != nil {
		return fmt.Errorf("copy: %w", err)
	}
	return nil
}
