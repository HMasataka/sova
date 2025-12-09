package editor

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/HMasataka/sova/internal/clipboard"
	"github.com/HMasataka/sova/internal/history"
)

// EditAndCopy opens a temporary file in nvim, and copies the content to clipboard.
// It also saves the content to history.
func EditAndCopy() error {
	tmpFile, err := os.CreateTemp("", "edit_tmp_*.txt")
	if err != nil {
		return fmt.Errorf("failed to create temp file: %w", err)
	}
	defer os.Remove(tmpFile.Name())

	originalContent := ""
	if err := tmpFile.Close(); err != nil {
		return fmt.Errorf("failed to close temp file: %w", err)
	}

	cmd := exec.Command("nvim", tmpFile.Name())
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("nvim execution failed: %w", err)
	}

	editedContent, err := os.ReadFile(tmpFile.Name())
	if err != nil {
		return fmt.Errorf("failed to read edited file: %w", err)
	}

	if strings.TrimSpace(string(editedContent)) == strings.TrimSpace(originalContent) {
		fmt.Println("No changes detected.")
		return nil
	}

	content := strings.TrimSuffix(string(editedContent), "\n")
	if err := clipboard.Copy(content); err != nil {
		fmt.Println("Failed to copy to clipboard. Output:")
		fmt.Println(content)
		return nil
	}

	fmt.Println("Content copied to clipboard successfully")

	if err := history.Save(content); err != nil {
		fmt.Fprintf(os.Stderr, "Warning: Failed to save to history: %v\n", err)
	}

	return nil
}
