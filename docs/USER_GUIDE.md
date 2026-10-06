# USER_GUIDE.md

| Field | Value |
| ----- | ----- |
| **Project** | herdr-glab-tasks — Herdr plugin for GitLab Issues |
| **Version** | v0.3.0 |
| **Created** | 2026-10-05 |
| **Updated** | 2026-10-06 01:30 |
| **Author** | Alfonso de la Guarda Reyes |
| **Location** | `docs/USER_GUIDE.md` |

How to use the GitLab Issues plugin from Herdr. For setup, see
[INSTALL.md](INSTALL.md).

## Panel: see workspace issues

Press `Ctrl+B i` to open the **GitLab Issues** pane in Herdr. It runs the
`issues` action, which opens the pane entry with an interactive tree that
stays alive until you quit it.

### Tree navigation

| Keys | Action |
| ---- | ------ |
| `j` / `k` or arrows | Move the cursor |
| `enter` | Open the issue detail window |
| `tab` / `space` | Collapse or expand a group |
| `q` / `esc` / `ctrl+c` | Quit the panel |

Group headers show `-` (expanded) or `+` (collapsed). The footer shows
the current branch issue and CI status when the latest commit carries a
`Ref #N` reference.

### Colors

| Element | Color |
| ------- | ----- |
| `todo` header | Bright yellow |
| `review` header | Bright magenta |
| `backlog` header | Gray |
| `done` header | Green |
| `none` header | Cyan |
| Issue title `priority::critical` | Bright red |
| Issue title `priority::high` | Bright yellow |
| Issue title `priority::medium` | Bright cyan |
| Issue title `priority::low` | Gray |

Label chips keep their own family colors. The color profile is forced
on pane start so the palette renders even in terminals advertised as
dumb; `NO_COLOR` still disables colors when you set it.

### Issue detail window

Press `enter` on an issue to open its detail window: title, IID, state,
author, date, URL, labels, and the full description.

| Keys | Action |
| ---- | ------ |
| `j` / `k` or arrows | Scroll the description |
| `esc` / `q` | Back to the tree (cursor preserved) |

Long descriptions scroll in a window sized to the terminal height.

### Batch mode for scripts

```sh
bin/herdr-gitlab-issues panel --print
```

Prints the classic grouped output and exits. Use it in scripts; the pane
itself always runs the interactive tree.

## Filter by state and label

```sh
bin/herdr-gitlab-issues panel --state all
bin/herdr-gitlab-issues panel --label priority::high
bin/herdr-gitlab-issues panel --state all --label type::bug
```

States are `opened` (default), `closed`, and `all`. The label flag
matches one GitLab label exactly, for example `status::review`.

## Configuration

Defaults live in `config.toml` inside the plugin config directory
(`herdr plugin config-dir herdr-gitlab-issues`):

| Key | Default | Meaning |
| --- | ------- | ------- |
| `state` | `opened` | Initial issue state: `opened`, `closed`, `all` |
| `label` | empty | Initial label filter |
| `branch_type` | `feature` | Branch type used by the `new` flow |
| `branch_scope` | `tasks` | Branch scope used by the `new` flow |

Precedence: CLI flag > `config.toml` > built-in default. Example:

```toml
state = "all"
label = "priority::high"
branch_type = "fix"
branch_scope = "auth"
```

## Create issue plus branch

Press `Ctrl+B Shift+I` for the **Nuevo issue + rama** action. From the
terminal the same flow is:

```sh
bin/herdr-gitlab-issues new --title "feat(auth): login (#N)" \
  --labels "priority::high,type::feature,source::user" \
  --type feature --scope auth
```

Expected: `created #42 ... on branch feature/auth-#42`. Branch names
always follow `<type>/<scope>-#IID`, and the tree must be clean
(commit or stash first) or the command refuses to run.

## Open issues from links

Ctrl+click any GitLab issue URL inside a Herdr pane. The link handler
matches `.../-/issues/<id>` and opens the issues action.

## FAQ

**Does the plugin store my GitLab token?**
No. It shells out to `glab`, which already holds your session.

**Why is my panel empty?**
Either the project has no issues in that state, or the label filter
matches nothing. Try `--state all` first.

**Can I edit or close issues from the panel?**
No. The panel is read plus create. Edit, comment, and merge in GitLab.

**Which Herdr version do I need?**
0.9.0 or newer, per `min_herdr_version` in `herdr-plugin.toml`.
Background reading: [Herdr plugin docs][herdr-plugins].

[herdr-plugins]: https://herdr.dev/docs/plugins/
