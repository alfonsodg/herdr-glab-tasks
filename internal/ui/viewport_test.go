package ui

import (
	"strings"
	"testing"

	"github.com/alfonsodg/herdr-glab-tasks/internal/gitlab"
)

func TestWindowRows(t *testing.T) {
	cases := []struct {
		rows, cursor, height int
		wantStart, wantEnd   int
	}{
		{10, 0, 5, 0, 5},
		{10, 4, 5, 0, 5},
		{10, 5, 5, 1, 6},
		{10, 7, 5, 3, 8},
		{10, 9, 5, 5, 10},
		{10, 0, 20, 0, 10},
		{10, 0, 0, 0, 10},
	}
	for _, tc := range cases {
		start, end := windowRows(tc.rows, tc.cursor, tc.height)
		if start != tc.wantStart || end != tc.wantEnd {
			t.Fatalf("windowRows(%d, %d, %d) = (%d, %d), want (%d, %d)",
				tc.rows, tc.cursor, tc.height, start, end, tc.wantStart, tc.wantEnd)
		}
	}
}

func TestRenderScrollsToCursor(t *testing.T) {
	var issues []gitlab.Issue
	for i := 1; i <= 30; i++ {
		issues = append(issues, gitlab.Issue{IID: i, Title: "issue", Labels: []string{"status::todo"}})
	}
	m := NewModel(issues, "")
	m.width, m.height = 80, 12

	out := m.render()
	if !strings.Contains(out, "#1 ") {
		t.Fatalf("top of list not rendered initially:\n%s", out)
	}

	for i := 0; i < 26; i++ {
		m.tree.MoveDown()
	}
	out = m.render()
	if !strings.Contains(out, "#26 ") {
		t.Fatalf("cursor row #26 not visible after scrolling:\n%s", out)
	}
	if strings.Contains(out, "#1 ") {
		t.Fatalf("top row should have scrolled off:\n%s", out)
	}
}
