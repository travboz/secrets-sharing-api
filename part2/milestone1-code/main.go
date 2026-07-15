package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"slices"
)

var (
	programName = "secret-share"
	usageString = `A HTTP client CLI for interacting with the Secret Sharing web application.
With this CLI we can create and view secrets using the application's API.

Usage: %s command [options]
`

	ActionView   = "view"
	ActionCreate = "create"
)

func printUsage(w io.Writer) {
	fmt.Fprintf(w, usageString, programName)
}

func main() {
	// Check if length of args is valid, we expect: cli create|view
	// So, if len of os.Args < 2, not valid.
	if len(os.Args) < 2 {
		printUsage(os.Stderr)
		os.Exit(1)
	}

	if slices.Contains(os.Args, "--help") || slices.Contains(os.Args, "-h") {
		fmt.Fprintln(os.Stdout, "Help called")
		printUsage(os.Stdout)
		os.Exit(1)
	}

	// We have more than 2 args, so determine which operation is which:
	subcommand := os.Args[1]
	args := os.Args[2:]

	var err error
	var result ParseResult

	switch subcommand {

	case ActionCreate:
		result, err = parseCreateArgs(os.Stderr, args)
		if err != nil {
			if errors.Is(err, ErrInvalidPosArgSpecified) {
				fmt.Fprintln(os.Stdout, err)
				result.Usage()
			}
			os.Exit(1)
		}

		result.Config.Action = ActionCreate

		if err = validateCreateArgs(result.Config); err != nil {
			fmt.Fprintln(os.Stdout, err)
			result.Usage()
			os.Exit(1)
		}

		if err = runCreateSecret(os.Stdout, result.Config); err != nil {
			fmt.Fprintln(os.Stdout, err)
			os.Exit(1)
		}

	case ActionView:
		result, err = parseViewArgs(os.Stderr, args)
		if err != nil {
			if errors.Is(err, ErrInvalidPosArgSpecified) {
				fmt.Fprintln(os.Stdout, err)
				result.Usage()
			}
			os.Exit(1)
		}

		result.Config.Action = ActionView

		if err = validateViewArgs(result.Config); err != nil {
			fmt.Fprintln(os.Stdout, err)
			result.Usage()
			os.Exit(1)
		}

		if err = runViewSecret(os.Stdout, result.Config); err != nil {
			fmt.Fprintln(os.Stdout, err)
			os.Exit(1)
		}
	}
}
