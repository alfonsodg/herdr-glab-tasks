// Command herdr-gitlab-issues is the entry point for the plugin commands.
package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/alfonsodg/herdr-glab-tasks/internal/branch"
	"github.com/alfonsodg/herdr-glab-tasks/internal/gitlab"
	"github.com/alfonsodg/herdr-glab-tasks/internal/pane"
	"github.com/alfonsodg/herdr-glab-tasks/internal/repo"
	"github.com/alfonsodg/herdr-glab-tasks/internal/ui"
)

const usage = "usage: herdr-gitlab-issues <panel|panel-open|new> [flags]"

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, usage)
		os.Exit(2)
	}
	if err := run(context.Background(), os.Args[1], os.Args[2:]); err != nil {
		fmt.Fprintf(os.Stderr, "herdr-gitlab-issues %s: %v\n", os.Args[1], err)
		os.Exit(1)
	}
}

func run(ctx context.Context, cmd string, args []string) error {
	switch cmd {
	case "panel":
		return runPanel(ctx, args)
	case "panel-open":
		return runPanelOpen(ctx, args)
	case "new":
		return runNew(ctx, args)
	default:
		return fmt.Errorf("unknown command: %s", cmd)
	}
}

func runPanelOpen(ctx context.Context, args []string) error {
	fs := flag.NewFlagSet("panel-open", flag.ContinueOnError)
	cwd := fs.String("cwd", "", "workspace directory holding the git remote")
	if err := fs.Parse(args); err != nil {
		return err
	}
	if *cwd == "" {
		*cwd = os.Getenv("HERDR_PLUGIN_CWD")
	}
	return (&pane.Opener{}).Open(ctx, "alfonsodg.herdr-gitlab-issues", "issues", *cwd)
}

func runPanel(ctx context.Context, args []string) error {
	fs := flag.NewFlagSet("panel", flag.ContinueOnError)
	label := fs.String("label", "", "filter by label")
	state := fs.String("state", "opened", "issue state")
	cwd := fs.String("cwd", "", "workspace directory holding the git remote")
	if err := fs.Parse(args); err != nil {
		return err
	}
	dir := *cwd
	if dir == "" {
		dir = os.Getenv("HERDR_PLUGIN_CWD")
	}
	if dir == "" {
		dir = "."
	}
	host, project, err := workspaceProject(dir)
	if err != nil {
		fmt.Printf("GitLab Issues: %v\n", err)
		fmt.Println("Open this pane from a workspace backed by a GitLab remote.")
		return err
	}
	client := gitlab.NewClient("glab", host)
	issues, err := client.ListIssues(ctx, project, *state, *label)
	if err != nil {
		return err
	}
	issues = ui.FilterByLabel(issues, *label)
	fmt.Print(ui.RenderPanel(ui.GroupByStatus(issues)))
	if iid, ok := branch.IssueRef(dir); ok {
		fmt.Printf("branch: Ref #%d\n", iid)
	}
	return nil
}

func runNew(ctx context.Context, args []string) error {
	fs := flag.NewFlagSet("new", flag.ContinueOnError)
	title := fs.String("title", "", "issue title")
	description := fs.String("description", "", "issue description")
	labels := fs.String("labels", "", "comma-separated labels")
	branchType := fs.String("type", "feature", "branch type")
	scope := fs.String("scope", "tasks", "branch scope")
	cwd := fs.String("cwd", "", "workspace directory holding the git remote")
	if err := fs.Parse(args); err != nil {
		return err
	}
	if strings.TrimSpace(*title) == "" {
		return fmt.Errorf("title is required")
	}
	dir := *cwd
	if dir == "" {
		dir = os.Getenv("HERDR_PLUGIN_CWD")
	}
	if dir == "" {
		dir = "."
	}
	host, project, err := workspaceProject(dir)
	if err != nil {
		return err
	}
	client := gitlab.NewClient("glab", host)
	var labelList []string
	if *labels != "" {
		for _, l := range strings.Split(*labels, ",") {
			if trimmed := strings.TrimSpace(l); trimmed != "" {
				labelList = append(labelList, trimmed)
			}
		}
	}
	issue, err := client.CreateIssue(ctx, project, *title, *description, labelList)
	if err != nil {
		return err
	}
	if issue.IID == 0 {
		return fmt.Errorf("could not determine created issue IID")
	}
	if err := branch.CreateIssueBranch(dir, *branchType, *scope, issue.IID); err != nil {
		return err
	}
	fmt.Printf("created #%d %s on branch %s\n", issue.IID, issue.Title, branch.Name(*branchType, *scope, issue.IID))
	return nil
}

func workspaceProject(dir string) (host, project string, err error) {
	out, rerr := exec.Command("git", "-C", dir, "remote", "get-url", "origin").Output()
	if rerr != nil {
		return "", "", fmt.Errorf("no git origin remote: %v", rerr)
	}
	host, project, ok := repo.ParseRemote(string(out))
	if !ok {
		return "", "", fmt.Errorf("cannot parse origin remote")
	}
	return host, project, nil
}
