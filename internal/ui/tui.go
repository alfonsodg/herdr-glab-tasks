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
)

const selMark = "▌"

type Model struct {
	tree    *Tree
	branch  string
	err     string
	loading bool
	detail  *Detail
	width   int
	height  int
}

func (m Model) InDetail() bool {
	return m.detail != nil
}

func (m *Model) OpenDetail(issue gitlab.Issue) {
	height := m.height - 8
	if height < 5 {
		height = 20
	}
	m.detail = &Detail{Issue: issue, height: height}
}

func (m *Model) CloseDetail() {
	m.detail = nil
}

func (m Model) DetailOffset() int {
	if m.detail == nil {
		return 0
	}
	return m.detail.offset
}

func (m *Model) DetailDown() {
	if m.detail == nil {
		return
	}
	maxOffset := len(DetailLines(m.detail.Issue, m.width)) - m.detail.height
	if maxOffset < 0 {
		maxOffset = 0
	}
	if m.detail.offset < maxOffset {
		m.detail.offset++
	}
}

func (m *Model) DetailUp() {
	if m.detail == nil || m.detail.offset <= 0 {
		return
	}
	m.detail.offset--
}

func NewModel(issues []gitlab.Issue, branchRef string) Model {
	return Model{tree: NewTree(issues), branch: branchRef}
}

func NewErrorModel(message string) Model {
	return Model{tree: NewTree(nil), err: message}
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
		if m.detail != nil {
			switch msg.String() {
			case "q", "esc":
				m.detail = nil
			case "up", "k":
				m.DetailUp()
			case "down", "j":
				m.DetailDown()
			}
			return m, nil
		}
		switch msg.String() {
		case "q", "ctrl+c", "esc":
			return m, tea.Quit
		case "up", "k":
			m.tree.MoveUp()
		case "down", "j":
			m.tree.MoveDown()
		case "tab", " ":
			m.tree.Toggle()
		case "enter":
			if row := m.tree.CurrentRow(); row != nil && row.Kind == RowIssue && row.Issue != nil {
				m.OpenDetail(*row.Issue)
			} else {
				m.tree.Toggle()
			}
		}
	case tea.WindowSizeMsg:
		m.width, m.height = msg.Width, msg.Height
		if m.detail != nil {
			m.detail.height = max(msg.Height-8, 5)
		}
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
	if m.detail != nil {
		return m.renderDetail()
	}
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
	if len(rows) == 0 && m.err == "" {
		out.WriteString(dimStyle.Render("No issues found.") + "\n")
	}
	if m.err != "" && len(rows) == 0 {
		out.WriteString(dimStyle.Render("Press q to close."))
		out.WriteString("\n")
	}
	viewHeight := m.height - 4
	if viewHeight < 3 {
		viewHeight = 20
	}
	start, end := windowRows(len(rows), m.tree.cursor, viewHeight)
	for i := start; i < end; i++ {
		row := rows[i]
		line := ""
		switch row.Kind {
		case RowHeader:
			marker := "-"
			if m.tree.collapsed[row.Column] {
				marker = "+"
			}
			line = columnStyle(row.Column).Render(fmt.Sprintf("%s %s", marker, row.Column))
		case RowIssue:
			head := fmt.Sprintf("#%d %s", row.Issue.IID, row.Issue.Title)
			line = "  " + issueStyle(row.Issue.Labels).Render(head)
			if chips := renderLabels(row.Issue.Labels); chips != "" {
				line += " " + chips
			}
		}
		if i == m.tree.cursor {
			line = selMark + " " + lipgloss.NewStyle().Bold(true).Render(line)
		} else {
			line = "  " + line
		}
		out.WriteString(line + "\n")
	}
	out.WriteString("\n")
	footer := "up/down navigate · enter detail · tab collapse · q quit"
	if m.branch != "" {
		footer += " · " + m.branch
	}
	out.WriteString(dimStyle.Render(footer))
	return out.String()
}

func (m Model) renderDetail() string {
	var out strings.Builder
	out.WriteString(groupStyle.Render("GitLab Issues"))
	out.WriteString("\n\n")
	lines := DetailLines(m.detail.Issue, m.width)
	height := m.detail.height
	if height <= 0 {
		height = 20
	}
	for _, line := range ScrollWindow(lines, m.detail.offset, height) {
		out.WriteString(line + "\n")
	}
	out.WriteString("\n")
	out.WriteString(dimStyle.Render("j/k scroll · esc back · q quit"))
	return out.String()
}

func Run(ctx context.Context, issues []gitlab.Issue, branchRef string) error {
	_ = ctx
	ForceColor()
	p := tea.NewProgram(NewModel(issues, branchRef))
	_, err := p.Run()
	return err
}

func RunError(ctx context.Context, message string) error {
	_ = ctx
	ForceColor()
	p := tea.NewProgram(NewErrorModel(message))
	_, err := p.Run()
	return err
}
