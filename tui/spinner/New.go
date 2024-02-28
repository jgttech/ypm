package spinner

import (
	"fmt"
	"jgttech/ypm/exceptions"
	"jgttech/ypm/tui"
	"os"

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
		exceptions.UnknownError(err)
	}

	if errorMsg != "" {
		fmt.Println("\n" + errorMsg)
		os.Exit(1)
	}

	return response.status, response.info
}
