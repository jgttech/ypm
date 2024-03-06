package confirm

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

func (self model) Init() tea.Cmd {
	return textinput.Blink
}
