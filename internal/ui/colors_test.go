package ui

import "testing"

func TestLabelColor(t *testing.T) {
	cases := []struct {
		label string
		want  string
	}{
		{"priority::critical", "1"},
		{"priority::high", "3"},
		{"priority::medium", "4"},
		{"priority::low", "8"},
		{"type::bug", "1"},
		{"type::security", "1"},
		{"type::feature", "2"},
		{"status::todo", "8"},
		{"status::review", "8"},
		{"area::api", ""},
		{"", ""},
	}
	for _, tc := range cases {
		if got := LabelColorName(tc.label); got != tc.want {
			t.Fatalf("LabelColorName(%q) = %q, want %q", tc.label, got, tc.want)
		}
	}
}
