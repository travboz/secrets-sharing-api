package main

import (
	"fmt"
	"os"
)

func main() {
	// Check if length of args is valid, we expect: cli subcommand
	if len(os.Args) < 2 {
		printUsage(os.Stderr)
		os.Exit(1)
	}

	// Help check
	if helpCalled(os.Args) {
		printUsage(os.Stdout)
		os.Exit(1)
	}

	// We have at least 2 args, so determine if a subcommand has been called.
	result, err := parseArgs(os.Stdout, os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	err = validateConfig(result.Config)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	output, err := performAction(result.Config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	fmt.Fprintln(os.Stdout, output)
}
