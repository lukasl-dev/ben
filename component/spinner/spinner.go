package spinner

import (
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Render renders a spinner in the terminal inside a goroutine whereby text()
// returns the string to be displayed on the right side next to the spinner. The
// returned function can be called to stop the spinner.
func Render(text func() string) (func(), <-chan error) {
	if text == nil {
		panic("Render() text cannot be nil")
	}

	program := createProgram(text)

	errs := make(chan error, 1)
	go func() {
		if err := program.Start(); err != nil {
			errs <- err
		}
		close(errs)
	}()

	return program.Quit, errs
}

// RenderStatic renders a spinner with a static text using Render().
func RenderStatic(text string) (func(), <-chan error) {
	return Render(func() string { return text })
}

// createProgram creates a bubbletea program that renders the spinner.
func createProgram(text func() string) *tea.Program {
	return tea.NewProgram(model{
		spinner: createSpinner(),
		text:    text,
	})
}

// createSpinner creates the spinner model to use.
func createSpinner() spinner.Model {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("#38bdf8"))
	return s
}
