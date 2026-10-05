package ui

import (
	"fmt"
	"strings"

	"github.com/alfonsodg/herdr-glab-tasks/internal/gitlab"
)

type Detail struct {
	Issue  gitlab.Issue
	offset int
	height int
}

func DetailLines(issue gitlab.Issue, width int) []string {
	var lines []string
	lines = append(lines, fmt.Sprintf("#%d %s", issue.IID, issue.Title))
	meta := issue.State
	if issue.Author != "" {
		meta += " · " + issue.Author
	}
	if !issue.UpdatedAt.IsZero() {
		meta += " · " + issue.UpdatedAt.Format("2006-01-02")
	}
	lines = append(lines, dimStyle.Render(meta))
	if issue.WebURL != "" {
		lines = append(lines, dimStyle.Render(issue.WebURL))
	}
	if len(issue.Labels) > 0 {
		lines = append(lines, renderLabels(issue.Labels))
	}
	lines = append(lines, "")
	if issue.Description != "" {
		lines = append(lines, strings.Split(issue.Description, "\n")...)
	} else {
		lines = append(lines, dimStyle.Render("No description."))
	}
	_ = width
	return lines
}

func renderLabels(labels []string) string {
	parts := make([]string, 0, len(labels))
	for _, l := range labels {
		parts = append(parts, labelStyle(l).Render(l))
	}
	return strings.Join(parts, " ")
}

func ScrollWindow(lines []string, offset, height int) []string {
	if height <= 0 {
		return nil
	}
	if offset < 0 {
		offset = 0
	}
	maxOffset := len(lines) - height
	if maxOffset < 0 {
		maxOffset = 0
	}
	if offset > maxOffset {
		offset = maxOffset
	}
	end := offset + height
	if end > len(lines) {
		end = len(lines)
	}
	return lines[offset:end]
}

func windowRows(total, cursor, height int) (int, int) {
	if height <= 0 || height >= total {
		return 0, total
	}
	start := 0
	if cursor >= height {
		start = cursor - height + 1
	}
	end := start + height
	if end > total {
		end = total
		start = max(end-height, 0)
	}
	return start, end
}
