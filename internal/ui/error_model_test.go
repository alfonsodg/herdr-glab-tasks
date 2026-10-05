package ui

import (
	"strings"
	"testing"

	tea "charm.land/bubbletea/v2"
)

func TestErrorModelRendersMessage(t *testing.T) {
	m := NewErrorModel("no git origin remote")
	out := m.render()
	if !strings.Contains(out, "no git origin remote") {
		t.Fatalf("error not rendered:\n%s", out)
	}
	if strings.Contains(out, "No issues found") {
		t.Fatalf("error model must not show empty list hint:\n%s", out)
	}
	if !strings.Contains(out, "q quit") && !strings.Contains(out, "q") {
		t.Fatalf("footer missing quit hint:\n%s", out)
	}
}

func TestErrorModelQuitsOnQ(t *testing.T) {
	m := NewErrorModel("boom")
	updated, cmd := m.Update(tea.KeyPressMsg{Text: "q", Code: 'q'})
	if cmd == nil {
		t.Fatal("q must return a quit command")
	}
	_ = updated
}
