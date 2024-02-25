package spinner

import (
	"fmt"
	"jgttech/ypm/tui"
	"strconv"
)

func (m model) View() (s string) {
	var text string
	var status string
	var indicator string
	var explanation string

	indicator = m.spinner.View()
	text = m.msg

	if m.status == 0 {
		status = "[---] "
	} else {
		status = "[" + strconv.Itoa(m.status) + "] "
	}

	status = tui.DarkGray(status)

	if m.status != 0 && m.status != SUCCESS && m.status != WARNING && m.status != FAILURE {
		indicator = tui.Indicators.Error
		explanation = "Invalid status code '" + strconv.Itoa(m.status) + "'. Must be "
		explanation += strconv.Itoa(SUCCESS) + " (success), "
		explanation += strconv.Itoa(WARNING) + " (warning), or "
		explanation += strconv.Itoa(FAILURE) + " (failure)."
	} else {
		switch m.status {
		case SUCCESS:
			indicator = tui.GreenText(tui.Indicators.Check)
			text = tui.DarkGray(text)
			break
		case WARNING:
			indicator = tui.YellowText(tui.Indicators.Warning)
			text = tui.YellowText(text)
			break
		case FAILURE:
			indicator = tui.ErrorText(tui.Indicators.Error)
			text = tui.ErrorText(text)
			break
		}
	}

	if m.info != "" {
		explanation = m.info
	}

	s += fmt.Sprintf("%s %s\n", indicator, tui.WhiteText(text))

	if explanation != "" {
		s += fmt.Sprintf("  %s\n", tui.Gray(explanation))
	}

	return s
}
