package ui

import (
	"testing"

	"github.com/alfonsodg/herdr-glab-tasks/internal/gitlab"
)

func fixtureIssues() []gitlab.Issue {
	return []gitlab.Issue{
		{IID: 1, Title: "first", Labels: []string{"status::todo"}},
		{IID: 2, Title: "second", Labels: []string{"status::todo"}},
		{IID: 3, Title: "third", Labels: []string{"status::review"}},
	}
}

func TestTreeRows(t *testing.T) {
	tree := NewTree(fixtureIssues())
	if tree.RowCount() != 5 {
		t.Fatalf("RowCount = %d, want 5 (2 headers + 3 issues)", tree.RowCount())
	}
}

func TestTreeMove(t *testing.T) {
	tree := NewTree(fixtureIssues())
	tree.MoveDown()
	tree.MoveDown()
	if tree.Cursor() != 2 {
		t.Fatalf("Cursor = %d, want 2", tree.Cursor())
	}
	tree.MoveUp()
	if tree.Cursor() != 1 {
		t.Fatalf("Cursor = %d, want 1", tree.Cursor())
	}
}

func TestTreeMoveSkipsCollapsed(t *testing.T) {
	tree := NewTree(fixtureIssues())
	tree.Toggle()
	if tree.RowCount() != 3 {
		t.Fatalf("RowCount collapsed = %d, want 3", tree.RowCount())
	}
	tree.MoveDown()
	tree.MoveDown()
	row := tree.CurrentRow()
	if row == nil || row.Issue == nil || row.Issue.IID != 3 {
		t.Fatalf("current = %+v, want issue 3", row)
	}
}

func TestTreeFilter(t *testing.T) {
	tree := NewTree(fixtureIssues())
	tree.SetFilter("status::review")
	if tree.RowCount() != 2 {
		t.Fatalf("RowCount filtered = %d, want 2", tree.RowCount())
	}
	tree.SetFilter("")
	if tree.RowCount() != 5 {
		t.Fatalf("RowCount unfiltered = %d, want 5", tree.RowCount())
	}
}
