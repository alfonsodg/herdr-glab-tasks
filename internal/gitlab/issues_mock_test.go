package gitlab

import (
	"context"
	"testing"
)

func stubClient(response string) *Client {
	c := NewClient("glab", "git.example.com")
	c.run = func(ctx context.Context, glab string, args ...string) ([]byte, error) {
		return []byte(response), nil
	}
	return c
}

func TestListIssuesParsesResponse(t *testing.T) {
	resp := `{"data":{"project":{"issues":{"nodes":[{"iid":"3","title":"Third","state":"opened","webUrl":"https://git.example.com/g/p/-/issues/3","description":"","updatedAt":"2026-10-01T00:00:00Z","labels":{"nodes":[{"title":"status::todo"}]}}]}}}}`
	issues, err := stubClient(resp).ListIssues(context.Background(), "g/p", "opened", "")
	if err != nil {
		t.Fatalf("ListIssues: %v", err)
	}
	if len(issues) != 1 || issues[0].IID != 3 || issues[0].StatusColumn() != "todo" {
		t.Fatalf("issues = %+v", issues)
	}
}

func TestLatestPipelineParsesResponse(t *testing.T) {
	resp := `[{"status":"success","ref":"feature/auth-#42","web_url":"https://git.example.com/g/p/-/pipelines/7"}]`
	pipeline, err := stubClient(resp).LatestPipeline(context.Background(), "g/p", "feature/auth-#42")
	if err != nil {
		t.Fatalf("LatestPipeline: %v", err)
	}
	if pipeline.Status != "success" || pipeline.Ref != "feature/auth-#42" {
		t.Fatalf("pipeline = %+v", pipeline)
	}
}

func TestLatestPipelineWithoutResults(t *testing.T) {
	pipeline, err := stubClient(`[]`).LatestPipeline(context.Background(), "g/p", "feature/auth-#42")
	if err != nil {
		t.Fatalf("LatestPipeline: %v", err)
	}
	if pipeline.Status != "none" || pipeline.Ref != "feature/auth-#42" {
		t.Fatalf("pipeline = %+v; want none for the branch", pipeline)
	}
}

func TestListIssuesProjectNotFound(t *testing.T) {
	resp := `{"data":{"project":null}}`
	if _, err := stubClient(resp).ListIssues(context.Background(), "g/missing", "opened", ""); err == nil {
		t.Fatal("expected error for missing project")
	}
}

func TestGetIssueNotFound(t *testing.T) {
	resp := `{"data":{"project":{"issue":null}}}`
	if _, err := stubClient(resp).GetIssue(context.Background(), "g/p", 999); err == nil {
		t.Fatal("expected error for missing issue")
	}
}
