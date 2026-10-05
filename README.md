# herdr-glab-tasks

Herdr plugin for GitLab Issues integration: workspace issue panel grouped by
`status::*`, create-issue-plus-branch action, branch pipeline state, and issue
URL link handling.

Derived from [hlouis/herdr-glab](https://github.com/hlouis/herdr-glab) (MIT),
which covers merge requests. This plugin covers Issues, which no marketplace
plugin addresses.

## Status

Bootstrap (v0.1.0 scaffold): `herdr-plugin.toml` manifest, Go module,
`internal/gitlab` issue model and client stub, `internal/ui` grouping helper,
and `cmd/herdr-gitlab-issues` entry point. Panel rendering and GraphQL queries
are follow-up work tracked in GitHub issues.

## Security

The plugin never manages tokens. All GitLab access delegates to the `glab` CLI
already authenticated on the machine. The manifest declares no `[[startup]]`
command, so nothing runs without an explicit user action.

## Build

```sh
sh install.sh
```

## Test

```sh
go test ./...
go vet ./...
```
