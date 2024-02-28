package spinner

import (
	"fmt"
	"jgttech/ypm/tui"
	"strconv"
)

func (self model) View() (s string) {
	var text string
	var indicator string
	var explanation string

	indicator = self.spinner.View()
	text = self.msg

	if self.status != 0 && self.status != SUCCESS && self.status != WARNING && self.status != FAILURE {
		indicator = tui.Indicators.Error
		explanation = "Invalid status code '" + strconv.Itoa(self.status) + "'. Must be "
		explanation += strconv.Itoa(SUCCESS) + " (success), "
		explanation += strconv.Itoa(WARNING) + " (warning), or "
		explanation += strconv.Itoa(FAILURE) + " (failure)."
	} else {
		switch self.status {
		case SUCCESS:
			indicator = tui.Theme.GreenText(tui.Indicators.Check)
			text = tui.Theme.DarkGray(text)
			break
		case WARNING:
			indicator = tui.Theme.YellowText(tui.Indicators.Warning)
			text = tui.Theme.YellowText(text)
			break
		case FAILURE:
			indicator = tui.Theme.ErrorText(tui.Indicators.Error)
			text = tui.Theme.ErrorText(text)
			break
		}
	}

	if self.info != "" {
		explanation = self.info
	}

	s += fmt.Sprintf("%s %s\n", indicator, tui.Theme.WhiteText(text))

	if explanation != "" {
		s += fmt.Sprintf("  %s\n", tui.Theme.Gray(explanation))
	}

	return s
}
