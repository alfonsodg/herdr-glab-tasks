// Package config reads the plugin config.toml from the Herdr config dir.
package config

import (
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

type Config struct {
	State       string `toml:"state"`
	Label       string `toml:"label"`
	BranchType  string `toml:"branch_type"`
	BranchScope string `toml:"branch_scope"`
	file        string
}

func Default() Config {
	return Config{
		State:       "opened",
		BranchType:  "feature",
		BranchScope: "tasks",
	}
}

func Load(dir string) Config {
	cfg := Default()
	if dir == "" {
		return cfg
	}
	path := filepath.Join(dir, "config.toml")
	cfg.file = path
	if _, err := os.Stat(path); err != nil {
		return cfg
	}
	var file Config
	if _, err := toml.DecodeFile(path, &file); err != nil {
		return cfg
	}
	merge(&cfg, file)
	return cfg
}

func merge(cfg *Config, file Config) {
	if file.State != "" {
		cfg.State = file.State
	}
	if file.Label != "" {
		cfg.Label = file.Label
	}
	if file.BranchType != "" {
		cfg.BranchType = file.BranchType
	}
	if file.BranchScope != "" {
		cfg.BranchScope = file.BranchScope
	}
}

func (c Config) Path() string {
	return c.file
}
