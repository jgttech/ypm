package spinner

import (
	"fmt"
	"jgttech/ypm/tui"
	"strconv"
)

func (m model) View() (s string) {
	var text string
	var indicator any
	var explanation string

	indicator = m.spinner.View()
	text = m.msg

	if m.status != -1 && m.status != 0 && m.status != 1 && m.status != 2 {
		explanation = "Invalid status code \"" + strconv.Itoa(m.status) + "\". Must be 0 (success), 1 (warning), or 2 (error)."
		indicator = tui.Indicators.Error
	} else {
		switch m.status {
		case -1:
			indicator = tui.Indicators.Error
			break
		case 1:
			indicator = tui.Indicators.Check
			break
		case 2:
			indicator = tui.Indicators.Warning
			break
		}
	}

	if m.info != "" {
		explanation = m.info
	}

	s += fmt.Sprintf("%s %s\n", indicator, tui.TextStyle(text))

	if explanation != "" {
		s += fmt.Sprintf("  %s\n", tui.TextStyle(explanation))
	}

	return s
}
