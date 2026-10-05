// Package pane opens the plugin pane through the herdr CLI.
package pane

import (
	"context"
	"os"
	"os/exec"
)

type Opener struct {
	Herdr string
	Run   func(ctx context.Context, name string, args ...string) error
}

func OpenArgs(pluginID, entrypoint string) []string {
	return []string{"plugin", "pane", "open", "--plugin", pluginID, "--entrypoint", entrypoint}
}

func (o *Opener) Open(ctx context.Context, pluginID, entrypoint string) error {
	herdr := o.Herdr
	if herdr == "" {
		herdr = os.Getenv("HERDR_BIN_PATH")
	}
	if herdr == "" {
		herdr = "herdr"
	}
	run := o.Run
	if run == nil {
		run = func(ctx context.Context, name string, args ...string) error {
			return exec.CommandContext(ctx, name, args...).Run()
		}
	}
	return run(ctx, herdr, OpenArgs(pluginID, entrypoint)...)
}
