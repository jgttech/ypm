package textinput

import (
	"jgttech/ypm/utils"

	"github.com/charmbracelet/bubbles/textinput"
)

type model struct {
	textinput   textinput.Model
	validate    textinput.ValidateFunc
	message     string
	placeholder string
}

const SUCCESS = utils.SUCCESS
const WARNING = utils.WARNING
const FAILURE = utils.FAILURE
