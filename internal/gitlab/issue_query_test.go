package gitlab

import (
	"encoding/json"
	"testing"
)

func TestConvertIssueNode(t *testing.T) {
	raw := `{"iid":"42","title":"Fix login","state":"opened","webUrl":"https://git.example.com/g/p/-/issues/42","description":"Details here","updatedAt":"2026-10-01T10:00:00Z","author":{"username":"ana"},"labels":{"nodes":[{"title":"status::todo"},{"title":"priority::high"}]}}`
	var node issueNode
	if err := json.Unmarshal([]byte(raw), &node); err != nil {
		t.Fatalf("unmarshal fixture: %v", err)
	}
	issue := convertIssue(node)
	if issue.IID != 42 || issue.Title != "Fix login" || issue.Author != "ana" {
		t.Fatalf("converted = %+v", issue)
	}
	if issue.StatusColumn() != "todo" {
		t.Fatalf("StatusColumn() = %q, want todo", issue.StatusColumn())
	}
	if issue.Description != "Details here" {
		t.Fatalf("Description = %q", issue.Description)
	}
	if issue.UpdatedAt.IsZero() {
		t.Fatal("UpdatedAt not parsed")
	}
}

func TestConvertIssueNodeClosed(t *testing.T) {
	raw := `{"iid":"7","title":"Old","state":"closed","webUrl":"https://git.example.com/g/p/-/issues/7","updatedAt":"2026-09-01T00:00:00Z"}`
	var node issueNode
	if err := json.Unmarshal([]byte(raw), &node); err != nil {
		t.Fatalf("unmarshal fixture: %v", err)
	}
	if got := convertIssue(node).StatusColumn(); got != "done" {
		t.Fatalf("StatusColumn() = %q, want done", got)
	}
}

func TestBuildListQuery(t *testing.T) {
	q := buildListQuery("mygroup/myproject", "opened", "status::todo")
	for _, want := range []string{"mygroup/myproject", "opened", "status::todo", "issues("} {
		if !contains(q, want) {
			t.Fatalf("query missing %q:\n%s", want, q)
		}
	}
}
