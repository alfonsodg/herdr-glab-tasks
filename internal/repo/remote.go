// Package repo maps workspaces to GitLab projects.
package repo

import (
	"net/url"
	"strings"
)

func ParseRemote(raw string) (host, project string, ok bool) {
	s := strings.TrimSpace(raw)
	if strings.Contains(s, "://") {
		u, err := url.Parse(s)
		if err != nil {
			return "", "", false
		}
		host, project = u.Hostname(), u.Path
	} else {
		hostPart, pathPart, found := strings.Cut(s, ":")
		if !found {
			return "", "", false
		}
		if i := strings.LastIndex(hostPart, "@"); i >= 0 {
			hostPart = hostPart[i+1:]
		}
		host, project = hostPart, pathPart
	}
	project = strings.Trim(strings.TrimSuffix(strings.Trim(project, "/"), ".git"), "/")
	if host == "" || project == "" {
		return "", "", false
	}
	return strings.ToLower(host), strings.ToLower(project), true
}
