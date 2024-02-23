package spinner

import (
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case cmd:
		m.info = msg.info
		m.status = msg.status

		if msg.exit {
			if m.quit != nil {
				m.quit()
			}

			return m, tea.Quit
		}

		return m, nil
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc", "q":
			m.exit = true

			if m.quit != nil {
				m.quit()
			}

			return m, tea.Quit
		}
	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}

	return m, m.spinner.Tick
}
