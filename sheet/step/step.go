package step

// Step represents a general step in a configuration file. Only one of the
// embedded fields can be set.
type Step struct {
	Base
	*Command
	*Copy
}
