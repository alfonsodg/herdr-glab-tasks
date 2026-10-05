package ui

import (
	"context"
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"

	"github.com/alfonsodg/herdr-glab-tasks/internal/gitlab"
)

var (
	groupStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("6"))
	dimStyle   = lipgloss.NewStyle().Faint(true)
	errStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("1"))
	selStyle   = lipgloss.NewStyle().Background(lipgloss.Color("238")).Foreground(lipgloss.Color("231"))
)

type Model struct {
	tree    *Tree
	branch  string
	err     string
	loading bool
}

func NewModel(issues []gitlab.Issue, branchRef string) Model {
	return Model{tree: NewTree(issues), branch: branchRef}
}

func (m Model) Init() tea.Cmd {
	return nil
}

type IssuesLoaded struct {
	Issues []gitlab.Issue
	Branch string
	Err    error
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c", "esc":
			return m, tea.Quit
		case "up", "k":
			m.tree.MoveUp()
		case "down", "j":
			m.tree.MoveDown()
		case "enter", "tab", " ":
			m.tree.Toggle()
		}
	case tea.WindowSizeMsg:
		_ = msg
	case IssuesLoaded:
		if msg.Err != nil {
			m.err = msg.Err.Error()
			m.loading = false
			return m, nil
		}
		m.tree = NewTree(msg.Issues)
		m.branch = msg.Branch
		m.loading = false
	}
	return m, nil
}

func (m Model) View() tea.View {
	v := tea.NewView(m.render())
	v.AltScreen = true
	return v
}

func (m Model) render() string {
	var out strings.Builder
	out.WriteString(groupStyle.Render("GitLab Issues"))
	if m.tree.filter != "" {
		fmt.Fprintf(&out, "  [%s]", m.tree.filter)
	}
	out.WriteString("\n\n")
	if m.err != "" {
		out.WriteString(errStyle.Render(m.err) + "\n\n")
	}
	rows := m.tree.Rows()
	if len(rows) == 0 {
		out.WriteString(dimStyle.Render("No issues found.") + "\n")
	}
	for i, row := range rows {
		line := ""
		switch row.Kind {
		case RowHeader:
			marker := "-"
			if m.tree.collapsed[row.Column] {
				marker = "+"
			}
			line = groupStyle.Render(fmt.Sprintf("%s %s", marker, row.Column))
		case RowIssue:
			labels := ""
			if len(row.Issue.Labels) > 0 {
				labels = " [" + strings.Join(row.Issue.Labels, ", ") + "]"
			}
			line = fmt.Sprintf("  #%d %s%s", row.Issue.IID, row.Issue.Title, labels)
		}
		if i == m.tree.cursor {
			line = selStyle.Render(line)
		}
		out.WriteString(line + "\n")
	}
	out.WriteString("\n")
	footer := "up/down navigate · enter collapse · q quit"
	if m.branch != "" {
		footer += " · " + m.branch
	}
	out.WriteString(dimStyle.Render(footer))
	return out.String()
}

func Run(ctx context.Context, issues []gitlab.Issue, branchRef string) error {
	_ = ctx
	p := tea.NewProgram(NewModel(issues, branchRef))
	_, err := p.Run()
	return err
}
