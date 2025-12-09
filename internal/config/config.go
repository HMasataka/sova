package config

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// Config represents the application configuration.
type Config struct {
	Editor            string `yaml:"editor"`
	HistoryPath       string `yaml:"history_path"`
	MaxHistoryEntries int    `yaml:"max_history_entries"`
}

// Load loads the configuration from ~/.sova/config.yaml.
// If the file doesn't exist, it returns the default configuration.
func Load() (*Config, error) {
	cfg := defaultConfig()

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("failed to get home directory: %w", err)
	}

	configPath := filepath.Join(homeDir, ".sova", "config.yaml")
	data, err := os.ReadFile(configPath)
	if err != nil {
		if os.IsNotExist(err) {
			return cfg, nil
		}
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	if err := yaml.Unmarshal(data, cfg); err != nil {
		return nil, fmt.Errorf("failed to parse config file: %w", err)
	}

	// Expand ~ in paths
	cfg.HistoryPath = expandPath(cfg.HistoryPath, homeDir)

	return cfg, nil
}

// expandPath expands ~ in the path to the home directory.
func expandPath(path, homeDir string) string {
	if len(path) > 0 && path[0] == '~' {
		if len(path) == 1 || path[1] == '/' {
			return filepath.Join(homeDir, path[1:])
		}
	}
	return path
}

// defaultConfig returns the default configuration.
func defaultConfig() *Config {
	homeDir, _ := os.UserHomeDir()
	return &Config{
		Editor:            "nvim",
		HistoryPath:       filepath.Join(homeDir, ".sova", "history.txt"),
		MaxHistoryEntries: 0, // 0 means unlimited
	}
}
