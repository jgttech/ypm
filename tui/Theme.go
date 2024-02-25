package tui

import "github.com/charmbracelet/lipgloss"

var (
	SpinnerStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("69"))
	WhiteText    = lipgloss.NewStyle().Foreground(lipgloss.Color("#ffffff")).Render
	LightGray    = lipgloss.NewStyle().Foreground(lipgloss.Color("#8a8a8a")).Render
	Gray         = lipgloss.NewStyle().Foreground(lipgloss.Color("#626262")).Render
	DarkGray     = lipgloss.NewStyle().Foreground(lipgloss.Color("#3c3c3c")).Render
	GreenText    = lipgloss.NewStyle().Foreground(lipgloss.Color("#228B22")).Render
	ErrorText    = lipgloss.NewStyle().Foreground(lipgloss.Color("#c83c3c")).Render
	YellowText   = lipgloss.NewStyle().Foreground(lipgloss.Color("#fbb829")).Render
)
