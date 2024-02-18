package tuiSpinner

import "fmt"

func (m model) View() (s string) {
	var text string
	var indicator any

	if m.exit {
		if m.done != "" {
			text = m.done
		} else {
			text = m.msg
		}

		indicator = []byte("✓")
	} else {
		text = m.msg
		indicator = m.spinner.View()
	}

	s += fmt.Sprintf("%s %s\n", indicator, textStyle(text))

	return
}
