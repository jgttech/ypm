package errors

import (
	"fmt"
	"jgttech/ypm/tui"
	"os"
)

func PathNotFound(path string) {
	msg := "\n" + tui.Theme.RedBold("ERROR") + "\n"
	msg += tui.Theme.DarkRed("Path not found: " + "\"" + path + "\"")

	fmt.Println(msg)
	os.Exit(1)
}
