package spinner

import (
	"jgttech/ypm/exceptions"
	"jgttech/ypm/tui"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

type Init func() (int, string)
type Quit func()

type New struct {
	Msg  string
	Done string
	Init Init
	Quit Quit
}

func (instance New) Run() {
	m := model{
		msg:  instance.Msg,
		done: instance.Done,
		init: instance.Init,
		quit: instance.Quit,
	}

	m.spinner = spinner.New()
	m.spinner.Style = tui.SpinnerStyle
	m.spinner.Spinner = spinner.MiniDot

	if _, err := tea.NewProgram(m).Run(); err != nil {
		exceptions.UnknownError(err)
	}
}
