package spinner

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (m model) Init() tea.Cmd {
	return func() tea.Msg {
		msg := cmd{exit: true}

		if m.init != nil {
			code, message := m.init()

			msg.status = code

			if code == FAILURE {
				msg.err = FAILURE_MSG
			}

			if message != "" {
				msg.info = message
			}
		}

		return msg
	}
}
