package spinner

import (
	spnr "github.com/charmbracelet/bubbles/spinner"
)

type model struct {
	spinner spnr.Model
	exit    bool
	status  int
	info    string
	msg     string
	done    string
	init    Init
	quit    Quit
}

type cmd struct {
	exit   bool
	status int
	info   string
}
