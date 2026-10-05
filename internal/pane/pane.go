// Package pane opens the plugin pane through the herdr CLI.
package pane

import (
	"context"
	"encoding/json"
	"os"
	"os/exec"
)

type Opener struct {
	Herdr  string
	Run    func(ctx context.Context, name string, args ...string) error
	Output func(ctx context.Context, name string, args ...string) ([]byte, error)
}

func OpenArgs(pluginID, entrypoint, cwd string) []string {
	args := []string{"plugin", "pane", "open", "--plugin", pluginID, "--entrypoint", entrypoint}
	if cwd != "" {
		args = append(args, "--cwd", cwd)
	}
	return args
}

func FocusedCwd(out []byte) string {
	var data struct {
		Result struct {
			Pane struct {
				ForegroundCwd string `json:"foreground_cwd"`
				Cwd           string `json:"cwd"`
			} `json:"pane"`
		} `json:"result"`
	}
	if err := json.Unmarshal(out, &data); err != nil {
		return ""
	}
	if data.Result.Pane.ForegroundCwd != "" {
		return data.Result.Pane.ForegroundCwd
	}
	return data.Result.Pane.Cwd
}

func (o *Opener) herdrBin() string {
	if o.Herdr != "" {
		return o.Herdr
	}
	if env := os.Getenv("HERDR_BIN_PATH"); env != "" {
		return env
	}
	return "herdr"
}

func (o *Opener) output(ctx context.Context, name string, args ...string) ([]byte, error) {
	if o.Output != nil {
		return o.Output(ctx, name, args...)
	}
	return exec.CommandContext(ctx, name, args...).Output()
}

func (o *Opener) ResolveCwd(ctx context.Context) string {
	out, err := o.output(ctx, o.herdrBin(), "pane", "current")
	if err != nil {
		return ""
	}
	return FocusedCwd(out)
}

func (o *Opener) Open(ctx context.Context, pluginID, entrypoint, cwd string) error {
	if cwd == "" {
		cwd = o.ResolveCwd(ctx)
	}
	run := o.Run
	if run == nil {
		run = func(ctx context.Context, name string, args ...string) error {
			return exec.CommandContext(ctx, name, args...).Run()
		}
	}
	return run(ctx, o.herdrBin(), OpenArgs(pluginID, entrypoint, cwd)...)
}
