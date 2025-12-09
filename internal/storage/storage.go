package storage

import (
	"fmt"
	"os"
	"path/filepath"
)

const (
	sovaDir        = ".sova"
	configFileName = "config.yaml"
)

// GetSovaDir returns the path to the .sova directory.
func GetSovaDir() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to get home directory: %w", err)
	}
	return filepath.Join(homeDir, sovaDir), nil
}

// EnsureSovaDir creates the .sova directory if it doesn't exist.
func EnsureSovaDir() error {
	dir, err := GetSovaDir()
	if err != nil {
		return err
	}
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create .sova directory: %w", err)
	}
	return nil
}

// GetConfigPath returns the path to the config file.
func GetConfigPath() (string, error) {
	dir, err := GetSovaDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, configFileName), nil
}

// ExpandPath expands ~ in the path to the home directory.
func ExpandPath(path string) (string, error) {
	if len(path) == 0 {
		return path, nil
	}

	if path[0] == '~' {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return "", fmt.Errorf("failed to get home directory: %w", err)
		}
		if len(path) == 1 || path[1] == '/' {
			return filepath.Join(homeDir, path[1:]), nil
		}
	}
	return path, nil
}

// EnsureDir creates the directory for the given file path if it doesn't exist.
func EnsureDir(filePath string) error {
	dir := filepath.Dir(filePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}
	return nil
}

// ReadFile reads the content of a file.
func ReadFile(path string) ([]byte, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// WriteFile writes content to a file.
func WriteFile(path string, data []byte) error {
	if err := os.WriteFile(path, data, 0644); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}
	return nil
}

// AppendToFile appends content to a file.
func AppendToFile(path string, data []byte) error {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer f.Close()

	if _, err := f.Write(data); err != nil {
		return fmt.Errorf("failed to write to file: %w", err)
	}

	return nil
}

// FileExists checks if a file exists.
func FileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
