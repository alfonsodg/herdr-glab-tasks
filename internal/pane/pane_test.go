package pane

import (
	"context"
	"errors"
	"testing"
)

func TestOpenArgs(t *testing.T) {
	args := OpenArgs("alfonsodg.herdr-gitlab-issues", "issues")
	want := []string{"plugin", "pane", "open", "--plugin", "alfonsodg.herdr-gitlab-issues", "--entrypoint", "issues"}
	if len(args) != len(want) {
		t.Fatalf("args = %v, want %v", args, want)
	}
	for i := range want {
		if args[i] != want[i] {
			t.Fatalf("args = %v, want %v", args, want)
		}
	}
}

func TestOpenMissingHerdr(t *testing.T) {
	opener := &Opener{Herdr: "herdr-missing-binary-xyz", Run: nil}
	opener.Run = func(ctx context.Context, name string, args ...string) error {
		return errors.New("exec: not found")
	}
	if err := opener.Open(context.Background(), "p", "e"); err == nil {
		t.Fatal("expected error when herdr CLI missing")
	}
}
