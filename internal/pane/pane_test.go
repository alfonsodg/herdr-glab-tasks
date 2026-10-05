package pane

import (
	"context"
	"errors"
	"testing"
)

func TestOpenArgs(t *testing.T) {
	args := OpenArgs("alfonsodg.herdr-gitlab-issues", "issues", "/home/u/repo")
	want := []string{"plugin", "pane", "open", "--plugin", "alfonsodg.herdr-gitlab-issues", "--entrypoint", "issues", "--cwd", "/home/u/repo"}
	if len(args) != len(want) {
		t.Fatalf("args = %v, want %v", args, want)
	}
	for i := range want {
		if args[i] != want[i] {
			t.Fatalf("args = %v, want %v", args, want)
		}
	}
}

func TestOpenArgsNoCwd(t *testing.T) {
	args := OpenArgs("p", "e", "")
	for _, a := range args {
		if a == "--cwd" {
			t.Fatalf("args = %v, want no --cwd", args)
		}
	}
}

func TestFocusedCwd(t *testing.T) {
	out := `{"result":{"pane":{"foreground_cwd":"/home/u/repo","cwd":"/home/u/other"}}}`
	if got := FocusedCwd([]byte(out)); got != "/home/u/repo" {
		t.Fatalf("FocusedCwd = %q", got)
	}
	if got := FocusedCwd([]byte(`{"result":{"pane":{"cwd":"/home/u/other"}}}`)); got != "/home/u/other" {
		t.Fatalf("FocusedCwd fallback = %q", got)
	}
	if got := FocusedCwd([]byte(`not json`)); got != "" {
		t.Fatalf("FocusedCwd invalid = %q", got)
	}
}

func TestOpenMissingHerdr(t *testing.T) {
	opener := &Opener{Herdr: "herdr-missing-binary-xyz", Run: nil}
	opener.Run = func(ctx context.Context, name string, args ...string) error {
		return errors.New("exec: not found")
	}
	if err := opener.Open(context.Background(), "p", "e", ""); err == nil {
		t.Fatal("expected error when herdr CLI missing")
	}
}
