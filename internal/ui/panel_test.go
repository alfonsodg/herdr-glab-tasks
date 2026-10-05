package ui

import (
	"strings"
	"testing"

	"github.com/alfonsodg/herdr-glab-tasks/internal/gitlab"
)

func TestRenderPanelGroupsByColumn(t *testing.T) {
	issues := []gitlab.Issue{
		{IID: 1, Title: "first", Labels: []string{"status::todo"}},
		{IID: 2, Title: "second", Labels: []string{"status::review"}},
	}
	out := RenderPanel(GroupByStatus(issues))
	if !strings.Contains(out, "## todo") || !strings.Contains(out, "#1 first") {
		t.Fatalf("panel missing todo section:\n%s", out)
	}
	if !strings.Contains(out, "## review") || !strings.Contains(out, "#2 second") {
		t.Fatalf("panel missing review section:\n%s", out)
	}
}

func TestRenderPanelEmpty(t *testing.T) {
	out := RenderPanel(GroupByStatus(nil))
	if !strings.Contains(out, "No issues") {
		t.Fatalf("empty panel = %q, want No issues hint", out)
	}
}

func TestFilterByLabel(t *testing.T) {
	issues := []gitlab.Issue{
		{IID: 1, Title: "first", Labels: []string{"priority::high", "status::todo"}},
		{IID: 2, Title: "second", Labels: []string{"priority::low"}},
	}
	filtered := FilterByLabel(issues, "priority::high")
	if len(filtered) != 1 || filtered[0].IID != 1 {
		t.Fatalf("filtered = %+v, want issue 1", filtered)
	}
	if got := FilterByLabel(issues, ""); len(got) != 2 {
		t.Fatalf("empty filter = %d issues, want 2", len(got))
	}
	if got := FilterByLabel(issues, "type::bug"); len(got) != 0 {
		t.Fatalf("no-match filter = %d issues, want 0", len(got))
	}
}
