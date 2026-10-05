package branch

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func initRepo(t *testing.T) string {
	t.Helper()
	dir := t.TempDir()
	for _, args := range [][]string{{"init"}, {"config", "user.email", "t@t"}, {"config", "user.name", "t"}, {"commit", "--allow-empty", "-m", "chore(test): init (#0)", "-m", "Ref #0"}} {
		if out, err := exec.Command("git", append([]string{"-C", dir}, args...)...).CombinedOutput(); err != nil {
			t.Fatalf("git %v: %v\n%s", args, err, out)
		}
	}
	return dir
}

func TestBranchName(t *testing.T) {
	if got := Name("feature", "auth", 42); got != "feature/auth-#42" {
		t.Fatalf("Name() = %q", got)
	}
}

func TestCreateIssueBranch(t *testing.T) {
	dir := initRepo(t)
	if err := CreateIssueBranch(dir, "feature", "auth", 42); err != nil {
		t.Fatalf("CreateIssueBranch: %v", err)
	}
	out, _ := exec.Command("git", "-C", dir, "branch", "--show-current").Output()
	if got := string(out); got != "feature/auth-#42\n" {
		t.Fatalf("branch = %q", got)
	}
}

func TestCreateIssueBranchRefusesDirtyTree(t *testing.T) {
	dir := initRepo(t)
	if err := os.WriteFile(filepath.Join(dir, "dirty.txt"), []byte("x"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := CreateIssueBranch(dir, "feature", "auth", 1); err == nil {
		t.Fatal("expected error on dirty tree")
	}
}

func TestBranchIssueRef(t *testing.T) {
	dir := initRepo(t)
	if _, err := exec.Command("git", "-C", dir, "commit", "--allow-empty", "-m", "feat(auth): login (#7)\n\nRef #7").CombinedOutput(); err != nil {
		t.Fatal(err)
	}
	iid, ok := IssueRef(dir)
	if !ok || iid != 7 {
		t.Fatalf("IssueRef() = %d, %v; want 7, true", iid, ok)
	}
}

func TestBranchIssueRefNone(t *testing.T) {
	dir := initRepo(t)
	if out, err := exec.Command("git", "-C", dir, "commit", "--no-verify", "--allow-empty", "-m", "wip").CombinedOutput(); err != nil {
		t.Fatalf("setup commit: %v\n%s", err, out)
	}
	if _, ok := IssueRef(dir); ok {
		t.Fatal("expected no ref on plain commit")
	}
}
