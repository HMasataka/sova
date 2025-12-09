package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/HMasataka/sova/internal/storage"
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

	configPath, err := storage.GetConfigPath()
	if err != nil {
		return nil, err
	}

	data, err := storage.ReadFile(configPath)
	if err != nil {
		if os.IsNotExist(err) {
			return cfg, nil
		}
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	if err := yaml.Unmarshal(data, cfg); err != nil {
		return nil, fmt.Errorf("failed to parse config file: %w", err)
	}

	cfg.HistoryPath, err = storage.ExpandPath(cfg.HistoryPath)
	if err != nil {
		return nil, fmt.Errorf("failed to expand history path: %w", err)
	}

	return cfg, nil
}

// defaultConfig returns the default configuration.
func defaultConfig() *Config {
	sovaDir, _ := storage.GetSovaDir()
	return &Config{
		Editor:            "nvim",
		HistoryPath:       filepath.Join(sovaDir, "history.txt"),
		MaxHistoryEntries: 0, // 0 means unlimited
	}
}
