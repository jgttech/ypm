package textinput

import (
	"log"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type InputEvent func(input string) string

type New struct {
	Question    string
	Placeholder string
	CharLimit   int
	Width       int
	Focus       bool
	OnChange    InputEvent
	OnSubmit    InputEvent
}

func (self New) Run() (int, string) {
	ti := textinput.New()
	ti.Placeholder = self.Placeholder
	ti.CharLimit = self.CharLimit
	ti.Width = self.Width

	m := model{
		textinput:   ti,
		question:    self.Question,
		placeholder: self.Placeholder,
		onChange:    self.OnChange,
		onSubmit:    self.OnSubmit,
	}

	if self.Focus {
		ti.Focus()
	}

	if _, err := tea.NewProgram(m).Run(); err != nil {
		log.Fatal(err)
	}

	return 0, ""
}
