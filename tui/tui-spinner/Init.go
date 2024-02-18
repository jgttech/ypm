package tuiSpinner

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (m model) Init() tea.Cmd {
	return func() tea.Msg {
		if m.init != nil {
			m.init()
		}

		return "exit"
	}
}
