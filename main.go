package main

import (
	"fmt"
	"os"

	"github.com/HMasataka/sova/internal/editor"
	"github.com/HMasataka/sova/internal/history"
	"github.com/jessevdk/go-flags"
)

type Options struct {
	History bool `short:"H" long:"history" description:"Show clipboard history"`
}

func main() {
	var opts Options
	parser := flags.NewParser(&opts, flags.Default)

	if _, err := parser.Parse(); err != nil {
		if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.ErrHelp {
			os.Exit(0)
		}
		os.Exit(1)
	}

	if opts.History {
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
