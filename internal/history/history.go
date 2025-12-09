package history

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/HMasataka/sova/internal/config"
	"github.com/HMasataka/sova/internal/storage"
)

// Save saves the given content to the history file with a timestamp.
func Save(cfg *config.Config, content string) error {
	if err := storage.EnsureDir(cfg.HistoryPath); err != nil {
		return err
	}

	// If max entries is set, enforce the limit
	if cfg.MaxHistoryEntries > 0 {
		if err := enforceMaxEntries(cfg); err != nil {
			return err
		}
	}

	timestamp := time.Now().Format("2006-01-02 15:04:05")
	entry := fmt.Appendf(nil, "=== %s ===\n%s\n\n", timestamp, content)

	if err := storage.AppendToFile(cfg.HistoryPath, entry); err != nil {
		return err
	}

	return nil
}

// enforceMaxEntries removes old entries if the count exceeds the limit.
func enforceMaxEntries(cfg *config.Config) error {
	data, err := storage.ReadFile(cfg.HistoryPath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}

	content := string(data)
	entries := strings.Split(content, "=== ")
	entries = entries[1:] // Skip the first empty element

	if len(entries) < cfg.MaxHistoryEntries {
		return nil
	}

	// Keep only the most recent entries, leaving room for the new entry
	keepCount := cfg.MaxHistoryEntries - 1
	if keepCount <= 0 {
		// Clear all entries to make room for the new one
		return storage.WriteFile(cfg.HistoryPath, []byte(""))
	}

	entriesToKeep := entries[len(entries)-keepCount:]
	newContent := "=== " + strings.Join(entriesToKeep, "=== ")

	return storage.WriteFile(cfg.HistoryPath, []byte(newContent))
}

// Show displays all history entries from the history file.
func Show(cfg *config.Config) error {
	data, err := storage.ReadFile(cfg.HistoryPath)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("No history found.")
			return nil
		}
		return err
	}

	if len(data) == 0 {
		fmt.Println("No history found.")
		return nil
	}

	fmt.Print(string(data))

	return nil
}
