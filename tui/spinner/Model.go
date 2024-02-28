package spinner

import "github.com/charmbracelet/bubbles/spinner"

type model struct {
	spinner spinner.Model
	err     string
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
	err    string
	info   string
}

const FAILURE_MSG = "Oops, looks like something went wrong!"

// For ease of use, the status codes that should be
// returned from the spinner function should be one of
// the valid HTTP status codes listed here. I decided
// on using HTTP status codes because it leverages
// existing web development knowledge.
//
// Status Codes:
// 199: Miscellaneous Warning
// 200: Success
// 501: Internal Server Error
const SUCCESS = 200
const WARNING = 199
const FAILURE = 501
