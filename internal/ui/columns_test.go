package ui

import (
	"strings"
	"testing"

	"github.com/alfonsodg/herdr-glab-tasks/internal/gitlab"
)

func TestColumnColorName(t *testing.T) {
	seen := map[string]string{}
	for _, col := range columns {
		name := ColumnColorName(col)
		if name == "" {
			t.Fatalf("column %q has no color", col)
		}
		if prev, dup := seen[name]; dup {
			t.Fatalf("column %q shares color %q with %q", col, name, prev)
		}
		seen[name] = col
	}
	if ColumnColorName("nope") != "" {
		t.Fatal("unknown column must have no color")
	}
}

func TestPriorityColorName(t *testing.T) {
	cases := []struct {
		labels []string
		want   string
	}{
		{[]string{"priority::critical"}, "9"},
		{[]string{"priority::high"}, "11"},
		{[]string{"priority::medium"}, "14"},
		{[]string{"priority::low"}, "8"},
		{nil, ""},
		{[]string{"status::todo"}, ""},
	}
	for _, tc := range cases {
		if got := PriorityColorName(tc.labels); got != tc.want {
			t.Fatalf("PriorityColorName(%v) = %q, want %q", tc.labels, got, tc.want)
		}
	}
}

func TestSelectedRowKeepsChipColor(t *testing.T) {
	m := NewModel([]gitlab.Issue{{IID: 7, Title: "x", Labels: []string{"priority::critical", "status::todo"}}}, "")
	out := m.render()
	if !strings.Contains(out, "\x1b[9m") && !strings.Contains(out, "\x1b[91m") {
		t.Fatalf("selected line lost critical chip color:\n%q", out)
	}
	if !strings.Contains(out, "\u258c") {
		t.Fatalf("selected line missing bar marker:\n%q", out)
	}
}
