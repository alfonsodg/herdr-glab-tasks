package ui

import (
	"strings"
	"testing"

	"github.com/alfonsodg/herdr-glab-tasks/internal/gitlab"
)

func detailFixture() gitlab.Issue {
	return gitlab.Issue{
		IID:         42,
		Title:       "Fix login",
		State:       "opened",
		Labels:      []string{"priority::high", "status::todo"},
		WebURL:      "https://git.example.com/g/p/-/issues/42",
		Description: "line1\nline2\nline3\nline4\nline5",
		Author:      "ana",
	}
}

func TestDetailLines(t *testing.T) {
	lines := DetailLines(detailFixture(), 40)
	joined := strings.Join(lines, "\n")
	for _, want := range []string{"#42", "Fix login", "ana", "line1", "line5"} {
		if !strings.Contains(joined, want) {
			t.Fatalf("detail missing %q:\n%s", want, joined)
		}
	}
}

func TestDetailScrollWindow(t *testing.T) {
	lines := []string{"a", "b", "c", "d", "e"}
	visible := ScrollWindow(lines, 1, 3)
	if len(visible) != 3 || visible[0] != "b" || visible[2] != "d" {
		t.Fatalf("window = %v, want [b c d]", visible)
	}
	clamped := ScrollWindow(lines, 10, 3)
	if len(clamped) != 3 || clamped[0] != "c" {
		t.Fatalf("clamped window = %v, want last 3", clamped)
	}
}

func TestDetailModelScroll(t *testing.T) {
	m := NewModel(nil, "")
	issue := detailFixture()
	issue.Description = "l1\nl2\nl3\nl4\nl5\nl6\nl7\nl8\nl9\nl10\nl11\nl12\nl13\nl14\nl15\nl16\nl17\nl18\nl19\nl20\nl21\nl22\nl23\nl24\nl25"
	m.OpenDetail(issue)
	if !m.InDetail() {
		t.Fatal("expected detail open")
	}
	m.DetailDown()
	m.DetailDown()
	if m.DetailOffset() != 2 {
		t.Fatalf("offset = %d, want 2", m.DetailOffset())
	}
	m.CloseDetail()
	if m.InDetail() {
		t.Fatal("expected detail closed")
	}
	if m.tree.Cursor() != 0 {
		t.Fatalf("cursor moved, want preserved 0")
	}
}
