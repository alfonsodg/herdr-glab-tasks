package gitlab

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

var (
	ErrGlabMissing  = errors.New("glab not found")
	ErrUnauthorized = errors.New("glab is not authenticated")
)

type Client struct {
	glab string
	host string
	run  func(ctx context.Context, glab string, args ...string) ([]byte, error)
}

type Pipeline struct {
	Status string `json:"status"`
	Ref    string `json:"ref"`
	WebURL string `json:"web_url"`
}

func NewClient(glab, host string) *Client {
	return &Client{glab: glab, host: host, run: defaultRun}
}

func defaultRun(ctx context.Context, glab string, args ...string) ([]byte, error) {
	cmd := exec.CommandContext(ctx, glab, args...)
	out, err := cmd.Output()
	if err != nil {
		return out, err
	}
	return out, nil
}

func (c *Client) ListIssues(ctx context.Context, project, state, label string) ([]Issue, error) {
	var data struct {
		Project *struct {
			Issues struct {
				Nodes []issueNode `json:"nodes"`
			} `json:"issues"`
		} `json:"project"`
	}
	if err := c.graphql(ctx, buildListQuery(project, state, label), nil, &data); err != nil {
		return nil, err
	}
	if data.Project == nil {
		return nil, fmt.Errorf("project %s not found", project)
	}
	issues := make([]Issue, 0, len(data.Project.Issues.Nodes))
	for _, n := range data.Project.Issues.Nodes {
		issues = append(issues, convertIssue(n))
	}
	return issues, nil
}

func (c *Client) GetIssue(ctx context.Context, project string, iid int) (Issue, error) {
	var data struct {
		Project *struct {
			Issue *issueNode `json:"issue"`
		} `json:"project"`
	}
	vars := []string{"-f", "project=" + project, "-f", "iid=" + strconv.Itoa(iid)}
	if err := c.graphql(ctx, oneIssueQuery, vars, &data); err != nil {
		return Issue{}, err
	}
	if data.Project == nil || data.Project.Issue == nil {
		return Issue{}, fmt.Errorf("%s#%d not found", project, iid)
	}
	return convertIssue(*data.Project.Issue), nil
}

func (c *Client) CreateIssue(ctx context.Context, project, title, description string, labels []string) (Issue, error) {
	args := []string{"issue", "create", "--repo", project, "--title", title}
	if description != "" {
		args = append(args, "--description", description)
	}
	for _, l := range labels {
		args = append(args, "--label", l)
	}
	out, err := c.run(ctx, c.glab, args...)
	if err != nil {
		return Issue{}, classify(err)
	}
	if iid := parseCreatedIID(string(out)); iid != 0 {
		return c.GetIssue(ctx, project, iid)
	}
	return Issue{Title: title, Description: description, Labels: labels, State: "opened"}, nil
}

func (c *Client) LatestPipeline(ctx context.Context, project, branch string) (Pipeline, error) {
	args := []string{
		"ci", "list", "--repo", project, "--ref", branch,
		"--order", "updated_at", "--sort", "desc", "--per-page", "1",
		"--output", "json",
	}
	out, err := c.run(ctx, c.glab, args...)
	if err != nil {
		return Pipeline{}, classify(err)
	}
	var pipelines []Pipeline
	if err := json.Unmarshal(out, &pipelines); err != nil {
		return Pipeline{}, fmt.Errorf("parse pipeline response: %w", err)
	}
	if len(pipelines) == 0 {
		return Pipeline{Status: "none", Ref: branch}, nil
	}
	return pipelines[0], nil
}

var createdURLPattern = regexp.MustCompile(`/-/work_items/(\d+)|/-/issues/(\d+)`)

func parseCreatedIID(out string) int {
	m := createdURLPattern.FindStringSubmatch(out)
	if m == nil {
		return 0
	}
	for _, g := range m[1:] {
		if g == "" {
			continue
		}
		if iid, err := strconv.Atoi(g); err == nil {
			return iid
		}
	}
	return 0
}

func (c *Client) graphql(ctx context.Context, query string, vars []string, out any) error {
	args := append([]string{"api", "graphql", "--hostname", c.host, "-f", "query=" + query}, vars...)
	stdout, err := c.run(ctx, c.glab, args...)
	if err != nil {
		return classify(err)
	}
	var resp struct {
		Data   json.RawMessage `json:"data"`
		Errors []struct {
			Message string `json:"message"`
		} `json:"errors"`
	}
	if err := json.Unmarshal(stdout, &resp); err != nil {
		return fmt.Errorf("parse graphql response: %w", err)
	}
	if len(resp.Errors) > 0 {
		msgs := make([]string, 0, len(resp.Errors))
		for _, e := range resp.Errors {
			msgs = append(msgs, e.Message)
		}
		return fmt.Errorf("graphql: %s", strings.Join(msgs, "; "))
	}
	return json.Unmarshal(resp.Data, out)
}

func classify(err error) error {
	if errors.Is(err, exec.ErrNotFound) || errors.Is(err, fs.ErrNotExist) {
		return fmt.Errorf("%w: %v", ErrGlabMissing, err)
	}
	return err
}
