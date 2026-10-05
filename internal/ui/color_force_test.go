package ui

import (
	"os"
	"testing"
)

func TestForceColorRespectsNoColor(t *testing.T) {
	t.Setenv("NO_COLOR", "1")
	os.Unsetenv("CLICOLOR_FORCE")
	ForceColor()
	if os.Getenv("CLICOLOR_FORCE") != "" {
		t.Fatal("must not force when NO_COLOR is set")
	}
}

func TestForceColorSetsDefault(t *testing.T) {
	os.Unsetenv("NO_COLOR")
	os.Unsetenv("CLICOLOR_FORCE")
	ForceColor()
	if os.Getenv("CLICOLOR_FORCE") != "1" {
		t.Fatalf("CLICOLOR_FORCE = %q, want 1", os.Getenv("CLICOLOR_FORCE"))
	}
}

func TestForceColorKeepsExisting(t *testing.T) {
	os.Unsetenv("NO_COLOR")
	t.Setenv("CLICOLOR_FORCE", "0")
	ForceColor()
	if os.Getenv("CLICOLOR_FORCE") != "0" {
		t.Fatal("must not override explicit user value")
	}
}
