package tui

import "github.com/charmbracelet/lipgloss"

var (
	SpinnerStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("69"))
	TextStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("252")).Render
)
