package tuiSpinner

import (
	"jgttech/ypm/exceptions"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type element struct {
	model model
}

type Runner struct {
	Msg  string
	Done string
	Init func()
}

var (
	spinnerStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("69"))
	textStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("252")).Render
)

func New() *element {
	m := model{
		msg:  "...",
		exit: false,
	}

	m.spinner = spinner.New()
	m.spinner.Style = spinnerStyle
	m.spinner.Spinner = spinner.MiniDot

	return &element{
		model: m,
	}
}

func (e *element) Run(runner Runner) {
	e.model.init = runner.Init
	e.model.msg = runner.Msg
	e.model.done = runner.Done

	if _, err := tea.NewProgram(e.model).Run(); err != nil {
		exceptions.UnknownError(err)
	}
}
