// Package ui renders the issues panel grouped by status column.
package ui

import "github.com/alfonsodg/herdr-glab-tasks/internal/gitlab"

func GroupByStatus(issues []gitlab.Issue) map[string][]gitlab.Issue {
	grouped := map[string][]gitlab.Issue{}
	for _, issue := range issues {
		col := issue.StatusColumn()
		grouped[col] = append(grouped[col], issue)
	}
	return grouped
}
