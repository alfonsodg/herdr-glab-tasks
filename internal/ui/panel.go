package ui

import (
	"fmt"
	"slices"
	"strings"

	"github.com/alfonsodg/herdr-glab-tasks/internal/gitlab"
)

var columns = []string{"todo", "review", "backlog", "done", "none"}

func FilterByLabel(issues []gitlab.Issue, label string) []gitlab.Issue {
	if label == "" {
		return issues
	}
	var filtered []gitlab.Issue
	for _, issue := range issues {
		if slices.Contains(issue.Labels, label) {
			filtered = append(filtered, issue)
		}
	}
	return filtered
}

func RenderPanel(grouped map[string][]gitlab.Issue) string {
	var out strings.Builder
	total := 0
	for _, col := range columns {
		for _, issue := range grouped[col] {
			total++
			_ = issue
		}
	}
	if total == 0 {
		return "No issues found.\n"
	}
	for _, col := range columns {
		issues := grouped[col]
		if len(issues) == 0 {
			continue
		}
		fmt.Fprintf(&out, "## %s (%d)\n", col, len(issues))
		for _, issue := range issues {
			labels := ""
			if len(issue.Labels) > 0 {
				labels = " [" + strings.Join(issue.Labels, ", ") + "]"
			}
			fmt.Fprintf(&out, "#%d %s%s\n", issue.IID, issue.Title, labels)
		}
		out.WriteString("\n")
	}
	return out.String()
}
