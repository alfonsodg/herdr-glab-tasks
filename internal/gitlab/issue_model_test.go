package gitlab

import "testing"

func TestIssueStatusColumn(t *testing.T) {
	cases := []struct {
		name   string
		labels []string
		state  string
		want   string
	}{
		{"todo", []string{"priority::high", "status::todo"}, "opened", "todo"},
		{"review", []string{"status::review"}, "opened", "review"},
		{"backlog", []string{"status::backlog"}, "opened", "backlog"},
		{"done label", []string{"status::done"}, "opened", "done"},
		{"closed state", nil, "closed", "done"},
		{"no status", []string{"priority::low"}, "opened", "none"},
		{"empty", nil, "opened", "none"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			issue := Issue{Labels: tc.labels, State: tc.state}
			if got := issue.StatusColumn(); got != tc.want {
				t.Fatalf("StatusColumn() = %q, want %q", got, tc.want)
			}
		})
	}
}
