package main

import "testing"

func TestFormatBranchStatus(t *testing.T) {
	if got := formatBranchStatus(42, "success"); got != "branch: Ref #42 · CI: success" {
		t.Fatalf("formatBranchStatus() = %q", got)
	}
}
