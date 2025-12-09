package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "--history" {
		if err := showHistory(); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		return
	}

	if err := editTempFile(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func editTempFile() error {
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
	if err := copyToClipboard(content); err != nil {
		fmt.Println("Failed to copy to clipboard. Output:")
		fmt.Println(content)
		return nil
	}

	fmt.Println("Content copied to clipboard successfully")

	if err := saveToHistory(content); err != nil {
		fmt.Fprintf(os.Stderr, "Warning: Failed to save to history: %v\n", err)
	}

	return nil
}

func copyToClipboard(text string) error {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("pbcopy")
	case "linux":
		if _, err := exec.LookPath("xclip"); err == nil {
			cmd = exec.Command("xclip", "-selection", "clipboard")
		} else if _, err := exec.LookPath("xsel"); err == nil {
			cmd = exec.Command("xsel", "--clipboard", "--input")
		} else {
			return fmt.Errorf("no clipboard command found (xclip or xsel)")
		}
	case "windows":
		cmd = exec.Command("clip")
	default:
		return fmt.Errorf("unsupported platform: %s", runtime.GOOS)
	}

	cmd.Stdin = strings.NewReader(text)

	return cmd.Run()
}

func saveToHistory(content string) error {
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

func showHistory() error {
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
