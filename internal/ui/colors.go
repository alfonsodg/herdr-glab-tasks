package ui

import (
	"os"

	"charm.land/lipgloss/v2"
)

func ColumnColorName(col string) string {
	switch col {
	case "todo":
		return "11"
	case "review":
		return "13"
	case "backlog":
		return "8"
	case "done":
		return "10"
	case "none":
		return "6"
	default:
		return ""
	}
}

func PriorityColorName(labels []string) string {
	for _, l := range labels {
		switch l {
		case "priority::critical":
			return "9"
		case "priority::high":
			return "11"
		case "priority::medium":
			return "14"
		case "priority::low":
			return "8"
		}
	}
	return ""
}

func columnStyle(col string) lipgloss.Style {
	name := ColumnColorName(col)
	if name == "" {
		return groupStyle
	}
	return lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color(name))
}

func issueStyle(labels []string) lipgloss.Style {
	name := PriorityColorName(labels)
	if name == "" {
		return lipgloss.NewStyle()
	}
	return lipgloss.NewStyle().Foreground(lipgloss.Color(name))
}

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

func ForceColor() {
	if os.Getenv("NO_COLOR") != "" {
		return
	}
	if _, set := os.LookupEnv("CLICOLOR_FORCE"); !set {
		os.Setenv("CLICOLOR_FORCE", "1")
	}
}
