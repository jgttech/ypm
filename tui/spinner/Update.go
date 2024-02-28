package spinner

import (
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

func (self model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case cmd:
		self.info = msg.info
		self.status = msg.status
		self.err = msg.err

		if msg.exit {
			if self.quit != nil {
				self.quit()
			}

			return self, tea.Quit
		}

		return self, nil
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc", "q":
			self.exit = true

			if self.quit != nil {
				self.quit()
			}

			return self, tea.Quit
		}
	case spinner.TickMsg:
		var cmd tea.Cmd
		self.spinner, cmd = self.spinner.Update(msg)
		return self, cmd
	}

	return self, self.spinner.Tick
}
