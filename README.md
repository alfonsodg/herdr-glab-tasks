# herdr-glab-tasks

Herdr plugin for GitLab Issues integration: workspace issue panel grouped by
`status::*`, create-issue-plus-branch action, branch state line, and issue
URL link handling.

Derived from [hlouis/herdr-glab](https://github.com/hlouis/herdr-glab) (MIT),
which covers merge requests. This plugin covers Issues, which no marketplace
plugin addresses.

## Features

- Interactive issue tree grouped by `todo`, `review`, `backlog`, `done`,
  `none` columns with collapse and cursor navigation.
- Color scheme per column and per priority: headers colored by status,
  issue titles colored by `priority::*`, label chips per family.
- Scrollable detail window on `enter` (title, author, date, URL,
  labels, description) with `j/k` scroll and `esc` back.
- Label filter (`--label`) and state selector (`--state`).
- Batch mode (`panel --print`) for scripts keeping the classic output.
- Create issue plus branch in one step (`new` command).
- Branch state line showing the `Ref #N` issue and CI status of the current branch.
- Link handler opening GitLab issue URLs from Herdr panes.

## Keys

| Keys | Action |
| ---- | ------ |
| `Ctrl+B i` | Open the GitLab Issues pane |
| `Ctrl+B Shift+I` | Open the new issue plus branch form |

Bound in the local Herdr config; the `i` key is free (no overlap with
annotate, crabbox, file-viewer, or telegram bindings).

## Stack

- Go 1.27, single binary, TUI via bubbletea v2 and lipgloss v2.
- GitLab access exclusively through the `glab` CLI (no tokens managed).
- See [docs/STANDARDS.md](docs/STANDARDS.md) for conventions.

## Quick start

Install from the marketplace index or straight from GitHub:

```sh
herdr plugin install alfonsodg/herdr-glab-tasks
```

Or link a local clone for development:

```sh
herdr plugin link /path/to/herdr-glab-tasks --enabled
sh install.sh
bin/herdr-gitlab-issues panel --state all --label priority::medium
```

## Security

The plugin never manages tokens. All GitLab access delegates to the `glab`
CLI already authenticated on the machine. The manifest declares no
`[[startup]]` command, so nothing runs without an explicit user action.

## Docs

- [docs/INSTALL.md](docs/INSTALL.md): prerequisites, install, verification.
- [docs/USER_GUIDE.md](docs/USER_GUIDE.md): panel, filters, new flow, FAQ.
- [docs/STANDARDS.md](docs/STANDARDS.md): stack, linters, patterns, testing.
- [CHANGELOG.md](CHANGELOG.md): version history.

## Herdr references

- [Herdr plugin docs](https://herdr.dev/docs/plugins/): manifest model,
  trust rules (listings are unreviewed, pin with `--ref`), and install flow.
- [Herdr plugin marketplace](https://herdr.dev/plugins/): the index this
  plugin targets once released.
- Base project [hlouis/herdr-glab](https://github.com/hlouis/herdr-glab)
  (MIT): MR panel this plugin complements with Issues coverage.
