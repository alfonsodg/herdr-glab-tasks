# Changelog

All notable changes to this project are documented here. Format follows
[Keep a Changelog](https://keepachangelog.com/en/1.1.0/).

## [Unreleased]

### Added

- Prebuilt release packages (linux/macos, amd64/arm64) published by the
  release workflow; `install.sh` downloads the matching binary when the
  machine has no Go toolchain (#22).
- Branch footer now includes the latest GitLab CI status for the referenced
  issue branch (#4).
- Interactive new-issue form for the Herdr keybinding, while preserving the
  flag-based CLI flow (#8).

## [0.3.0] - 2026-10-06

Config, navigation and resilience release: interactive polish plus plugin
id rename.

### Added in 0.3.0

- Per-column header colors: todo yellow, review magenta, backlog gray,
  done green, none cyan (#14).
- Priority-colored issue titles: critical bright red, high bright
  yellow, medium bright cyan, low gray (#14).
- Scrollable issue detail window on `enter` with `j/k` scroll and
  `esc` back, preserving the tree cursor (#12).
- Label chips keep their family colors inside the selected row (#14).
- Viewport scrolling for the tree: the window follows the cursor (#17).
- Plugin `config.toml` defaults: `state`, `label`, `branch_type`,
  `branch_scope`; CLI flags win over config (#18).

### Fixed in 0.3.0

- Color profile forced at TUI start so the Herdr pane renders ANSI
  even when it advertises dumb/no-TTY; `NO_COLOR` still respected (#13).
- Detail scroll window clamped to the terminal height (#12).
- Interactive panel stays open showing the reason when the focused
  pane is not a GitLab workspace (#16).

### Changed in 0.3.0

- Plugin id renamed `alfonsodg.herdr-gitlab-issues` ->
  `herdr-gitlab-issues`; keybindings and config dir move with it (#19).

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
