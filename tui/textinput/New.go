package textinput

import (
	"log"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type New struct {
	Message     string
	Placeholder string
	CharLimit   int
	Width       int
	Validate    textinput.ValidateFunc
}

func (self New) Run() (int, string) {
	width := 20
	charLimit := 156

	ti := textinput.New()
	ti.Placeholder = self.Placeholder

	ti.Focus()

	if self.Width > 0 {
		width = self.Width
	}

	if self.CharLimit > 0 {
		charLimit = self.CharLimit
	}

	ti.CharLimit = charLimit
	ti.Width = width

	m := model{
		textinput:   ti,
		message:     self.Message,
		placeholder: self.Placeholder,
		validate:    self.Validate,
	}

	if _, err := tea.NewProgram(m).Run(); err != nil {
		log.Fatal(err)
	}

	return 0, ""
}
