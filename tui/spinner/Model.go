package spinner

import (
	"jgttech/ypm/utils"

	"github.com/charmbracelet/bubbles/spinner"
)

type model struct {
	spinner spinner.Model
	err     string
	exit    bool
	status  int
	info    string
	msg     string
	done    string
	init    initFunc
	quit    quitFunc
}

type cmd struct {
	exit   bool
	status int
	err    string
	info   string
}

const FAILURE_MSG = "Oops, looks like something went wrong!"
const SUCCESS = utils.SUCCESS
const WARNING = utils.WARNING
const FAILURE = utils.FAILURE
