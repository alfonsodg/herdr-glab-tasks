package main

import (
	"strings"
	"testing"
)

func TestFormatBranchStatus(t *testing.T) {
	if got := formatBranchStatus(42, "success"); got != "branch: Ref #42 · CI: success" {
		t.Fatalf("formatBranchStatus() = %q", got)
	}
}

func TestReadInteractiveTitle(t *testing.T) {
	var out strings.Builder
	title, ok, err := readInteractiveTitle(strings.NewReader("Fix login\n"), &out)
	if err != nil {
		t.Fatalf("readInteractiveTitle: %v", err)
	}
	if !ok || title != "Fix login" {
		t.Fatalf("title = %q, ok = %v; want Fix login, true", title, ok)
	}
	if !strings.Contains(out.String(), "Issue title:") {
		t.Fatalf("prompt = %q", out.String())
	}
}

func TestReadInteractiveTitleBlankCancels(t *testing.T) {
	var out strings.Builder
	title, ok, err := readInteractiveTitle(strings.NewReader("\n"), &out)
	if err != nil {
		t.Fatalf("readInteractiveTitle: %v", err)
	}
	if ok || title != "" {
		t.Fatalf("title = %q, ok = %v; want empty, false", title, ok)
	}
	if !strings.Contains(out.String(), "Cancelled") {
		t.Fatalf("cancel output = %q", out.String())
	}
}
