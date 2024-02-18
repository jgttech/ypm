package tuiSpinner

import (
	"github.com/charmbracelet/bubbles/spinner"
)

type modelMsg struct {
	text string
}

type model struct {
	// The message to display with the spinner
	msg string

	// The message that should show on exit
	done string

	// Indicates that the program should exit
	exit bool

	// The custom init function we can run
	init func()

	// The config for the spinner
	spinner spinner.Model
}
