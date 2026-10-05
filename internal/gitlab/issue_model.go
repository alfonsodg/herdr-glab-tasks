// Package gitlab fetches issues through `glab api graphql`.
package gitlab

import "slices"

type Issue struct {
	IID    int      `json:"iid"`
	Title  string   `json:"title"`
	State  string   `json:"state"`
	Labels []string `json:"labels,omitempty"`
	WebURL string   `json:"web_url"`
}

func (i Issue) StatusColumn() string {
	if i.State == "closed" || slices.Contains(i.Labels, "status::done") {
		return "done"
	}
	for _, col := range []string{"todo", "review", "backlog"} {
		if slices.Contains(i.Labels, "status::"+col) {
			return col
		}
	}
	return "none"
}
