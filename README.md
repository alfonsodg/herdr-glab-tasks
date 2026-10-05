# herdr-glab-tasks

Herdr plugin for GitLab Issues integration: workspace issue panel grouped by
`status::*`, create-issue-plus-branch action, branch state line, and issue
URL link handling.

Derived from [hlouis/herdr-glab](https://github.com/hlouis/herdr-glab) (MIT),
which covers merge requests. This plugin covers Issues, which no marketplace
plugin addresses.

## Features

- Issue panel grouped by `todo`, `review`, `backlog`, `done`, `none` columns.
- Label filter (`--label`) and state selector (`--state`).
- Create issue plus branch in one step (`new` command).
- Branch state line showing the `Ref #N` issue of the current branch.
- Link handler opening GitLab issue URLs from Herdr panes.

## Stack

- Go 1.27, single binary, no runtime dependencies.
- GitLab access exclusively through the `glab` CLI (no tokens managed).
- See [docs/STANDARDS.md](docs/STANDARDS.md) for conventions.

## Quick start

```sh
herdr plugin link /path/to/herdr-glab-tasks --enabled
sh install.sh
bin/herdr-gitlab-issues panel
bin/herdr-gitlab-issues panel --state all --label priority::medium
bin/herdr-gitlab-issues new --title "feat(scope): subject (#N)"
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
