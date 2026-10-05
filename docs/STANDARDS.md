# STANDARDS.md

| Field | Value |
| ----- | ----- |
| **Project** | herdr-glab-tasks — Herdr plugin for GitLab Issues |
| **Version** | v0.2.0 |
| **Created** | 2026-10-05 |
| **Updated** | 2026-10-05 01:30 |
| **Author** | Alfonso de la Guarda Reyes |
| **Location** | `docs/STANDARDS.md` |

Development standards for this repository.

## Stack

- Go 1.27, module `github.com/alfonsodg/herdr-glab-tasks`.
- Single binary `bin/herdr-gitlab-issues`, built with `sh install.sh`.
- Zero third-party runtime dependencies (standard library only).
- CI: `.github/workflows/ci.yml` runs gofmt check, `go vet`, `go test`.

## Lint and format

- `gofmt -l .` must print nothing.
- `go vet ./...` must pass with zero findings.
- Markdown files must pass `markdownlint-cli`.

## Code patterns

- Small packages under `internal/`: `gitlab`, `ui`, `branch`, `repo`.
- `internal/gitlab` shells out to `glab api graphql`; never handles tokens.
- `Client.run` is injectable so tests use fixture JSON, never live `glab`.
- Errors are typed sentinels (`ErrGlabMissing`, `ErrUnauthorized`).
- Branch names follow `<type>/<scope>-#IID`; commits carry `Ref #N`.

## Testing

- TDD: failing test first, minimal implementation, refactor.
- Commands: `go test ./... -count=1`, `go vet ./...`, `go build ./...`.
- Branch tests use temp git repos with `--no-verify` where the global
  commit hook would reject fixture messages; the real repo is never mutated.
- Existing tests are inviolable: never alter assertions to make new code pass.

## Commits and issues

- Conventional commits `<type>(<scope>): <subject> (#N)` with `Ref #N` footer.
- Every change references a GitHub issue; docs changes use `type::docs`.
- Labels: one `priority::*`, one `type::*`, one `source::*` per issue.
