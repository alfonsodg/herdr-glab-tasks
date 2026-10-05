# INSTALL.md

| Field | Value |
| ----- | ----- |
| **Project** | herdr-glab-tasks — Herdr plugin for GitLab Issues |
| **Version** | v0.2.0 |
| **Created** | 2026-10-05 |
| **Updated** | 2026-10-05 11:30 |
| **Author** | Alfonso de la Guarda Reyes |
| **Location** | `docs/INSTALL.md` |

How to install the GitLab Issues plugin in Herdr.

## Prerequisites

- Herdr 0.9.0 or newer (see [Herdr plugin docs][herdr-plugins]).
- Go toolchain for the local build.
- `glab` CLI installed and authenticated (`glab auth status` shows
  your GitLab host). The plugin never asks for tokens: it reuses
  the `glab` session already on the machine.
- A local clone of this repository.

## Steps

1. Link the local plugin into Herdr:

   ```sh
   herdr plugin link /path/to/herdr-glab-tasks --enabled
   ```

   Expected: JSON output with `"plugin_id": "alfonsodg.herdr-gitlab-issues"`
   and `"enabled": true`.

2. Build the binary (pulls bubbletea/lipgloss TUI deps on first run):

   ```sh
   sh install.sh
   ```

   Expected: `built bin/herdr-gitlab-issues with go version ...`.

3. Bind the keys in the local Herdr config (`~/.config/herdr/config.toml`,
   never in git):

   ```toml
   [[keys.command]]
   key = "prefix+i"
   type = "plugin_action"
   command = "alfonsodg.herdr-gitlab-issues.issues"
   description = "GitLab Issues del workspace"

   [[keys.command]]
   key = "prefix+shift+i"
   type = "plugin_action"
   command = "alfonsodg.herdr-gitlab-issues.new-issue"
   description = "Nuevo issue + rama"
   ```

   The `i` key is free: no overlap with annotate, crabbox, file-viewer,
   or telegram bindings. Reload Herdr so it picks the keys up.

4. Confirm Herdr lists the plugin:

   ```sh
   herdr plugin list
   ```

   Expected: a line with `alfonsodg.herdr-gitlab-issues (GitLab Issues)`.

## Verification

Press `Ctrl+B i` from a workspace whose `origin` remote points to GitLab.
Expected: the interactive issue tree stays open with grouped sections.
Batch equivalent for scripts:

```sh
bin/herdr-gitlab-issues panel --print
```

For usage details, see [USER_GUIDE.md](USER_GUIDE.md).

## Troubleshooting

| Symptom | Cause | Fix |
| ------- | ----- | --- |
| `not implemented: panel` | Old binary | Re-run `sh install.sh` |
| `no git origin remote` | Outside a git clone | Move into a GitLab clone |
| Bad remote URL | Odd remote format | Check remote with `git remote -v` |
| `glab not found` | Missing from `PATH` | Install `glab` and re-run |
| GraphQL `401` / logged out | Session expired | Run `glab auth login` again |
| Empty panel | No match for state/filter | Try `--state all`, drop `--label` |

[herdr-plugins]: https://herdr.dev/docs/plugins/
