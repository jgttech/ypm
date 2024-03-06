package confirm

import (
	"jgttech/ypm/utils"
	"log"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type New struct {
	Message  string
	Validate textinput.ValidateFunc
}

func (self New) Run() int {
	ti := textinput.New()
	ti.Prompt = ""
	ti.Cursor.Blink = false

	ti.Focus()

	ti.Width = 4
	ti.CharLimit = 1

	m := model{
		textinput: ti,
		message:   self.Message,
		validate:  self.Validate,
	}

	if _, err := tea.NewProgram(m).Run(); err != nil {
		log.Fatal(err)
	}

	return utils.UNKNOWN
}
