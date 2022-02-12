package step

import (
	"fmt"
	"reflect"
)

// Step represents a general step in a configuration file. Only one of the
// embedded fields can be set.
type Step struct {
	Base
	*Command
	*Copy
	*Mkdir
	*Remove
	*Rename
}

// Validate validates s.
func (s Step) Validate() error {
	set := 0

	typ, val := reflect.TypeOf(s), reflect.ValueOf(s)
	for i := 0; i < typ.NumField(); i++ {
		if !val.Field(i).IsNil() {
			set++
		}
	}

	if set != 1 {
		return fmt.Errorf("step: %s: exactly one of the fields must be set: command, copy, mkdir, remove, rename", s.Name)
	}

	return nil
}
