package confirm

import (
	"jgttech/ypm/utils"

	"github.com/charmbracelet/bubbles/textinput"
)

type model struct {
	textinput textinput.Model
	validate  textinput.ValidateFunc
	message   string
	answer    string
	err       error
}

const YES = utils.YES
const NO = utils.NO
const UNKNOWN = utils.UNKNOWN
