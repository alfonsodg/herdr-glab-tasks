package ui

import "charm.land/lipgloss/v2"

func LabelColorName(label string) string {
	switch label {
	case "priority::critical", "type::bug", "type::security":
		return "1"
	case "priority::high":
		return "3"
	case "priority::medium":
		return "4"
	case "priority::low":
		return "8"
	case "type::feature":
		return "2"
	case "status::todo", "status::review", "status::backlog", "status::done":
		return "8"
	default:
		return ""
	}
}

func labelStyle(label string) lipgloss.Style {
	name := LabelColorName(label)
	if name == "" {
		return lipgloss.NewStyle()
	}
	return lipgloss.NewStyle().Foreground(lipgloss.Color(name))
}
