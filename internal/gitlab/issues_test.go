package gitlab

import "testing"

func TestNewClientDelegatesAuthToGlab(t *testing.T) {
	c := NewClient("/usr/bin/glab", "git.example.com")
	if c == nil {
		t.Fatal("NewClient returned nil")
	}
	if ErrGlabMissing == nil || ErrUnauthorized == nil {
		t.Fatal("glab auth sentinel errors must be defined")
	}
}
