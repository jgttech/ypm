package textinput

import "fmt"

func (self model) View() (s string) {
	s += fmt.Sprintf("%s\n%s", self.message, self.textinput.View())
	return
}
