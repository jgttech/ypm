package spinner

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (self model) Init() tea.Cmd {
	return func() tea.Msg {
		msg := cmd{exit: true}

		if self.init != nil {
			code, message := self.init()

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
