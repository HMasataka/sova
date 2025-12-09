package editor

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/HMasataka/sova/internal/clipboard"
	"github.com/HMasataka/sova/internal/config"
	"github.com/HMasataka/sova/internal/history"
)

// EditAndCopy opens a temporary file in the configured editor, and copies the content to clipboard.
// It also saves the content to history.
func EditAndCopy(cfg *config.Config) error {
	tmpFile, err := os.CreateTemp("", "edit_tmp_*.txt")
	if err != nil {
		return fmt.Errorf("failed to create temp file: %w", err)
	}
	defer os.Remove(tmpFile.Name())

	originalContent := ""
	if err := tmpFile.Close(); err != nil {
		return fmt.Errorf("failed to close temp file: %w", err)
	}

	// Parse editor command and arguments
	editorParts := strings.Fields(cfg.Editor)
	if len(editorParts) == 0 {
		return fmt.Errorf("editor command is empty")
	}

	// Build command with editor arguments and temp file
	args := append(editorParts[1:], tmpFile.Name())
	cmd := exec.Command(editorParts[0], args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("%s execution failed: %w", cfg.Editor, err)
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

	if err := history.Save(cfg, content); err != nil {
		fmt.Fprintf(os.Stderr, "Warning: Failed to save to history: %v\n", err)
	}

	return nil
}
