package gitlab

import "testing"

func TestParseCreatedIID(t *testing.T) {
	cases := []struct {
		out  string
		want int
	}{
		{"- Creating issue\nhttps://scovil.labtau.com/g/p/-/work_items/53\n", 53},
		{"https://git.example.com/g/p/-/issues/42", 42},
		{"no url here", 0},
		{"", 0},
	}
	for _, tc := range cases {
		if got := parseCreatedIID(tc.out); got != tc.want {
			t.Fatalf("parseCreatedIID(%q) = %d, want %d", tc.out, got, tc.want)
		}
	}
}
