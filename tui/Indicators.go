package tui

import "fmt"

func str(data []byte) string {
	return fmt.Sprintf("%s", data)
}

type indicators struct {
	Info    string
	Check   string
	Warning string
	Error   string
}

var Indicators indicators = indicators{
	Info:    str([]byte("i")),
	Check:   str([]byte("✓")),
	Warning: str([]byte("!")),
	Error:   str([]byte("ⅹ")),
}
