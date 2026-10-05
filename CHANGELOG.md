# Changelog

All notable changes to this project are documented here. Format follows
[Keep a Changelog](https://keepachangelog.com/en/1.1.0/).

## [Unreleased]

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
