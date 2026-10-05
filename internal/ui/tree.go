package ui

import "github.com/alfonsodg/herdr-glab-tasks/internal/gitlab"

type RowKind int

const (
	RowHeader RowKind = iota
	RowIssue
)

type Row struct {
	Kind   RowKind
	Column string
	Issue  *gitlab.Issue
}

type Tree struct {
	issues    []gitlab.Issue
	collapsed map[string]bool
	filter    string
	cursor    int
}

func NewTree(issues []gitlab.Issue) *Tree {
	return &Tree{issues: issues, collapsed: map[string]bool{}}
}

func (t *Tree) SetFilter(label string) {
	t.filter = label
	t.cursor = 0
}

func (t *Tree) Toggle() {
	rows := t.Rows()
	if t.cursor < 0 || t.cursor >= len(rows) {
		return
	}
	row := rows[t.cursor]
	t.collapsed[row.Column] = !t.collapsed[row.Column]
	if t.cursor >= len(t.Rows()) {
		t.cursor = len(t.Rows()) - 1
	}
}

func (t *Tree) Rows() []Row {
	issues := t.issues
	if t.filter != "" {
		issues = FilterByLabel(issues, t.filter)
	}
	grouped := GroupByStatus(issues)
	var rows []Row
	for _, col := range columns {
		list := grouped[col]
		if len(list) == 0 {
			continue
		}
		rows = append(rows, Row{Kind: RowHeader, Column: col})
		if t.collapsed[col] {
			continue
		}
		for i := range list {
			rows = append(rows, Row{Kind: RowIssue, Column: col, Issue: &list[i]})
		}
	}
	return rows
}

func (t *Tree) RowCount() int {
	return len(t.Rows())
}

func (t *Tree) Cursor() int {
	return t.cursor
}

func (t *Tree) MoveDown() {
	if t.cursor < len(t.Rows())-1 {
		t.cursor++
	}
}

func (t *Tree) MoveUp() {
	if t.cursor > 0 {
		t.cursor--
	}
}

func (t *Tree) CurrentRow() *Row {
	rows := t.Rows()
	if t.cursor < 0 || t.cursor >= len(rows) {
		return nil
	}
	return &rows[t.cursor]
}
