package spinner

import (
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"strings"
)

type model struct {
	spinner spinner.Model
	opts    Options
}

var _ tea.Model = (*model)(nil)

func New(opts Options) tea.Model {
	opts.normalize()

	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color(opts.Color))

	return &model{spinner: s, opts: opts}
}

func (m *model) Init() tea.Cmd {
	return m.spinner.Tick
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	m.spinner, cmd = m.spinner.Update(msg)
	return m, cmd
}

func (m *model) View() string {
	var builder strings.Builder

	if m.opts.Prefix != nil {
		builder.WriteString(*m.opts.Prefix)
	}

	builder.WriteString(m.spinner.View())

	if m.opts.Suffix != nil {
		builder.WriteString(*m.opts.Suffix)
	}

	return builder.String()
}
