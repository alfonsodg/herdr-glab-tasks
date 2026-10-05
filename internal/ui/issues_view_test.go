package ui

import (
	"testing"

	"github.com/alfonsodg/herdr-glab-tasks/internal/gitlab"
)

func TestGroupByStatus(t *testing.T) {
	issues := []gitlab.Issue{
		{IID: 1, Title: "first", Labels: []string{"status::todo"}},
		{IID: 2, Title: "second", Labels: []string{"status::review"}},
		{IID: 3, Title: "third"},
	}
	grouped := GroupByStatus(issues)
	if len(grouped["todo"]) != 1 || grouped["todo"][0].IID != 1 {
		t.Fatalf("todo group = %+v, want issue 1", grouped["todo"])
	}
	if len(grouped["review"]) != 1 || grouped["review"][0].IID != 2 {
		t.Fatalf("review group = %+v, want issue 2", grouped["review"])
	}
	if len(grouped["none"]) != 1 || grouped["none"][0].IID != 3 {
		t.Fatalf("none group = %+v, want issue 3", grouped["none"])
	}
}
