# Changelog

All notable changes to this project are documented here. Format follows
[Keep a Changelog](https://keepachangelog.com/en/1.1.0/).

## [Unreleased]

## [0.2.0] - 2026-10-05

Interactive panel release: keyboard-driven issue tree plus keybindings.

### Added in 0.2.0

- Interactive bubbletea issue tree: cursor, collapse, filter (#10).
- `panel --print` batch mode for scripts (#10).
- `Ctrl+B i` keybinding opening the pane via `panel-open` (#8).
- Focused-workspace cwd resolution so the pane renders (#9).

### Fixed

- Pane died silently on wrong cwd; errors now print visibly (#9).

## [0.1.0] - 2026-10-05

First functional release: Herdr GitLab Issues plugin v1 scope.

### Added

- Issue panel grouped by `status::*` columns with label filter (#3).
- GraphQL issue queries via `glab` CLI: list, read, create (#2).
- Create-issue-plus-branch flow with `Ref #N` branch state line (#4).
- Plugin manifest without `[[startup]]`, link handler for issue URLs (#1).
- Enum and `work_items` URL fixes verified against live E2E (#5).

### Security

- No token management: all GitLab access delegates to `glab`.
