package confirm

import (
	"fmt"

	"jgttech/ypm/tui"
)

func (self model) View() (s string) {
	msg := tui.Theme.White(self.message)
	sep := tui.Theme.Dim("(y/n) ")
	view := self.textinput.View()
	errMsg := ""

	if self.err != nil {
		view = tui.Theme.Red(view)
		errMsg = "\n" + tui.Theme.RedBold(self.err.Error())
	}

	if self.answer != "" {
		view = ""

		if self.answer == "y" {
			sep = tui.Theme.Green("yes")
		}
		if self.answer == "n" {
			sep = tui.Theme.Dim("no")
		}
	}

	s += fmt.Sprintf("%s %s %s%s\n", msg, sep, view, errMsg)

	return
}
