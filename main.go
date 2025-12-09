package main

import (
	"fmt"
	"os"

	"github.com/HMasataka/sova/internal/editor"
	"github.com/HMasataka/sova/internal/history"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "--history" {
		if err := history.Show(); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		return
	}

	if err := editor.EditAndCopy(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
