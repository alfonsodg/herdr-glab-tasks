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
	"github.com/alfonsodg/herdr-glab-tasks/internal/config"
	"github.com/alfonsodg/herdr-glab-tasks/internal/gitlab"
	"github.com/alfonsodg/herdr-glab-tasks/internal/pane"
	"github.com/alfonsodg/herdr-glab-tasks/internal/repo"
	"github.com/alfonsodg/herdr-glab-tasks/internal/ui"
)

const usage = "usage: herdr-gitlab-issues <panel|panel-open|new> [flags]"

const pluginID = "herdr-gitlab-issues"

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
	return (&pane.Opener{}).Open(ctx, pluginID, "issues", *cwd)
}

func runPanel(ctx context.Context, args []string) error {
	fs := flag.NewFlagSet("panel", flag.ContinueOnError)
	label := fs.String("label", "", "filter by label (overrides config)")
	state := fs.String("state", "", "issue state: opened, closed, all (overrides config)")
	cwd := fs.String("cwd", "", "workspace directory holding the git remote")
	printOut := fs.Bool("print", false, "print static panel and exit instead of interactive TUI")
	if err := fs.Parse(args); err != nil {
		return err
	}
	cfg := loadConfig()
	useState := *state
	if useState == "" {
		useState = cfg.State
	}
	useLabel := *label
	if useLabel == "" {
		useLabel = cfg.Label
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
		msg := fmt.Sprintf("Not a GitLab workspace: %s\n%s\n\nOpen this pane from a repository with a GitLab remote.", dir, err)
		if *printOut {
			fmt.Println(msg)
			return err
		}
		return ui.RunError(ctx, msg)
	}
	client := gitlab.NewClient("glab", host)
	issues, err := client.ListIssues(ctx, project, useState, useLabel)
	if err != nil {
		if *printOut {
			return err
		}
		return ui.RunError(ctx, fmt.Sprintf("Cannot list issues for %s: %v", project, err))
	}
	issues = ui.FilterByLabel(issues, useLabel)
	branchRef := ""
	if iid, ok := branch.IssueRef(dir); ok {
		branchRef = fmt.Sprintf("branch: Ref #%d", iid)
	}
	if *printOut {
		fmt.Print(ui.RenderPanel(ui.GroupByStatus(issues)))
		if branchRef != "" {
			fmt.Println(branchRef)
		}
		return nil
	}
	return ui.Run(ctx, issues, branchRef)
}

func runNew(ctx context.Context, args []string) error {
	fs := flag.NewFlagSet("new", flag.ContinueOnError)
	title := fs.String("title", "", "issue title")
	description := fs.String("description", "", "issue description")
	labels := fs.String("labels", "", "comma-separated labels")
	branchType := fs.String("type", "", "branch type (overrides config)")
	scope := fs.String("scope", "", "branch scope (overrides config)")
	cwd := fs.String("cwd", "", "workspace directory holding the git remote")
	if err := fs.Parse(args); err != nil {
		return err
	}
	if strings.TrimSpace(*title) == "" {
		return fmt.Errorf("title is required")
	}
	cfg := loadConfig()
	useType := *branchType
	if useType == "" {
		useType = cfg.BranchType
	}
	useScope := *scope
	if useScope == "" {
		useScope = cfg.BranchScope
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
	if err := branch.CreateIssueBranch(dir, useType, useScope, issue.IID); err != nil {
		return err
	}
	fmt.Printf("created #%d %s on branch %s\n", issue.IID, issue.Title, branch.Name(useType, useScope, issue.IID))
	return nil
}

func loadConfig() config.Config {
	dir := os.Getenv("HERDR_PLUGIN_CONFIG_DIR")
	if dir == "" {
		if out, err := exec.Command("herdr", "plugin", "config-dir", pluginID).Output(); err == nil {
			dir = strings.TrimSpace(string(out))
		}
	}
	return config.Load(dir)
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
