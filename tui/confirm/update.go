package confirm

import (
	"errors"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

func (self model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	value := strings.ToLower(self.textinput.Value())

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return self, tea.Quit
		}
	}

	if value != "" {
		if value == "y" || value == "n" {
			self.answer = value
			return self, tea.Quit
		} else {
			self.err = errors.New("Invalid input")
		}
	} else {
		self.err = nil
	}

	self.textinput, cmd = self.textinput.Update(msg)
	return self, cmd
}
