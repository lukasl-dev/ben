package spinner

import (
	"fmt"
	"github.com/briandowns/spinner"
	"github.com/fatih/color"
	"strings"
	"time"
)

// Spinner represents a terminal spinner.
type Spinner struct {
	// opts is the spinner configuration.
	opts Options

	// spin is the underlying spinner.
	spin *spinner.Spinner
}

// New returns a new Spinner.
func New(opts Options) *Spinner {
	return &Spinner{opts: opts, spin: createSpinner(opts)}
}

// createSpinner constructs a new spinner from opts.
func createSpinner(opts Options) *spinner.Spinner {
	spin := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
	spin.Prefix = strings.Repeat(" ", int(opts.Indent))
	spin.Suffix = " " + opts.Pending
	_ = spin.Color("blue")
	return spin
}

// Start starts the spinner.
func (s *Spinner) Start() {
	s.spin.Start()
}

// Error stops the spinner with the error message that is specified in the
// Options.
func (s *Spinner) Error() {
	s.spin.FinalMSG = fmt.Sprintf(" %s %s", color.RedString("×"), s.opts.Error)
	s.spin.Stop()
}

// Success stops the spinner with the success message that is specified in the
// Options.
func (s *Spinner) Success() {
	s.spin.FinalMSG = fmt.Sprintf(" %s %s", color.GreenString("✓"), s.opts.Success)
	s.spin.Stop()
}

// Update replaces the spinner's message with the specified one.
func (s *Spinner) Update(text string) {
	s.spin.Suffix = " " + text
}
