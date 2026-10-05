package repo

import "testing"

func TestParseRemote(t *testing.T) {
	cases := []struct {
		raw     string
		host    string
		project string
		ok      bool
	}{
		{"git@github.com:alfonsodg/herdr-glab-tasks.git", "github.com", "alfonsodg/herdr-glab-tasks", true},
		{"https://git.example.com/group/project.git", "git.example.com", "group/project", true},
		{"not-a-remote", "", "", false},
	}
	for _, tc := range cases {
		host, project, ok := ParseRemote(tc.raw)
		if host != tc.host || project != tc.project || ok != tc.ok {
			t.Fatalf("ParseRemote(%q) = %q, %q, %v", tc.raw, host, project, ok)
		}
	}
}
