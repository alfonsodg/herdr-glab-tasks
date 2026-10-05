// Command herdr-gitlab-issues is the entry point for the plugin commands.
package main

import (
	"fmt"
	"os"
)

const usage = "usage: herdr-gitlab-issues <panel|new>"

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, usage)
		os.Exit(2)
	}
	if err := run(os.Args[1]); err != nil {
		fmt.Fprintf(os.Stderr, "herdr-gitlab-issues %s: %v\n", os.Args[1], err)
		os.Exit(1)
	}
}

func run(cmd string) error {
	switch cmd {
	case "panel", "new":
		return fmt.Errorf("not implemented: %s", cmd)
	default:
		return fmt.Errorf("unknown command: %s", cmd)
	}
}
