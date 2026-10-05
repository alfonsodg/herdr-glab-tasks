package gitlab

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type issueNode struct {
	IID         string `json:"iid"`
	Title       string `json:"title"`
	State       string `json:"state"`
	WebURL      string `json:"webUrl"`
	Description string `json:"description"`
	UpdatedAt   string `json:"updatedAt"`
	Author      *struct {
		Username string `json:"username"`
	} `json:"author"`
	Labels *struct {
		Nodes []struct {
			Title string `json:"title"`
		} `json:"nodes"`
	} `json:"labels"`
}

const issueFragment = `
fragment issue on Issue {
  iid title state webUrl description updatedAt
  author { username }
  labels(first: 20) { nodes { title } }
}
`

func buildListQuery(project, state, label string) string {
	p, _ := json.Marshal(project)
	switch state {
	case "opened", "closed", "all":
	default:
		state = "opened"
	}
	var filter string
	if label != "" {
		l, _ := json.Marshal(label)
		filter = fmt.Sprintf(", labelName: [%s]", l)
	}
	return fmt.Sprintf("query { project(fullPath: %s) { issues(state: %s%s, first: 50) { nodes { ...issue } } } }\n%s", p, state, filter, issueFragment)
}

const oneIssueQuery = `
query($project: ID!, $iid: String!) {
  project(fullPath: $project) { issue(iid: $iid) { ...issue } }
}
` + issueFragment

func convertIssue(n issueNode) Issue {
	iid, _ := strconv.Atoi(n.IID)
	updated, _ := time.Parse(time.RFC3339, n.UpdatedAt)
	issue := Issue{
		IID:         iid,
		Title:       n.Title,
		State:       n.State,
		WebURL:      n.WebURL,
		Description: n.Description,
		UpdatedAt:   updated,
	}
	if n.Author != nil {
		issue.Author = n.Author.Username
	}
	if n.Labels != nil {
		for _, l := range n.Labels.Nodes {
			issue.Labels = append(issue.Labels, l.Title)
		}
	}
	return issue
}

func contains(s, sub string) bool {
	return strings.Contains(s, sub)
}
