# USER_GUIDE.md

| Field | Value |
| ----- | ----- |
| **Project** | herdr-glab-tasks — Herdr plugin for GitLab Issues |
| **Version** | v0.1.0 |
| **Created** | 2026-10-05 |
| **Updated** | 2026-10-05 02:00 |
| **Author** | Alfonso de la Guarda Reyes |
| **Location** | `docs/USER_GUIDE.md` |

How to use the GitLab Issues plugin from Herdr. For setup, see
[INSTALL.md](INSTALL.md).

## Panel: see workspace issues

Press `Ctrl+B i` to open the **GitLab Issues** pane in Herdr. It runs the
`issues` action, which opens the pane entry. Equivalent terminal form:

```sh
bin/herdr-gitlab-issues panel
```

Output groups issues by status column:

```text
## todo (2)
#12 Fix login redirect [priority::high, status::todo]
#18 Update docs [status::todo]

## review (1)
#21 Refactor auth [status::review]
```

The last line shows the current branch issue when the latest commit
carries it: `branch: Ref #21`.

## Filter by state and label

```sh
bin/herdr-gitlab-issues panel --state all
bin/herdr-gitlab-issues panel --label priority::high
bin/herdr-gitlab-issues panel --state all --label type::bug
```

States are `opened` (default), `closed`, and `all`. The label flag
matches one GitLab label exactly, for example `status::review`.

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
