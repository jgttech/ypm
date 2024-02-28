package tui

import "github.com/charmbracelet/lipgloss"

type render func(strs ...string) string

func color(color string) render {
	return lipgloss.NewStyle().Foreground(lipgloss.Color(color)).Render
}

type theme struct {
	SpinnerStyle lipgloss.Style
	WhiteText    render
	LightGray    render
	Gray         render
	DarkGray     render
	GreenText    render
	ErrorText    render
	YellowText   render
}

var Theme = theme{
	SpinnerStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("69")),
	WhiteText:    color("#ffffff"),
	LightGray:    color("#8a8a8a"),
	Gray:         color("#626262"),
	DarkGray:     color("#3c3c3c"),
	GreenText:    color("#228B22"),
	ErrorText:    color("#c83c3c"),
	YellowText:   color("#fbb829"),
}
