package spinner

import (
	"fmt"
	"jgttech/ypm/errors"
	"jgttech/ypm/tui"
	"os"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

type initFunc func() (int, string)
type quitFunc func()

type New struct {
	Msg  string
	Done string
	Init initFunc
	Quit quitFunc
}

func (self New) Run() (int, string) {
	m := model{
		msg:  self.Msg,
		done: self.Done,
		init: self.Init,
		quit: self.Quit,
	}

	m.spinner = spinner.New()
	m.spinner.Style = tui.Theme.SpinnerStyle
	m.spinner.Spinner = spinner.MiniDot

	result, err := tea.NewProgram(m).Run()
	response := result.(model)
	errorMsg := response.err

	if err != nil {
		errors.UnknownError(err)
	}

	if errorMsg != "" {
		fmt.Println("\n" + errorMsg)
		os.Exit(1)
	}

	return response.status, response.info
}
