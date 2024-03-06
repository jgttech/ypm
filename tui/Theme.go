package tui

import "github.com/charmbracelet/lipgloss"

type render func(strs ...string) string

func color(colorCode string, bold bool) render {
	return lipgloss.NewStyle().Bold(bold).Foreground(lipgloss.Color(colorCode)).Render
}

type theme struct {
	SpinnerStyle lipgloss.Style
	White        render
	LightGray    render
	Gray         render
	DarkGray     render
	Green        render
	Red          render
	RedBold      render
	DarkRed      render
	Yellow       render
	Dim          render
}

var Theme = theme{
	SpinnerStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("69")),
	White:        color("#ffffff", false),
	LightGray:    color("#8a8a8a", false),
	Gray:         color("#626262", false),
	DarkGray:     color("#3c3c3c", false),
	Green:        color("#228B22", false),
	Red:          color("#c83c3c", false),
	RedBold:      color("#c83c3c", true),
	DarkRed:      color("#8c2a2a", false),
	Yellow:       color("#fbb829", false),
	Dim:          lipgloss.NewStyle().Foreground(lipgloss.Color("240")).Render,
}
