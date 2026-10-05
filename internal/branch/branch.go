// Package branch bridges issues and code: branch creation and ref parsing.
package branch

import (
	"errors"
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

var refPattern = regexp.MustCompile(`Ref #(\d+)`)

func Name(branchType, scope string, iid int) string {
	return fmt.Sprintf("%s/%s-#%d", branchType, scope, iid)
}

func CreateIssueBranch(dir, branchType, scope string, iid int) error {
	if dirty(dir) {
		return errors.New("working tree is dirty: commit or stash first")
	}
	name := Name(branchType, scope, iid)
	if out, err := exec.Command("git", "-C", dir, "checkout", "-b", name).CombinedOutput(); err != nil {
		return fmt.Errorf("checkout -b %s: %v\n%s", name, err, out)
	}
	return nil
}

func IssueRef(dir string) (int, bool) {
	out, err := exec.Command("git", "-C", dir, "log", "-1", "--pretty=%B").Output()
	if err != nil {
		return 0, false
	}
	m := refPattern.FindStringSubmatch(string(out))
	if m == nil {
		return 0, false
	}
	iid, err := strconv.Atoi(m[1])
	if err != nil {
		return 0, false
	}
	return iid, true
}

func dirty(dir string) bool {
	out, err := exec.Command("git", "-C", dir, "status", "--porcelain").Output()
	if err != nil {
		return true
	}
	return strings.TrimSpace(string(out)) != ""
}
