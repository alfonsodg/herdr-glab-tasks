package config

import (
	"os"
	"path/filepath"
	"testing"
)

func writeConfig(t *testing.T, body string) string {
	t.Helper()
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "config.toml"), []byte(body), 0o644); err != nil {
		t.Fatal(err)
	}
	return dir
}

func TestLoadMissingFileUsesDefaults(t *testing.T) {
	cfg := Load(t.TempDir())
	if cfg.State != "opened" || cfg.BranchType != "feature" || cfg.BranchScope != "tasks" {
		t.Fatalf("defaults wrong: %+v", cfg)
	}
	if cfg.Label != "" {
		t.Fatalf("default label = %q, want empty", cfg.Label)
	}
}

func TestLoadPartialFileKeepsDefaults(t *testing.T) {
	cfg := Load(writeConfig(t, `state = "all"`))
	if cfg.State != "all" {
		t.Fatalf("state = %q, want all", cfg.State)
	}
	if cfg.BranchType != "feature" || cfg.BranchScope != "tasks" {
		t.Fatalf("defaults lost on partial file: %+v", cfg)
	}
}

func TestLoadFullFile(t *testing.T) {
	cfg := Load(writeConfig(t, `
state = "closed"
label = "priority::high"
branch_type = "fix"
branch_scope = "auth"
`))
	if cfg.State != "closed" || cfg.Label != "priority::high" || cfg.BranchType != "fix" || cfg.BranchScope != "auth" {
		t.Fatalf("full file not applied: %+v", cfg)
	}
}

func TestLoadInvalidTomlFallsBack(t *testing.T) {
	cfg := Load(writeConfig(t, `state = [broken`))
	if cfg.State != "opened" {
		t.Fatalf("invalid toml must fall back to defaults, got %+v", cfg)
	}
}
