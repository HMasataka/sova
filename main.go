package main

import (
	"fmt"
	"os"

	"github.com/HMasataka/sova/internal/config"
	"github.com/HMasataka/sova/internal/editor"
	"github.com/HMasataka/sova/internal/history"
	"github.com/jessevdk/go-flags"
)

type Options struct {
	History bool `short:"H" long:"history" description:"Show clipboard history"`
}

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading config: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Configuration loaded: %+v\n", cfg)

	var opts Options
	parser := flags.NewParser(&opts, flags.Default)

	if _, err := parser.Parse(); err != nil {
		if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.ErrHelp {
			os.Exit(0)
		}
		os.Exit(1)
	}

	if opts.History {
		if err := history.Show(cfg); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		return
	}

	if err := editor.EditAndCopy(cfg); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
