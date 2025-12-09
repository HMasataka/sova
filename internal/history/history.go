package history

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/HMasataka/sova/internal/config"
)

// Save saves the given content to the history file with a timestamp.
func Save(cfg *config.Config, content string) error {
	historyDir := filepath.Dir(cfg.HistoryPath)
	if err := os.MkdirAll(historyDir, 0755); err != nil {
		return fmt.Errorf("failed to create history directory: %w", err)
	}

	// If max entries is set, enforce the limit
	if cfg.MaxHistoryEntries > 0 {
		if err := enforceMaxEntries(cfg); err != nil {
			return err
		}
	}

	f, err := os.OpenFile(cfg.HistoryPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open history file: %w", err)
	}
	defer f.Close()

	timestamp := time.Now().Format("2006-01-02 15:04:05")
	entry := fmt.Sprintf("=== %s ===\n%s\n\n", timestamp, content)

	if _, err := f.WriteString(entry); err != nil {
		return fmt.Errorf("failed to write to history: %w", err)
	}

	return nil
}

// enforceMaxEntries removes old entries if the count exceeds the limit.
func enforceMaxEntries(cfg *config.Config) error {
	data, err := os.ReadFile(cfg.HistoryPath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return fmt.Errorf("failed to read history file: %w", err)
	}

	content := string(data)
	entries := strings.Split(content, "=== ")
	entries = entries[1:] // Skip the first empty element

	if len(entries) < cfg.MaxHistoryEntries {
		return nil
	}

	// Keep only the most recent entries
	keepCount := cfg.MaxHistoryEntries - 1
	if keepCount <= 0 {
		// Clear all entries to make room for the new one
		if err := os.WriteFile(cfg.HistoryPath, []byte(""), 0644); err != nil {
			return fmt.Errorf("failed to clear history: %w", err)
		}
		return nil
	}

	entriesToKeep := entries[len(entries)-keepCount:]
	newContent := "=== " + strings.Join(entriesToKeep, "=== ")

	if err := os.WriteFile(cfg.HistoryPath, []byte(newContent), 0644); err != nil {
		return fmt.Errorf("failed to write trimmed history: %w", err)
	}

	return nil
}

// Show displays all history entries from the history file.
func Show(cfg *config.Config) error {
	content, err := os.ReadFile(cfg.HistoryPath)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("No history found.")
			return nil
		}
		return fmt.Errorf("failed to read history file: %w", err)
	}

	if len(content) == 0 {
		fmt.Println("No history found.")
		return nil
	}

	fmt.Print(string(content))
	return nil
}
