package textinput

import "github.com/charmbracelet/bubbles/textinput"

type model struct {
	textinput   textinput.Model
	question    string
	placeholder string
	onChange    InputEvent
	onSubmit    InputEvent
}
