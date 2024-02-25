package tui

import "fmt"

func toString(data []byte) string {
	return fmt.Sprintf("%s", data)
}

type indicators struct {
	Info    string
	Check   string
	Warning string
	Error   string
}

var Indicators indicators = indicators{
	Info:    toString([]byte("i")),
	Check:   toString([]byte("✓")),
	Warning: toString([]byte("!")),
	Error:   toString([]byte("ⅹ")),
}
