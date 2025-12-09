package history

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// Save saves the given content to the history file with a timestamp.
func Save(content string) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to get home directory: %w", err)
	}

	sovaDir := filepath.Join(homeDir, ".sova")
	if err := os.MkdirAll(sovaDir, 0755); err != nil {
		return fmt.Errorf("failed to create .sova directory: %w", err)
	}

	historyFile := filepath.Join(sovaDir, "history.txt")
	f, err := os.OpenFile(historyFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
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

// Show displays all history entries from the history file.
func Show() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to get home directory: %w", err)
	}

	historyFile := filepath.Join(homeDir, ".sova", "history.txt")
	content, err := os.ReadFile(historyFile)
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
