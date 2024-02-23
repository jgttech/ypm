package tui

type indicators struct {
	Info    []byte
	Check   []byte
	Warning []byte
	Error   []byte
}

var Indicators indicators = indicators{
	Info:    []byte("i"),
	Check:   []byte("✓"),
	Warning: []byte("?"),
	Error:   []byte("ⅹ"),
}
