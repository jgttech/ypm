package textinput

import tea "github.com/charmbracelet/bubbletea"

func (self model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return self, tea.Quit
		case tea.KeyEnter:
			if self.validate != nil {
				err := self.validate(self.textinput.Value())

				if err != nil {
					return self, nil
				}
			}

			return self, tea.Quit
		}
	}

	self.textinput, cmd = self.textinput.Update(msg)
	return self, cmd
}